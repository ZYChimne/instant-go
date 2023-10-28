package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	var countries []schema.CountryResponse
	err := database.GetCountries(&countries)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetCountriesError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": countries})
}

func GetStates(c *gin.Context) {
	cID, err := strconv.ParseUint(c.Query("cID"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetStatesError))
		return
	}
	var states []schema.StateResponse
	err = database.GetStatesByCountryID(uint(cID), &states)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetStatesError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": states})
}

func GetCities(c *gin.Context) {
	sID, err := strconv.ParseUint(c.Query("sID"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetCitiesError))
		return
	}
	var cities []schema.CityResponse
	err = database.GetCitiesByStateID(uint(sID), &cities)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetCitiesError))
		return
	}
	c.JSON(http.StatusOK, gin.H{ "data": cities})
}
