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
	name := "Smth"
	middleName := ""
	lastName := "SSS"
	_, err := NewFullName(name, middleName, lastName)
	if err == nil {
		t.Error(err)
	}
}
