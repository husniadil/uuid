package uuid

import (
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrUUIDValidation represents an invalid input.
	ErrUUIDValidation = errors.New("validation error")
)

const (
	// VersionNil represents nil uuid.
	VersionNil = iota
	// Version1 represents version 1 uuid.
	Version1
	// Version2 represents version 2 uuid.
	Version2
	// Version3 represents version 3 uuid.
	Version3
	// Version4 represents version 4 uuid.
	Version4
	// Version5 represents version 5 uuid.
	Version5
)

// Request stores uuid generation request.
type Request struct {
	Version   int
	Domain    string
	ID        int
	Namespace string
	Data      string
}

// generate generates UUID.
func generate(req Request) (uuid.UUID, error) {
	var result uuid.UUID
	var domain uuid.Domain
	var namespace uuid.UUID
	var err error
	if req.Version < VersionNil || req.Version > Version5 {
		return uuid.Nil, errors.Wrapf(ErrUUIDValidation, "invalid version: %d", req.Version)
	}
	if req.Version == Version2 {
		domain, err = domainTypeFromString(req.Domain)
		if err != nil {
			return uuid.Nil, errors.Wrap(ErrUUIDValidation, err.Error())
		}
		if req.ID < 0 {
			return uuid.Nil, errors.Wrapf(ErrUUIDValidation, "invalid id: %d", req.ID)
		}
	}
	if req.Version == Version3 || req.Version == Version5 {
		namespace, err = namespaceTypeFromString(req.Namespace)
		if err != nil {
			return uuid.Nil, errors.Wrap(ErrUUIDValidation, err.Error())
		}
		if req.Data == "" {
			return uuid.Nil, errors.Wrapf(ErrUUIDValidation, "data is required")
		}
	}

	switch req.Version {
	case VersionNil:
		result = uuid.Nil
	case Version1:
		result, err = uuid.NewUUID()
	case Version2:
		result, err = uuid.NewDCESecurity(domain, uint32(req.ID))
	case Version3:
		result = uuid.NewMD5(namespace, []byte(req.Data))
	case Version4:
		result, err = uuid.NewRandom()
	case Version5:
		result = uuid.NewSHA1(namespace, []byte(req.Data))
	}
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error generating uuid")
	}
	return result, nil
}

// Generate generates UUIDs.
func Generate(size int, hypen, uppercase bool, req Request) ([]string, error) {
	var result []string
	if size < 1 {
		return nil, errors.Wrapf(ErrUUIDValidation, "invalid size: %d", size)
	}
	for i := 0; i < size; i++ {
		generatedUUID, err := generate(req)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		decoratedUUID := generatedUUID.String()
		if hypen == false {
			decoratedUUID = strings.Replace(decoratedUUID, "-", "", -1)
		}
		if uppercase {
			decoratedUUID = strings.ToUpper(decoratedUUID)
		}
		result = append(result, decoratedUUID)
	}
	return result, nil
}
