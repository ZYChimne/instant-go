package database

import (
	"errors"
	"zychimne/instant/pkg/schema"
)

func GetCountries(countries *[]schema.CountryResponse) error {
	return PostgresDB.Table("countries").Select("id", "name").Scan(&countries).Error
}

func GetStatesByCountryID(cID uint, states *[]schema.StateResponse) error {
	return PostgresDB.Table("states").Select("id", "name").Where("country_id = ?", cID).Scan(&states).Error
}

func GetCitiesByStateID(sID uint, cities *[]schema.CityResponse) error {
	return PostgresDB.Table("cities").Select("id", "name").Where("state_id = ?", sID).Scan(&cities).Error
}

func CheckGeo(country string, state string, city string, _ string) error {
	var count int64
	if err := PostgresDB.Table("cities").Select("id").Where("country_name = ? and state_name = ? and name = ?", country, state, city).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return errors.New(GeoCheckerError)
}
