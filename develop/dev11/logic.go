package main

import (
	"errors"
	"sync"
	"time"
)

var mutex sync.RWMutex
var cache map[int]UserData

func GetUser(uid int) (UserData, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	val, ok := cache[uid]
	if ok {
		return val, true
	}
	return val, false
}

func createEvent(userId int, event EventData) error {
	mutex.Lock()
	defer mutex.Unlock()
	val, ok := cache[userId]
	if !ok {
		user := UserData{
			Id:     userId,
			Events: map[time.Time]EventData{},
		}
		user.Events[event.Date] = event
		cache[userId] = user
		return nil
	}
	_, ok = val.Events[event.Date]
	if ok {
		return errors.New("event already exist. Use update function")
	}
	val.Events[event.Date] = event
	return nil
}

func updateEvent(userId int, event EventData) error {
	mutex.Lock()
	defer mutex.Unlock()
	val, ok := cache[userId]
	if !ok {
		return errors.New("unknown user")
	}
	_, ok = val.Events[event.Date]
	if !ok {
		return errors.New("unknown event")
	}
	val.Events[event.Date] = event
	return nil
}

func deleteEvent(userId int, date time.Time) error {
	mutex.Lock()
	defer mutex.Unlock()
	val, ok := cache[userId]
	if !ok {
		return errors.New("unknown user")
	}
	_, ok = val.Events[date]
	if !ok {
		return errors.New("unknown event")
	}
	delete(val.Events, date)
	return nil
}

func eventsForTime(userId int, months int, days int) ([]EventData, error) {
	var events []EventData
	mutex.RLock()
	defer mutex.RUnlock()
	user, ok := cache[userId]
	if !ok {
		return nil, errors.New("unknown user")
	}

	start := truncateToDay(time.Now())
	end := start.AddDate(0, months, days)
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		event, ok := user.Events[d]
		if ok {
			events = append(events, event)
		}
	}
	return events, nil
}
