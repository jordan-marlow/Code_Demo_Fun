package main

import (
	"log"
	"net/http"

	"web_service/api/credit_card_check"
	"web_service/api/patient"
)

func main() {
	http.HandleFunc("/api/cc_validate", credit_card_check.LuhnHandler)
	http.HandleFunc("/api/get_patients", patient.PatientHandler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
