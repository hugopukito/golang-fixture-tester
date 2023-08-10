package main

import (
	fixture "github.com/hugopukito/golang-fixture"
)

func main() {
	fixture.RunFixtures("db-fixture-test", "my-fixtures")
}
