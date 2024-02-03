
package planet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Planet struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Diameter int    `json:"diameter"`
	Distance int    `json:"distance"`
}

var planets = []Planet{
	{"Mercury", "Rocky", 4879, 57910000},
	{"Venus", "Rocky", 12104, 108200000},
	{"Earth", "Rocky", 12742, 149600000},
	{"Mars", "Rocky", 6779, 227900000},
	{"Jupiter", "Gas Giant", 139820, 778500000},
	{"Saturn", "Gas Giant", 116460, 1433000000},
	{"Uranus", "Ice Giant", 50724, 2871000000},
	{"Neptune", "Ice Giant", 49244, 4495000000},
}

func GetPlanetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(planets)
}

func GetPlanetByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["name"]

	var foundPlanet Planet
	for _, p := range planets {
		if p.Name == name {
			foundPlanet = p
			break
		}
	}

	if foundPlanet.Name != "" {
		json.NewEncoder(w).Encode(foundPlanet)
	} else {
		http.NotFound(w, r)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Planets App is healthy!\nAuthor: Dilnaz")
}
