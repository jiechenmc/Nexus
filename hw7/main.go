package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
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

	// Initialize Memcached client
	mc := memcache.New("localhost:11211") // Update this with your Memcached server address if necessary

	// Check cache for the player data
	cacheKey := "hw9_" + strings.Split(playerName, " ")[0] + strings.Split(playerName, " ")[1]
	fmt.Println("Cache Key", cacheKey)
	cachedItem, err := mc.Get(cacheKey)
	if err == nil {
		// Cache hit: decode the cached JSON data and return it
		fmt.Println("Cache Hit")
		var response Response
		if err := json.Unmarshal(cachedItem.Value, &response); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	fmt.Println("Cache Miss")
	// Cache miss: connect to the database
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

	// Cache the result in Memcached
	responseJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	mc.Set(&memcache.Item{
		Key:        cacheKey,
		Value:      responseJSON,
		Expiration: int32(60 * 5), // Cache expiration time in seconds (e.g., 5 minutes)
	})
	fmt.Println(cacheKey)
	fmt.Println(responseJSON)

	// Encode the response as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/hw9", hw9)
	fmt.Println("Server is listening on port 80...")
	http.ListenAndServe(":80", nil)
}
