package hash

import (
	"crypto/sha256"
	"fmt"
	"regexp"
)

const (
	//GlobalDBVersion ...
	GlobalDBVersion int = 1
)

// AsSha256 hash function
func AsSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// AsSha256OfOrderedMap hash function
// My tests showed this is unnecessary because apparently, the order of the key is the same
// if the content of the row is the same. So, two rows could have different field order, but
// it should be always the same for the specific row
func AsSha256OfOrderedMap(o map[string]interface{}, orderedKey []string) string {
	type ordered struct {
		key   string
		value interface{}
	}
	var list []ordered
	for _, j := range orderedKey {
		a := o[j]
		list = append(list, ordered{
			key:   j,
			value: a,
		})
	}
	return AsSha256(list)
}

// CleanAndTrim just alphanumeric chars
func CleanAndTrim(text string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(text, "")
}
