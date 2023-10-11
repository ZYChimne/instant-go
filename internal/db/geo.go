package database

import (
	"zychimne/instant/pkg/model"
)

func GetCountries(countries *[]model.Country) error {
	return PostgresDB.Select("id", "name").Find(&countries).Error
}

func GetStatesByCountryID(cID uint, states *[]model.State) error {
	return PostgresDB.Select("id", "name").Where("country_id = ?", cID).Find(&states).Error
}

func GetCitiesByStateID(sID uint, cities *[]model.City) error {
	return PostgresDB.Select("id", "name").Where("state_id = ?", sID).Find(&cities).Error
}
