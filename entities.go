package main

import (
	"time"

	"github.com/google/uuid"
)

type Dog struct {
	ID      uuid.UUID
	Name    string
	Breed   string
	Size    int
	Created time.Time
}

type Cat struct {
	ID         uuid.UUID
	Name       string
	Color      string
	TailLength int
	Info       Info
	Moree      More
}

type Info struct {
	ID   uuid.UUID
	Text string
}

type More struct {
	ID       uuid.UUID
	MoreText string
}

type User struct {
	Email    string
	Password string
}
