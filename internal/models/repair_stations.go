package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type RepairStation struct {
	RepairStationId uuid.UUID
	Coordinates     Point
	Address         string
}

type RepairStationModel struct {
	DB *pgx.Conn
}

func (m *RepairStationModel) Get(id uuid.UUID) (RepairStation, error) {
	return RepairStation{}, nil
}

func (m *RepairStationModel) GetNearby(coordinates Point) ([]RepairStation, error) {
	return []RepairStation{}, nil
}
