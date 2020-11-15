package uuid

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Metadata represents a uuid metadata.
type Metadata struct {
	UUID          string `json:"uuid,omitempty"`
	Version       string `json:"version,omitempty"`
	Variant       string `json:"variant,omitempty"`
	Time          string `json:"time,omitempty"`
	Domain        string `json:"domain,omitempty"`
	ID            uint32 `json:"id,omitempty"`
	NodeID        []byte `json:"node_id,omitempty"`
	URN           string `json:"urn,omitempty"`
	ClockSequence int    `json:"clock_sequence,omitempty"`
}

// Parse parses uuid metadata.
func Parse(s string) (Metadata, error) {
	if s == "" {
		return Metadata{}, errors.Wrap(ErrUUIDValidation, "uuid is required")
	}
	parsedUUID, err := uuid.Parse(s)
	if err != nil {
		return Metadata{}, errors.Wrapf(ErrUUIDValidation, "invalid uuid: %s", s)
	}
	metadata := Metadata{
		UUID:    parsedUUID.String(),
		Version: parsedUUID.Version().String(),
		Variant: parsedUUID.Variant().String(),
		URN:     parsedUUID.URN(),
	}
	if parsedUUID.Version() == Version1 || parsedUUID.Version() == Version2 {
		sec, nsec := parsedUUID.Time().UnixTime()
		parsedTime := time.Unix(sec, nsec)
		metadata.Time = parsedTime.UTC().String()
		metadata.NodeID = parsedUUID.NodeID()
		metadata.ClockSequence = parsedUUID.ClockSequence()
	}
	if parsedUUID.Version() == Version2 {
		metadata.Domain = parsedUUID.Domain().String()
		metadata.ID = parsedUUID.ID()
	}
	return metadata, nil
}
