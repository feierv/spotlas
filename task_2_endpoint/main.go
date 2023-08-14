package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"mywebapp/db"
	"strconv"
	"math"
)

func calculateBoundingBox(latitude, longitude, radius float64, areaType string) (minLat, maxLat, minLng, maxLng float64) {
	if areaType == "square" {
			minLat, maxLat = latitude-radius, latitude+radius
			minLng, maxLng = longitude-radius, longitude+radius
		} else if areaType == "circle" {
			const earthRadius = 6371000 // in meters
			latDiff := (radius / earthRadius) * (180 / math.Pi)
			lngDiff := (radius / earthRadius) * (180 / math.Pi) / math.Cos(latitude*math.Pi/180)
			minLat, maxLat = latitude-latDiff, latitude+latDiff
			minLng, maxLng = longitude-lngDiff, longitude+lngDiff
		}
	return minLat, maxLat, minLng, maxLng
}


func main() {
	// Create a new Gin router
	router := gin.Default()
	dbConn, err := db.Connect()
	fmt.Printf("dbConn value: %v\n", dbConn)
	if err != nil {
        fmt.Printf("Error connecting to the database: %s\n", err.Error())
        log.Fatal(err)
    }
	// Define a route and handler function
	router.GET("/spots", func(c *gin.Context) {
		// GET /spots?latitude=37.7749&longitude=-122.4194&radius=1000&type=circle
		latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
		longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
		radius, _ := strconv.ParseFloat(c.Query("radius"), 64)
		areaType := c.Query("type")
		minLat, maxLat, minLng, maxLng := calculateBoundingBox(latitude, longitude, radius, areaType)
		spots, err := db.QueryRecords(dbConn, minLng, minLat, maxLng, maxLat)

		if err != nil {
			log.Fatal(err)
		}
		// Return data as JSON in the response
		c.JSON(http.StatusOK, spots)
	})

	// Run the server on port 8002
	router.Run(":8002")
}
