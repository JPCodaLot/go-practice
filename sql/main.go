package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

var dsnFlag = flag.String("d", "", `Database DSN (Data Source Name)
Format:
	[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
Example:
	username:password@protocol(address)/dbname?param=value
`)
var hostFlag = flag.String("h", "localhost:3000", `Host`)

func main() {
	flag.Parse()
	db, err := sql.Open("mysql", *dsnFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		var search string
		path := strings.NewReader(r.URL.Path)
		_, err := fmt.Fscanf(path, "/api/creators/%s", &search)
		if err != nil {
			log.Print(err)
		}
		creators, err := getCreators(db, search)
		if err != nil {
			log.Print(err)
			return
		}
		data, err := json.Marshal(creators)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%s", data)
	})

	http.Handle("/app/", http.FileServer(http.Dir("frontend")))

	log.Fatal(http.ListenAndServe(*hostFlag, nil))
}

type Creator struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
}

func getCreators(db *sql.DB, search string) ([]Creator, error) {
	var creators []Creator
	rows, err := db.Query("SELECT id, cname, fname, lname FROM `Creators` WHERE id = ?", search)
	if err != nil {
		log.Fatalf("getCreators: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var creator Creator
		err := rows.Scan(&creator.ID, &creator.Username, &creator.FirstName, &creator.LastName)
		if err != nil {
			log.Printf("getCreators: %s", err)
		}
		creators = append(creators, creator)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("getCreators: %s", err)
	}
	return creators, nil
}
