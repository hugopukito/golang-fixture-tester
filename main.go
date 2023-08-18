package main

import (
	fixture "github.com/hugopukito/golang-fixture"
	"github.com/hugopukito/golang-fixture/database"
)

func main() {
	fixture.RunFixtures("my-fixtures", database.DatabaseParams{})
}
