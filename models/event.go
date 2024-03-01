package models

import (
	"errors"
	"fmt"
	"heitortanoue/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    string `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	// Save the event to the database
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // fecha se algo acontecer daqui para baixo

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?,
		description = ?,
		location = ?,
		dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	fmt.Println("event: ", e)
	defer stmt.Close() // fecha se algo acontecer daqui para baixo

	results, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	rowsAffected, _ := results.RowsAffected()
	fmt.Println("results: ", rowsAffected)
	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // fecha se algo acontecer daqui para baixo

	_, err = stmt.Exec(e.ID)
	return err
}

func (e Event) Register(userID int64) error {
	query := `
	INSERT INTO registrations (event_id, user_id)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // fecha se algo acontecer daqui para baixo

	_, err = stmt.Exec(e.ID, userID)
	if err != nil {
		return err
	}

	return err
}

func (e Event) Unregister(userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // fecha se algo acontecer daqui para baixo

	result, err := stmt.Exec(e.ID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("usuário não registrado para o evento")
	}

	return err
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	// como sabemos que o ID é unico, a função abaixo pega só 1 resultado
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// devolve todos os matches com a query
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetAllRegistrationsByUserId(userId int64) ([]Event, error) {
	query := `
	SELECT e.id, e.name, e.description, e.location, e.dateTime, e.user_id
	FROM events e
	INNER JOIN registrations r ON e.id = r.event_id
	WHERE r.user_id = ?
	`

	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
