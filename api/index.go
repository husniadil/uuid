// Package api contains handlers for serverless functions deployed on vercel.
package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/husniadil/uuid"
)

const defaultMaxSize = 500

func getMaxSize() int {
	strMaxSize := os.Getenv("MAX_SIZE")
	maxSize, err := strconv.Atoi(strMaxSize)
	if err != nil {
		maxSize = defaultMaxSize
	}
	return maxSize
}

// GenerateUUID handles http requests.
func GenerateUUID(w http.ResponseWriter, r *http.Request) {
	strVersion := r.URL.Query().Get("version")
	version, err := strconv.Atoi(strVersion)
	if err != nil {
		version = 0
	}

	var hypen bool = false
	if _, ok := r.URL.Query()["hypen"]; ok {
		hypen = true
	}

	var uppercase bool = false
	if _, ok := r.URL.Query()["uppercase"]; ok {
		uppercase = true
	}

	strSize := r.URL.Query().Get("size")
	size, err := strconv.Atoi(strSize)
	if err != nil {
		size = 1
	}
	maxSize := getMaxSize()
	if size > maxSize {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "max allowed size is %d", maxSize)
		return
	}

	domain := r.URL.Query().Get("domain")

	strID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strID)
	if err != nil || id < 0 {
		id = 1
	}

	namespace := r.URL.Query().Get("namespace")

	data := r.URL.Query().Get("data")

	var download bool = false
	if _, ok := r.URL.Query()["download"]; ok {
		download = true
	}

	req := uuid.Request{
		Version:   version,
		Domain:    domain,
		ID:        id,
		Namespace: namespace,
		Data:      data,
	}

	uuids, err := uuid.Generate(size, hypen, uppercase, req)
	if err != nil {
		if errors.Is(err, uuid.ErrUUIDValidation) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintln(w, err)
		return
	}
	if download {
		fileName := generateFileName(version)
		fileSize := generateFileSize(size, uuids[0])
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
		w.Header().Add("Content-Length", strconv.Itoa(fileSize))
	}
	w.WriteHeader(http.StatusOK)
	for _, uuid := range uuids {
		fmt.Fprintln(w, uuid)
	}
}

func generateFileName(version int) string {
	fileName := "nil_uuid.txt"
	if version > 0 {
		fileName = fmt.Sprintf("version%d_uuid.txt", version)
	}
	return fileName
}

func generateFileSize(size int, firstUUID string) int {
	fileSize := (len(firstUUID) + 1) * size
	return fileSize
}
