package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"os"
	"google.golang.org/appengine"
	"strings"
)

const WelcomeMessage = "Welcome to the climate registry. Please access the clima API via 'URL/clima/day' (day is an int)"
const DefaultDays = 3650

type Response struct {
	Climate string `json:"clima"`
	Day     int    `json:"dia"`
}

func main() {
	initializeRouter()
	appengine.Main()

	if verifyOfflineMode() {
		// Print simulation status per requirement
		days := DefaultDays
		sim := NewSimulation()
		sim.Simulate(days, NewSimluatorConfig(true, false))
	}
}
func initializeRouter() {
	mux := http.NewServeMux()

	// Expose REST api per bonus requirement
	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/clima/", climaHandle)
	mux.HandleFunc("/clima", climaHandle)
	mux.HandleFunc("/tasks/", tasksHandle)
	http.Handle("/", mux)
}

func tasksHandle(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Appengine-Cron") == "true" {
		taskParam := strings.TrimPrefix(r.URL.Path, "/tasks/")

		if taskParam == "initdb" {
			// Print simulation status per requirement
			days := DefaultDays
			sim := NewSimulation()
			_, error := sim.Simulate(days, NewSimluatorConfig(false, true))

			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Error :" + error.Error()))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("200 - Database was initialized successfully"))
			}
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 - Forbidden"))
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(WelcomeMessage)
}

func climaHandle(w http.ResponseWriter, r *http.Request) {
	dayParam := r.URL.Query().Get("dia")

	if dayParam == "" {
		dayParam = strings.TrimPrefix(r.URL.Path, "/clima/")
	}

	response := &Response{}
	dayIntParam := int64(0)
	dayIntParam, err := strconv.ParseInt(dayParam, 10, 64)

	if err == nil {
		var statusCode int
		response, statusCode = getClimateResponse(dayIntParam)
		w.WriteHeader(statusCode)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response = &Response{
			Climate: "invalid value in path " + r.URL.Path,
			Day:     -1,
		}
	}

	json.NewEncoder(w).Encode(response)
}

func getClimateResponse(dayIntParam int64) (*Response, int) {
	response := &Response{}
	var statusCode int
	days := int(dayIntParam)

	err, climate := GetClimate(int(days))

	if err == nil {
		statusCode = http.StatusOK
		response = &Response{
			Climate: climate,
			Day:     int(days),
		}
	} else {
		statusCode = http.StatusBadRequest
		response = &Response{
			Climate: "error: " + err.Error(),
			Day:     -1,
		}
	}

	return response, statusCode
}

func verifyOfflineMode() bool {
	offlineModeEnabled := false
	if len(os.Args) > 1 {
		offlineModeEnabled = os.Args[1] == "offline"
	}
	return offlineModeEnabled
}
