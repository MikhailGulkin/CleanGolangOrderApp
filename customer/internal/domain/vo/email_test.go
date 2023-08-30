package vo

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewValidEmail(t *testing.T) {
	email := "some@example.com"
	_, err := NewEmail(email)
	if err != nil {
		t.Error(err)
	}
}
func TestNewInvalidEmail(t *testing.T) {
	email := strings.Repeat("a", 256)
	_, err := NewEmail(fmt.Sprintf("some%s@emxample.com", email))
	if err == nil {
		t.Error(err)
	}
}
