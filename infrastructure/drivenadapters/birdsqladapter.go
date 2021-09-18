package drivenadapters

import (
	"LealTechTest/domain/entities"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "onlydemopassword"
	dbname   = "birddemo"
)

var db *sql.DB
var err error


func init() {
	psqlConnInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,port,user,password,dbname)
	db, err = sql.Open("postgres", psqlConnInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func List() []entities.Bird{

    birds := []entities.Bird{}
	rows, err := db.Query("select  * from birdinfo.birds")
	if err != nil {
		fmt.Println("Error can't obtain rows of response with list of birds, ", err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var bt entities.Bird
		 err = rows.Scan(&bt.ID, &bt.Specie, &bt.Name, &bt.Characteristics)
		birds = append(birds, bt)
	}

	return birds
}
