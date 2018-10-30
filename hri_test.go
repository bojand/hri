package hri

import (
	"log"
	"regexp"
	"testing"
)

func TestRandom(t *testing.T) {
	res := Random()
	matched, err := regexp.MatchString("\\b([a-z])+-([a-z])+-(\\d{1,2})\\b", res)
	if err != nil {
		log.Fatal(err)
	}

	if !matched {
		t.Errorf("Result does not match regex. Result: %s", res)
	}
}
