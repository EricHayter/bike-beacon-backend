package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"context"
)

type RepairStationPhoto struct {
	Id 				uuid.UUID
	PhotoKey      	string
}

type RepairStationPhotoModel struct {
	DB 			*pgx.Conn
}

func (m *RepairStationPhotoModel) Get(ctx context.Context, stationId uuid.UUID) ([]RepairStationPhoto, error) {
	query := `SELECT
	repair_station_photo_id,
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
		err := rows.Scan(&repairStationPhoto.Id, &repairStationPhoto.PhotoKey)
		if err != nil {
			return nil, err
		}

		repairStationPhotos = append(repairStationPhotos, repairStationPhoto)
	}

	return repairStationPhotos, nil
}
