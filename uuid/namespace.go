package uuid

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	namespaceDNS  = "dns"
	namespaceURL  = "url"
	namespaceOID  = "oid"
	namespaceX500 = "x500"
)

func namespaceTypeFromString(str string) (uuid.UUID, error) {
	if str == "" {
		return uuid.Nil, errors.New("namespace is required")
	}
	switch str {
	case namespaceDNS:
		return uuid.NameSpaceDNS, nil
	case namespaceURL:
		return uuid.NameSpaceURL, nil
	case namespaceOID:
		return uuid.NameSpaceOID, nil
	case namespaceX500:
		return uuid.NameSpaceX500, nil
	}
	return uuid.Nil, fmt.Errorf("invalid namespace: %s", str)
}
