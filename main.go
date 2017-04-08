package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"os"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"strings"
)

const WelcomeMessage = "Welcome to the climate registry. Please access the clima API via 'URL/clima/day' (day is an int)"

type Response struct {
	Climate string `json:"clima"`
	Day     int    `json:"dia"`
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/clima/", climaHandle)
	appengine.Main()

	if verifyOfflineMode() {
		// Print simulation status per requirement
		days := 3650
		sim := NewSimulation()
		sim.Simulate(days, NewSimluatorConfig(true, false))

		// Expose REST api
		router := mux.NewRouter()
		router.HandleFunc("/", IndexEndpoint).Methods("GET")
		router.HandleFunc("/clima/{dia:[0-9]+}", GetClimateEndpoint).Methods("GET")
		router.Queries("{dia:[0-9]+}")

		return
	}

	initDb, days := verifyInitializeDbMode()
	if initDb {
		// Persist simulation status per bonus requirement
		sim := NewSimulation()
		sim.Simulate(days, NewSimluatorConfig(false, true))
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
	dayParam := strings.TrimPrefix(r.URL.Path, "/clima/")
	dayIntParam := int64(0)
	dayIntParam, err := strconv.ParseInt(dayParam, 10, 64)

	response := getClimateResponse(err, dayIntParam)
	json.NewEncoder(w).Encode(response)
}

func IndexEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(WelcomeMessage)
	return
}

func GetClimateEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	days, err := strconv.ParseInt(params["dia"], 10, 64)
	response := &Response{}

	if err == nil {
		sim := NewSimulation()
		climate := sim.Simulate(int(days), NewSimluatorConfig(false, false))

		response = &Response{
			Climate: climate,
			Day:     int(days),
		}

	} else {
		response = &Response{
			Climate: "invalid value",
			Day:     -1,
		}
	}


	json.NewEncoder(w).Encode(response)
}

func getClimateResponse(err error, dayIntParam int64) *Response {
	response := &Response{}

	if err == nil {
		days := int(dayIntParam)

		sim := NewSimulation()
		climate := sim.Simulate(int(days), NewSimluatorConfig(false, false))

		response = &Response{
			Climate: climate,
			Day:     int(days),
		}
	} else {
		response = &Response{
			Climate: "invalid value",
			Day:     -1,
		}
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
