package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"context"
)

type RepairStationPhoto struct {
	Id 				uuid.UUID
	RepairStationId uuid.UUID
	PhotoKey      	string
	S3Url			string
}

type RepairStationPhotoModel struct {
	DB 			*pgx.Conn
}

func (m *RepairStationPhotoModel) Get(ctx context.Context, stationId uuid.UUID) ([]RepairStationPhoto, error) {
	query := `SELECT
	repair_station_photo_id,
	repair_station_id,
	photo_key
	FROM repair_station_photo
	WHERE repair_station_id = $1`
	rows, err := m.DB.Query(ctx, query, stationId)
	if err != nil {
		return nil, err
	}

	var repairStationPhotos []RepairStationPhoto
	for rows.Next() {
		repairStationPhoto := RepairStationPhoto{}
		err := rows.Scan(&repairStationPhoto.Id, &repairStationPhoto.RepairStationId, &repairStationPhoto.PhotoKey)
		if err != nil {
			return nil, err
		}

		repairStationPhoto.S3Url = fmt.Sprintf("")
	}



	return repairStationPhotos, nil
}

///func (m *RepairStationModel) GetNearby(ctx context.Context, coordinates Point) ([]RepairStation, error) {
///	//// TODO here at least change it back to taking in the zoom.
///	//// We really don't even need to know the zoom to be honnest...
///	//// make the coordinates optaional in the parameters as well.
///	//	query := `SELECT
///	//		repair_station_id,
///	//		address_str,
///	//        ST_X(location::geometry) as lng,
///	//        ST_Y(location::geometry) as lat,
///	//		created_at
///	//	FROM repair_station
///	//	WHERE ST_DWithin(
///	//		location,
///	//		ST_MakePoint($1, $2)::geography,
///	//		5000  -- 5000 meters = 5km
///	//	)
///	//	`
///	query := `SELECT
///	repair_station_id,
///	address_str,
///	ST_X(location::geometry) as lng,
///	ST_Y(location::geometry) as lat,
///	created_at
///	FROM repair_station`
///
///
///	//	rows, err := m.DB.Query(ctx, query, coordinates.Lng, coordinates.Lat)
///	rows, err := m.DB.Query(ctx, query)
///	if err != nil {
///		return nil, err
///	}
///
///	var repairStations []RepairStation
///	for rows.Next() {
///		repairStation := RepairStation{}
///		err := rows.Scan(&repairStation.Id, &repairStation.AddressStr, &repairStation.Location.Lng, &repairStation.Location.Lat, &repairStation.CreatedAt)
///		if err != nil {
///			return nil, err
///		}
///		repairStations = append(repairStations, repairStation)
///	}
///
///	return repairStations, nil
///}
