package main

import (
	"encoding/json"
	"time"
)

type UserData struct {
	Id     int                     `json:"id"`
	Events map[time.Time]EventData `json:"events"`
}

type EventData struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
type eventData struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func truncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func JsonToUser(jsonData []byte) (UserData, error) {
	var user UserData
	err := json.Unmarshal(jsonData, &user)
	return user, err
}

func JsonToEvent(jsonData []byte) (EventData, error) {
	var event eventData
	var event_date EventData
	err := json.Unmarshal(jsonData, &event)
	if err != nil {
		return event_date, err
	}
	event_date.Date, err = time.Parse("2006-01-02", event.Date)
	event_date.Date = truncateToDay(event_date.Date)
	if err != nil {
		return event_date, err
	}
	event_date.Name = event.Name
	return event_date, err
}
