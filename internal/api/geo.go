package api

import (
	"net/http"
	"strconv"

	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	var countries []model.Country
	err := database.GetCountries(&countries)
	if err != nil {
		handleError(c, err, "Database Error", DatabaseError)
	}
	resp := make([]schema.CountryResponse, len(countries))
	for i, country := range countries {
		resp[i] = schema.CountryResponse{
			ID:   country.ID,
			Name: country.Name,
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetStates(c *gin.Context) {
	cID := c.Query("cID")
	errMsg := "Country ID is required"
	if cID == "" {
		handleError(c, nil, errMsg, ParameterError)
	}
	id, err := strconv.ParseUint(cID, 10, 64)
	if err != nil {
		handleError(c, nil, errMsg, ParameterError)
	}
	var states []model.State
	err = database.GetStatesByCountryID(uint(id), &states)
	if err != nil {
		handleError(c, err, "Database Error", DatabaseError)
	}
	resp := make([]schema.StateResponse, len(states))
	for i, state := range states {
		resp[i] = schema.StateResponse{
			ID:   state.ID,
			Name: state.Name,
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": states})
}

func GetCities(c *gin.Context) {
	sID := c.Query("sID")
	errMsg := "State ID is required"
	if sID == "" {
		handleError(c, nil, errMsg, ParameterError)
	}
	id, err := strconv.ParseUint(sID, 10, 64)
	if err != nil {
		handleError(c, nil, errMsg, ParameterError)
	}
	var cities []model.City
	err = database.GetCitiesByStateID(uint(id), &cities)
	if err != nil {
		handleError(c, err, "DatabaseError", DatabaseError)
	}
	resp := make([]schema.CityResponse, len(cities))
	for i, city := range cities {
		resp[i] = schema.CityResponse{
			ID:   city.ID,
			Name: city.Name,
		}
	}
	c.JSON(http.StatusOK, gin.H{ "data": cities})
}
