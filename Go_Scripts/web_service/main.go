package main

import (
	"log"
	"net/http"

	"html/template"
	"path/filepath"
	"web_service/api/credit_card_check"
	"web_service/api/patient"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/cc_validate", credit_card_check.LuhnHandler)
	http.HandleFunc("/api/get_patients", patient.PatientHandler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
