package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func hw9(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-CSE356", "66cfe2a89ba30e1a6c706756")

	db, err := sql.Open("mysql", "root:example@/hw9")
	if err != nil {
		panic(err)
	}

	query := "select A.Player as p1,B.Player as p2,C.Player as p3,D.Player as p4 from assists A, assists B, assists C, assists D where A.POS=B.POS and B.POS=C.POS and C.POS=D.POS and A.Club<>B.Club and A.club<>C.Club and A.Club<>C.Club and A.Club<>D.Club and B.Club<>C.Club and B.Club<>D.Club and C.Club<>D.Club and A.Player='PLAYER_NAME_HERE' order by A.A+B.A+C.A+D.A desc, A.A desc, B.A desc, C.A desc, D.A desc, p1, p2, p3, p4 limit 1;"
	res, err := db.Exec(query)

	if err != nil {
		fmt.Println(res)
	}

	fmt.Fprintf(w, "hello\n")
}

func main() {

	http.HandleFunc("/hw9", hw9)

	http.ListenAndServe(":8080", nil)
}
