package model

import "time"

type Shape struct {
	ID        string    `json:"shape_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Sides     string    `json:"sides"`
}

func (p *Shape) Archive() {
	// p.Archived = true
}
