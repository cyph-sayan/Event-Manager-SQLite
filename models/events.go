package models

import (
	"events-management/database"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required`
	Description string    `binding:"required`
	Location    string    `binding:"required`
	DateTime    time.Time `binding:"required`
	UserId      int
}

func (e *Event) Save() error {
	query := ` INSERT INTO events (
		name, description, location, dateTime, userId
	) VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.Id = id
	return err
}

func GetEventById(id int) (Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := database.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		events = append(events, event)
	}
	return events, nil
}

func (event *Event) UpdateEvent() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return err
}

func (event *Event) DeleteEvent() error {
	query := `DELETE FROM events where id = ?`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Id)
	return err
}
