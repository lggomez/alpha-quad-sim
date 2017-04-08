package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"os"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"strings"
)

const WelcomeMessage = "Welcome to the climate registry. Please access the clima API via 'URL/clima/day' (day is an int)"

type Response struct {
	Path string
	Climate string `json:"clima"`
	Day     int    `json:"dia"`
}

func main() {
	mux := http.NewServeMux()

	// Expose REST api per bonus requirement
	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/clima/", climaHandle)
	mux.HandleFunc("/clima", climaHandle)
	http.Handle("/", mux)
	appengine.Main()

	initDb, days := verifyInitializeDbMode()
	if initDb {
		// Persist simulation status per bonus requirement
		sim := NewSimulation()
		sim.Simulate(days, NewSimluatorConfig(false, true))
	}

	if verifyOfflineMode() {
		// Print simulation status per requirement
		days := 3650
		sim := NewSimulation()
		sim.Simulate(days, NewSimluatorConfig(true, false))
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Index hit")
	json.NewEncoder(w).Encode(WelcomeMessage)
}

func climaHandle(w http.ResponseWriter, r *http.Request) {
	dayParam := r.URL.Query().Get("dia")

	if dayParam == ""  {
		dayParam = strings.TrimPrefix(r.URL.Path, "/clima/")
	}

	response := &Response{}
	dayIntParam := int64(0)
	dayIntParam, err := strconv.ParseInt(dayParam, 10, 64)

	if err != nil {
		response = &Response{
			Climate: "invalid value in path " + r.URL.Path,
			Day:     -1,
		}
	} else {
		response = getClimateResponse(dayIntParam)
	}

	json.NewEncoder(w).Encode(response)
}

func getClimateResponse(dayIntParam int64) *Response {
	response := &Response{}

	days := int(dayIntParam)

	sim := NewSimulation()
	climate := sim.Simulate(int(days), NewSimluatorConfig(false, false))

	response = &Response{
		Climate: climate,
		Day:     int(days),
	}

	return response
}

func verifyOfflineMode() bool {
	offlineModeEnabled := false
	if len(os.Args) > 1 {
		offlineModeEnabled = os.Args[1] == "offline"
	}
	return offlineModeEnabled
}

func verifyInitializeDbMode() (bool, int) {
	initDb := false
	days := int64(0)
	if len(os.Args) > 2 {
		initDb = os.Args[1] == "initdb"
		days, _ = strconv.ParseInt(os.Args[2], 10, 64)
	}
	return initDb, int(days)
}
