package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/husniadil/uuid"
)

// ParseUUIDMetadata parses uuid metadata.
func ParseUUIDMetadata(w http.ResponseWriter, r *http.Request) {
	paramUUID := r.URL.Query().Get("uuid")
	metadata, err := uuid.Parse(paramUUID)
	if err != nil {
		if errors.Is(err, uuid.ErrUUIDValidation) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintln(w, err)
		return
	}
	byteJSON, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(byteJSON))
}
