package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"context"
	"time"
)

type RepairStation struct {
	Id 				uuid.UUID
	AddressStr      string
	Location     	Point
	CreatedAt 		time.Time
}

type RepairStationModel struct {
	DB *pgx.Conn
}

func (m *RepairStationModel) Get(ctx context.Context, id uuid.UUID) (*RepairStation, error) {
	query := `SELECT
	repair_station_id,
	address_str,
	ST_X(location::geometry) as lng,
	ST_Y(location::geometry) as lat,
	created_at
	FROM repair_station
	WHERE repair_station_id = $1`

	row := m.DB.QueryRow(ctx, query, id)

	repairStation := RepairStation{}
	err := row.Scan(&repairStation.Id, &repairStation.AddressStr, &repairStation.Location.Lng, &repairStation.Location.Lat, &repairStation.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &repairStation, nil
}

func (m *RepairStationModel) GetNearby(ctx context.Context, coordinates Point) ([]RepairStation, error) {
	//// TODO here at least change it back to taking in the zoom.
	//// We really don't even need to know the zoom to be honnest...
	//// make the coordinates optaional in the parameters as well.
	//	query := `SELECT
	//		repair_station_id,
	//		address_str,
	//        ST_X(location::geometry) as lng,
	//        ST_Y(location::geometry) as lat,
	//		created_at
	//	FROM repair_station
	//	WHERE ST_DWithin(
	//		location,
	//		ST_MakePoint($1, $2)::geography,
	//		5000  -- 5000 meters = 5km
	//	)
	//	`
	query := `SELECT
	repair_station_id,
	address_str,
	ST_X(location::geometry) as lng,
	ST_Y(location::geometry) as lat,
	created_at
	FROM repair_station`


	//	rows, err := m.DB.Query(ctx, query, coordinates.Lng, coordinates.Lat)
	rows, err := m.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var repairStations []RepairStation
	for rows.Next() {
		repairStation := RepairStation{}
		err := rows.Scan(&repairStation.Id, &repairStation.AddressStr, &repairStation.Location.Lng, &repairStation.Location.Lat, &repairStation.CreatedAt)
		if err != nil {
			return nil, err
		}
		repairStations = append(repairStations, repairStation)
	}

	return repairStations, nil
}
