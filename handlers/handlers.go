package handlers

import (
	"encoding/json"
	"net/http"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []models.Job

	jobs = append(jobs, models.Job{ID: 1, Name: "Accounting"})
	jobs = append(jobs, models.Job{ID: 2, Name: "Programming"})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
