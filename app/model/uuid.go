package model

import "github.com/google/uuid"

type UUID string

func NewUUID() UUID {
	return UUID(uuid.New().String())
}
