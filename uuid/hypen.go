package uuid

import "strings"

func removeHypen(str string) string {
	return strings.Replace(str, "-", "", -1)
}
