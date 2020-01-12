package main

import "github.com/gin-gonic/gin"
import "log"
import "net/http"
import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"

const NameParam = "name"

var db *sqlx.DB
var router *gin.Engine

type Patient struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Illness string `json:"illness" db:"illness"`
	BirthDate string `json:"birthDate" db:"birth_date"`
	LocationId int `json:"locationId" db:"location_id"`
}

type Location struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	HospitalId string `json:"hospitalId" db:"hospital_id"`
}

type Hospital struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	MaxPatientAmount int `json:"maxPatientCount" db:"max_patient_amount"`
}

func main() {
	initDatabase()
	initAPI()

	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func initDatabase() {
	db = sqlx.MustConnect("postgres", "postgres://ormkzpwm:x6uyb8EA034nb8rP8YMgy7udyYrYUuVw@manny.db.elephantsql.com:5432/ormkzpwm")
}

func initAPI() {
	router = gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	router.GET("/greet/:"+NameParam, func(c *gin.Context) {
		name := c.Param(NameParam)
		message := "Hello " + name + "!"
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})
	router.GET("/patient/all", GetAllPatients)
	router.GET("/location/all", GetAllLocations)
	router.GET("/hospital/all", GetAllHospitals)
}

func GetAllPatients(c *gin.Context) {
	var patients []Patient
	err := db.Select(&patients, "select * from patient")
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, patients)
}

func GetAllLocations(c *gin.Context) {
	var locations []Location
	err := db.Select(&locations, "select * from location")
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, locations)
}

func GetAllHospitals(c *gin.Context) {
	var hospitals []Hospital
	err := db.Select(&hospitals, "select * from hospital")
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, hospitals)
}
