package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

type Country struct {
	Code          	string  `json:"code,omitempty"  db:"Code"`
	Name        	string 	`json:"name,omitempty"  db:"Name"`
	Continent 		string 	`json:"countryCode,omitempty"  db:"Continent"`
	Region    		string 	`json:"district,omitempty"  db:"Region"`
	SurfaceArea  	float64   `json:"population,omitempty"  db:"SurfaceArea"`
	IndepYear      	int 	`json:"indepyear,omitempty"  db:"IndepYear"`
	Population 		int 	`json:"poplation,omitempty"  db:"Population"`
	LifeExpectancy 	int 	`json:"lifeexpectancy,omitempty"  db:"LifeExpectancy"`
	GNP  			float64  	`json:"gnp,omitempty"  db:"GNP"`
	GNPOld  		float64  	`json:"gnpold,omitempty"  db:"GNPOld"`
	LocalName       string 	`json:"localname,omitempty"  db:"LocalName"`
	GovernmentForm 	string 	`json:"governmentform,omitempty"  db:"GovernmentForm"`
	HeadOfState    	string 	`json:"headofstate,omitempty"  db:"HeadOfState"`
	Captal  		int    	`json:"captal,omitempty"  db:"Capital"`
	Code2  			string 	`json:"code2,omitempty"  db:"Code2"`
}

var (
	db * sqlx.DB
)

func main() {
	_db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	db = _db

	e := echo.New()

	e.GET("/cities/:cityName",getCityInfoHander)

	e.Start(":12800")
}

func getCityInfoHander(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	city := City{}
	db.Get(&city, "SELECT * FROM city WHERE Name=?",cityName)
	if city.Name == "" {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, city)
}
