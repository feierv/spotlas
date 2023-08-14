SELECT
    name,
    website,
    ST_X(coordinates::geometry) AS longitude,
    ST_Y(coordinates::geometry) AS latitude
FROM "MY_TABLE"
WHERE ST_Within(ST_SetSRID(ST_MakePoint(ST_X(coordinates::geometry), ST_Y(coordinates::geometry)), 4326), ST_MakeEnvelope(-122.4194 - 0.01, 37.7749 - 0.01, -122.4194 + 0.01, 37.7749 + 0.01, 4326))
ORDER BY ST_Distance(ST_SetSRID(ST_MakePoint(ST_X(coordinates::geometry), ST_Y(coordinates::geometry)), 4326), coordinates),
         CASE WHEN ST_Distance(ST_SetSRID(ST_MakePoint(ST_X(coordinates::geometry), ST_Y(coordinates::geometry)), 4326), coordinates) < 50 THEN rating ELSE 0 END DESC;
