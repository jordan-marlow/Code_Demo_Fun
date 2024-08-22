package patient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Num int `json:"num_patients"`
}

type Response struct {
	NameArray []string `json:"name_array"`
}

func PatientHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	patients := get_patient_name(req.Num)
	fmt.Println(patients)
	res := Response{NameArray: patients}
	json.NewEncoder(w).Encode(res)
}
