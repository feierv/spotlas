package db

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/satori/go.uuid"
)

type YourStruct struct {
	ID          uuid.UUID
	Name        string
	Website     sql.NullString
	Latitude    float64
	Longitude   float64
	Description sql.NullString
	Rating      float64
}

func QueryRecords(db *sql.DB, minLng, minLat, maxLng, maxLat float64) ([]YourStruct, error) {
	sqlQuery := fmt.Sprintf(`
	SELECT
		id,
		name,
		website,
		ST_Y(coordinates::geometry) AS latitude,
		ST_X(coordinates::geometry) AS longitude,
		description,
		rating
	FROM "MY_TABLE"
	WHERE ST_Within(ST_SetSRID(ST_MakePoint(ST_X(coordinates::geometry), ST_Y(coordinates::geometry)), 4326), ST_MakeEnvelope(%f, %f, %f, %f, 4326))
	ORDER BY ST_Distance(ST_SetSRID(ST_MakePoint(ST_X(coordinates::geometry), ST_Y(coordinates::geometry)), 4326), coordinates),
         CASE WHEN ST_Distance(ST_SetSRID(ST_MakePoint(ST_X(coordinates::geometry), ST_Y(coordinates::geometry)), 4326), coordinates) < 50 THEN rating ELSE 0 END DESC;
		 `, minLng, minLat, maxLng, maxLat)

	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []YourStruct
	for rows.Next() {
		var id uuid.UUID
		var name string
		var latitude, longitude float64
		var description, website sql.NullString
		var rating float64

		err := rows.Scan(&id, &name, &website, &latitude, &longitude, &description, &rating)
		if err != nil {
			log.Fatal(err)
		}

		result := YourStruct{
			ID:          id,
			Name:        name,
			Website:     website,
			Latitude:    latitude,
			Longitude:   longitude,
			Description: description,
			Rating:      rating,
		}

		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return results, err
}
