package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"unico/app/helper/io"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
}

func Conn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./unico.db")
	if err != nil {
		panic(err)
	}
	x := sqlite{
		db: db,
	}
	x.init("./dump.sql")
	x.load("./DEINFO_AB_FEIRASLIVRES_2014.csv")
	return db
}

func (s *sqlite) init(filename string) {
	// Read file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Execute all
	_, err = s.db.Exec(string(content))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *sqlite) load(filename string) {
	records := io.ReadCSV(filename)
	into := "INSERT INTO feiraslivres(id, lng, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome, registro, logradouro, numero, bairro, referencia)"
	values := "\nVALUES"
	for i, line := range records {
		values = values + "\n("
		for j := 0; j < len(line); j++ {
			if j == 0 {
				values = values + line[j]
			} else {
				values = values + `"` + line[j] + `"`
			}
			if j < len(line)-1 {
				values = values + ","
			}
		}
		values = values + ")"
		if i < len(records)-1 {
			values = values + ","
		}
	}
	// loadfile := "./load.sql"
	// io.Writefile(loadfile, data)
	_, err := s.db.Exec(into + values)
	if err != nil {
		fmt.Println(err.Error())
	}
}
