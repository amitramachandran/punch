package familytree

import (
	"context"
	"log"
)

func (e *Entity) CountSons(name string) int {
	query := "select count(*) from relationship where relatedAs = 'son' and (`relatedTo`) = (?)"
	row := e.c.QueryRowContext(context.Background(), query, name)

	var r int
	err := row.Scan(&r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func (e *Entity) CountDaughters(name string) int {
	query := "select count(*) from relationship where relatedAs = 'daughter' and (`relatedTo`) = (?)"
	row := e.c.QueryRowContext(context.Background(), query, name)

	var r int
	err := row.Scan(&r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func (e *Entity) CountWives(name string) int {
	query := "select count(*) from relationship where relatedAs = 'wife' and (`relatedTo`) = (?)"
	row := e.c.QueryRowContext(context.Background(), query, name)

	var r int
	err := row.Scan(&r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
