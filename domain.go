package uuid

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	domainPerson = "person"
	domainGroup  = "group"
	domainOrg    = "org"
)

func domainTypeFromString(str string) (uuid.Domain, error) {
	if str == "" {
		return 0, errors.New("domain is required")
	}
	switch str {
	case domainPerson:
		return uuid.Person, nil
	case domainGroup:
		return uuid.Group, nil
	case domainOrg:
		return uuid.Org, nil
	}
	return 0, fmt.Errorf("invalid domain: %s", str)
}
