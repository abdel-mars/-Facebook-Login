package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var votes = make(map[string]int)

func HandleVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	pet := r.FormValue("pet")
	if pet == "" {
		http.Error(w, "Please select a pet", http.StatusBadRequest)
		return
	}

	votes[pet]++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": fmt.Sprintf("Vote for %s recorded (simulation)", pet),
		"votes":   votes,
	})
}

func main() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	http.HandleFunc("/api/vote", HandleVote)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil) // Fixed the typo here!
}