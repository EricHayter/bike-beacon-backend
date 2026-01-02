package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
//	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
//	"github.com/aws/smithy-go"
)

type RepairStationPhoto struct {
	Id 				uuid.UUID
	RepairStationId uuid.UUID
	PhotoKey      	string
}

type RepairStationPhotoModel struct {
	DB 			*pgx.Conn
	S3Client 	*s3.Client
}

func (m *RepairStationPhotoModel) Get(ctx context.Context, stationId uuid.UUID) (*RepairStationPhoto, error) {
	query := `SELECT
	repair_station_photo_id,
	repair_station_id,
	photo_key
	FROM repair_station_photo
	WHERE repair_station_id = $1`
	row := m.DB.QueryRow(ctx, query, stationId)

	repairStationPhoto := RepairStationPhoto{}
	err := row.Scan(&repairStationPhoto.Id, &repairStationPhoto.RepairStationId, &repairStationPhoto.PhotoKey)
	if err != nil {
		return nil, err
	}

	// pull the data now from the S3 storage
	result, err := m.S3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String("iamges"),
		Key:    aws.String(repairStationPhoto.PhotoKey),
	})
	if err != nil {
		var noKey *types.NoSuchKey
		if errors.As(err, &noKey) {
			err = noKey
		}
		return nil, err
	}
	defer result.Body.Close()

	// pass in the data into the struct?
//	result.Body.Read()

	return &repairStationPhoto, nil
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
