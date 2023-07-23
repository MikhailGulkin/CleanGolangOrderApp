package aggregate

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/address/vo"
)

type Address struct {
	vo.AddressID
	BuildingNumber int
	StreetName     string
	City           string
	Country        string
}

func (Address) Create(id vo.AddressID, buildingNumber int, streetName string, city string, country string) Address {
	return Address{AddressID: id, City: city, StreetName: streetName, BuildingNumber: buildingNumber, Country: country}
}
func (a *Address) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %d", a.Country, a.City, a.StreetName, a.BuildingNumber)
}
