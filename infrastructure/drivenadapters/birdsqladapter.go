package birdsqladapter

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

type  birdSqlDrivenAdapter struct {
	 db *sql.DB

}


var Dba birdSqlDrivenAdapter

func init() {

	psqlConnInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	Dba.db, err = sql.Open("postgres", psqlConnInfo)

	if err != nil {
		panic(err)
	}

	err = Dba.db.Ping()
	if err != nil {
		panic(err)
	}
}

func (bda birdSqlDrivenAdapter) Add(b entities.Bird) error {

	sql := ` INSERT INTO birdinfo.birds (specie,name,characteristics) VALUES ($1, $2, $3) `
	_, err := bda.db.Exec(sql, b.Specie, b.Name, b.Characteristics)
	if err != nil {
		fmt.Println("Error on insert of new bird", err)
		return err
	}
	return nil
}

func (bda birdSqlDrivenAdapter) Update(b entities.Bird)  error {

	sql := `UPDATE birdinfo.birds SET specie = $2 ,name = $3 , characteristics = $4
		 WHERE id = $1`
	_, err := Dba.db.Exec(sql,b.ID, b.Specie,b.Name,b.Characteristics )

	if err != nil {
		panic(err)
		return err
	}

	return  nil
}

func (bda birdSqlDrivenAdapter) Delete(bird entities.Bird) error{

	sql := `DELETE  FROM birdinfo.birds b WHERE b.id =$1`
	_, err := bda.db.Exec(sql,bird.ID)
	if err != nil {
		fmt.Println("Error on delete record ", err)
		return err
	}
	return nil
}

func (bda birdSqlDrivenAdapter) FindById(id int) (entities.Bird,error) {

	bird := entities.Bird{}
	sql := `SELECT * FROM birdinfo.birds b WHERE b.id =$1`
	rows, err := bda.db.Query(sql, id)
	if err != nil {
		return bird, err
	}
	defer rows.Close()
	rows.Next()

	err = rows.Scan(&bird.ID, &bird.Specie, &bird.Name, &bird.Characteristics)
	if err != nil {
		return bird, err
	}
	return bird,nil
}


func (bda birdSqlDrivenAdapter) List() ( []entities.Bird,  error ) {

	birds := []entities.Bird{}
	rows, err := bda.db.Query("select  * from birdinfo.birds")
	if err != nil {
		fmt.Println("Error can't obtain rows of response with list of birds, ", err)
		return birds,err
	}
	defer rows.Close()

	for rows.Next() {
		var bt entities.Bird
		err = rows.Scan(&bt.ID, &bt.Specie, &bt.Name, &bt.Characteristics)
		birds = append(birds, bt)
	}
	return birds, nil
}
