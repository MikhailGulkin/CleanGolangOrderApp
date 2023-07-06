package address

type Address struct {
	BuildingNumber int
	StreetName     string
	City           string
	Country        string
}

func (Address) Create(buildingNumber int, streetName string, city string, country string) Address {
	return Address{City: city, StreetName: streetName, BuildingNumber: buildingNumber, Country: country}
}
