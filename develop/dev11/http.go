package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func UIDPage(w http.ResponseWriter, r *http.Request) {
	var errorString string
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			errorString = fmt.Sprintf(`"error": %v`, err)
			http.Error(w, errorString, http.StatusBadRequest)
			return
		}
		var body []byte
		body, err = io.ReadAll(r.Body)
		if err != nil {
			errorString = fmt.Sprintf(`"error": %v`, err)
			http.Error(w, errorString, http.StatusServiceUnavailable)
			return
		}
		switch r.URL.Path {
		case "/create_event":
			data, err := JsonToEvent(body)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			createEvent(user_id, data)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			fmt.Fprint(w, `"result": "success"`)
		case "/update_event":
			data, err := JsonToEvent(body)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			updateEvent(user_id, data)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			fmt.Fprint(w, `"result": "success"`)
		case "/delete_event":
			data, err := JsonToEvent(body)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			deleteEvent(user_id, data.Date)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			fmt.Fprint(w, `"result": "success"`)
		default:
			http.Error(w, "Unknown function", http.StatusBadRequest)
		}
	case "GET":
		var response []byte
		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			errorString = fmt.Sprintf(`"error": %v`, err)
			http.Error(w, errorString, http.StatusBadRequest)
			return
		}
		switch r.URL.Path {
		case "/events_for_day":
			events, err := eventsForTime(user_id, 0, 1)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			response, err = json.Marshal(events)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			fmt.Fprint(w, `"result": `)
			fmt.Fprint(w, string(response))
		case "/events_for_week":
			events, err := eventsForTime(user_id, 0, 7)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			response, err = json.Marshal(events)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			fmt.Fprint(w, `"result": `)
			fmt.Fprint(w, string(response))
		case "/events_for_month":
			events, err := eventsForTime(user_id, 1, 0)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			response, err = json.Marshal(events)
			if err != nil {
				errorString = fmt.Sprintf(`"error": %v`, err)
				http.Error(w, errorString, http.StatusServiceUnavailable)
				return
			}
			fmt.Fprint(w, `"result": `)
			fmt.Fprint(w, string(response))
		default:
			http.Error(w, "Unknown function", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Unsupported method", http.StatusBadRequest)
	}
}

func logRequest(handler http.Handler, file io.Writer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(file, "%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func StartHttp(port string, filename string) error {
	f, _ := os.Create(filename)
	http.HandleFunc("/", UIDPage)

	if err := http.ListenAndServe(port, logRequest(http.DefaultServeMux, f)); err != http.ErrServerClosed {
		return err
	}
	return nil
}
