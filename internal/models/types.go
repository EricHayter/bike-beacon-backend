package models

type Point struct {
	Lng float64 `json:"lng" db:"longitude"`
	Lat float64 `json:"lat" db:"latitude"`
}
