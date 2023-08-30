package vo

import "testing"

func TestValidFullName(t *testing.T) {
	name := "Mikhail"
	middleName := "Vladimirovich"
	lastName := "Gulkin"
	_, err := NewFullName(name, middleName, lastName)
	if err != nil {
		t.Error(err)
	}
}
func TestInvalidFullName(t *testing.T) {
	name := ""
	middleName := ""
	lastName := ""
	_, err := NewFullName(name, middleName, lastName)
	if err == nil {
		t.Error(err)
	}
}
