package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	Players []string `json:"players"`
}

func hw9(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-CSE356", "66cfe2a89ba30e1a6c706756") // Add X-CSE356 header with submission ID

	// Get player name from query parameters
	playerName := req.URL.Query().Get("player")
	if playerName == "" {
		http.Error(w, "Player name is required", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/hw9")
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Define the query
	query := `
        SELECT A.Player AS p1, B.Player AS p2, C.Player AS p3, D.Player AS p4
        FROM assists A, assists B, assists C, assists D
        WHERE A.POS = B.POS 
          AND B.POS = C.POS 
          AND C.POS = D.POS 
          AND A.Club <> B.Club 
          AND A.Club <> C.Club 
          AND A.Club <> D.Club 
          AND B.Club <> C.Club 
          AND B.Club <> D.Club 
          AND C.Club <> D.Club 
          AND A.Player = ?
        ORDER BY A.A + B.A + C.A + D.A DESC, A.A DESC, B.A DESC, C.A DESC, D.A DESC, p1, p2, p3, p4
        LIMIT 1;
    `

	// Execute the query
	rows, err := db.Query(query, playerName)
	if err != nil {
		http.Error(w, "Query execution failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process the results
	var response Response
	if rows.Next() {
		var p1, p2, p3, p4 string
		if err := rows.Scan(&p1, &p2, &p3, &p4); err != nil {
			http.Error(w, "Error reading query result", http.StatusInternalServerError)
			return
		}
		response.Players = []string{p1, p2, p3, p4}
	} else {
		response.Players = []string{} // No results, return an empty array
	}

	// Encode the response as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/hw9", hw9)

	http.ListenAndServe(":8080", nil)
}
