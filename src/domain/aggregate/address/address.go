package address

import "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"

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
