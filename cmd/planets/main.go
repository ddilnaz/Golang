
package main

import (
	"fmt"
	"log"
	"net/http"
	"TSIS1/pkg/planet"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/planets", planet.GetPlanetsHandler).Methods("GET")
	r.HandleFunc("/planets/{name}", planet.GetPlanetByNameHandler).Methods("GET")
	r.HandleFunc("/health", planet.HealthCheckHandler).Methods("GET")

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
