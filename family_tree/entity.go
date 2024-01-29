package familytree

import (
	"context"
	"database/sql"
	"log"

	"github.com/amitramachandran/dto"
)

type Entity struct {
	l *log.Logger
	c *sql.Conn
}

func NewEntity(l *log.Logger) *Entity {
	return &Entity{l, Connection()}
}

func (e *Entity) AddPerson(person string) {
	query := "Insert into `person` (`name`) values (?)"
	insertResult, err := e.c.ExecContext(context.Background(), query, person)
	if err != nil {
		log.Fatalf("impossible insert person: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)
}

func (e *Entity) AddRelation(relation string) {
	query := "Insert into `relations` (`relation_name`) values (?)"
	insertResult, err := e.c.ExecContext(context.Background(), query, relation)
	if err != nil {
		log.Fatalf("impossible insert relation: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)
}

func (e *Entity) GetPerson(person string) (string, error) {
	var personInfo []dto.Person
	query := "select * from `person` where (`name`) = (?)"
	rows, err := e.c.QueryContext(context.Background(), query, person)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var p dto.Person
		if err := rows.Scan(&p.Id, &p.Name); err != nil {
			return "", err
		}
		personInfo = append(personInfo, p)
	}
	// panic happening here while person not in DB
	return personInfo[0].Name, nil
}

func (e *Entity) GetRelation(relation string) (string, error) {
	var relationInfo []dto.Relation
	query := "select * from `relations` where (`relation_name`) = (?)"
	rows, err := e.c.QueryContext(context.Background(), query, relation)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var r dto.Relation
		if err := rows.Scan(&r.Id, &r.RelationName); err != nil {
			return "", err
		}
		relationInfo = append(relationInfo, r)
	}
	return relationInfo[0].RelationName, nil
}

func (e *Entity) CheckEntries(from string, to string, relation string) error {
	_, err := e.GetPerson(from)
	if err != nil {
		return err
	}
	_, err = e.GetPerson(to)
	if err != nil {
		return err
	}
	_, err = e.GetRelation(relation)
	if err != nil {
		return err
	}
	return nil

}

func (e *Entity) AddRelationship(from string, relation string, to string) {
	query := "Insert into `relationship` (`relatedFrom`, `relatedAs`, `relatedTo`) values (?, ?, ?)"
	insertResult, err := e.c.ExecContext(context.Background(), query, from, relation, to)
	if err != nil {
		log.Fatalf("impossible insert relationship: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)
}

func (e *Entity) GetFatherName(person string) (string, error) {
	var relationship []dto.Relationship
	query := `SELECT father.*
	FROM (
		  SELECT *
			FROM relationship
		   WHERE relatedFrom = (?)
		 ) father
   WHERE father.relatedAs = 'son' OR father.relatedAs = 'daughter'`

	// query := "select * from `relationship` where (`relatedFrom`) = (?)"
	rows, err := e.c.QueryContext(context.Background(), query, person)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var p dto.Relationship
		if err := rows.Scan(&p.Id, &p.RelatedFrom, &p.RelatedAs, &p.RelatedTo); err != nil {
			return "", err
		}
		relationship = append(relationship, p)
	}
	// panic happening here while person not in DB
	return relationship[0].RelatedTo, nil
}
