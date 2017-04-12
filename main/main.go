package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"os"
	"google.golang.org/appengine"
	"strings"
	"alpha-quad-sim/simulator"
	"alpha-quad-sim/database"
)

const WelcomeMessage = "Welcome to the climate registry. Please access the clima API via 'URL/clima/day' (day is an int)"
const DefaultDays = 3650

// Response - response object
type Response struct {
	Climate string `json:"clima"`
	Day     int    `json:"dia"`
}

func main() {
	if verifyOfflineMode() {
		// Print simulation status per requirement
		days := DefaultDays
		sim := simulator.NewSimulation()
		sim.Simulate(days, simulator.NewSimluatorConfig(true, false))
	}

	initializeRouter()
	appengine.Main()
}

// initializeRouter - Router initialization
func initializeRouter() {
	mux := http.NewServeMux()

	// Expose REST api per bonus requirement
	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/clima/", climaHandle)
	mux.HandleFunc("/clima", climaHandle)
	mux.HandleFunc("/tasks/", tasksHandle)
	http.Handle("/", mux)

	if verifyOfflineMode() {
		http.ListenAndServe(":3000", mux)
	}
}

// tasksHandle - Handler of task action requests
func tasksHandle(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Appengine-Cron") == "true" {
		taskParam := strings.TrimPrefix(r.URL.Path, "/tasks/")

		if taskParam == "initdb" {
			// Print simulation status per requirement
			days := DefaultDays
			sim := simulator.NewSimulation()
			_, error := sim.Simulate(days, simulator.NewSimluatorConfig(false, true))

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

// indexHandle - Handler of index requests
func indexHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(WelcomeMessage)
}

// climaHandle - Handler of /clima/ action requests
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

// getClimateResponse - Obtain climate object from the database and generate the response
func getClimateResponse(dayIntParam int64) (*Response, int) {
	response := &Response{}
	var statusCode int
	days := int(dayIntParam)

	err, climate := database.GetClimate(int(days))

	if err == nil {
		statusCode = http.StatusOK
		response = &Response{
			Climate: climate,
			Day:     int(days),
		}
	} else {
		statusCode = http.StatusInternalServerError
		response = &Response{
			Climate: "error: " + err.Error(),
			Day:     -1,
		}
	}

	return response, statusCode
}

// verifyOfflineMode - Verifies if the application has been launched in the offline mode
func verifyOfflineMode() bool {
	offlineModeEnabled := false
	if len(os.Args) > 1 {
		offlineModeEnabled = os.Args[1] == "offline"
	}
	return offlineModeEnabled
}
