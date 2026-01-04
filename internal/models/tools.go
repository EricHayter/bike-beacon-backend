package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"context"
)

type Tool struct {
	Id 				uuid.UUID
	RepairStationId uuid.UUID
	Type 			string
}

type ToolModel struct {
	DB *pgxpool.Pool
}

func (m *ToolModel) Get(ctx context.Context, stationId uuid.UUID) ([]Tool, error) {
	query := `SELECT
	tool_id,
	repair_station_id,
	tool_type_id
	FROM tool
	WHERE repair_station_id = $1`

	rows, err := m.DB.Query(ctx, query, stationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tools []Tool
	for rows.Next() {
		tool := Tool{}
		err := rows.Scan(&tool.Id, &tool.RepairStationId, &tool.Type)
		if err != nil {
			return nil, err
		}
		tools = append(tools, tool)
	}

	return tools, nil
}

