package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "mypropertyqr-landsurvey/Algs"
	"mypropertyqr-landsurvey/Events"
)

func main() {
	// Algs.InitPy()


	http.HandleFunc("/extractdata", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
	
		id, _ := body["id"].(string)
		memberId, _ := body["memberId"].(string)
	
		dataStr := Events.Extractdata(id, memberId)
		var data map[string]interface{}
		err = json.Unmarshal([]byte(dataStr), &data)
		if err != nil {
			http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("Server started at :5003")
	err := http.ListenAndServe(":5003", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}