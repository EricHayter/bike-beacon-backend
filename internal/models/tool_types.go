package models

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"context"
)

type ToolType string

type ToolTypeModel struct {
	DB *pgxpool.Pool
}

func (m *ToolTypeModel) Get(ctx context.Context) ([]ToolType, error) {
	query := `SELECT
	unnest(enum_range(NULL::tool_type))::text
	AS tool_type;`

	rows, err := m.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var toolTypes []ToolType
	for rows.Next() {
		var toolType ToolType
		err := rows.Scan(&toolType)
		if err != nil {
			return nil, err
		}
		toolTypes = append(toolTypes, toolType)
	}

	return toolTypes, nil
}

