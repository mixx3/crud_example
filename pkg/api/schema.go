package api

import "time"

type LogPostSchema struct {
	Message string `json:"message" example:"Error!!"`
}

type LogGetSchema struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	DtCreated time.Time `json:"dtCreated"`
}

type LogGetAllSchema struct {
	Items []LogGetSchema `json:"items"`
}
