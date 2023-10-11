package schema

type CountryResponse struct {
	ID uint `json:"countryID"`
	Name string `json:"name"`
}

type CityResponse struct {
	ID uint `json:"cityID"`
	Name string `json:"name"`
}

type StateResponse struct {
	ID uint `json:"stateID"`
	Name string `json:"name"`
}
