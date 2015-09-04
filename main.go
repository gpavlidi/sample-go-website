package main

import (
	"log"

	//"database/sql"
	//_ "github.com/lib/pq"
)

func main() {
	/*
		db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/goboostrap?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		return
	*/
	app := NewApp(&AppConfig{":8080", "web1"})
	log.Println(app.Config.Addr)
	app.ListenAndServe()
}
