package user

import (
	"encoding/json"
	f "fmt"
	"nexignal/ss/config"

	"github.com/lib/pq"
)

func (sif *SIForm) Get() (bool, ER) {

	// u for get user row
	var byte_row []byte
	var row map[string]interface{}
	var db_err bool = true

	// The order of inputs(columns) is important
	query := f.Sprintf("SELECT db.signin('%s', '%s')", sif.Identifier, sif.Password)

	err := db.QueryRow(query).Scan(&byte_row)

	json.Unmarshal(byte_row, &row)

	if err, ok := err.(*pq.Error); ok {
		f.Println("error: ", err)
		return db_err, CNC{err.Code, err.Constraint}.ErrorMessage()
	}

	return !db_err, ER{}
}

func GetRows() (bool, ER) {

	// u for get user row
	// var row map[string]interface{}
	var db_err bool = true

	// The order of inputs(columns) is important
	query := f.Sprintf("SELECT * from %s", config.USER_MODEL)

	rows, err := db.Query(query)

	if err, ok := err.(*pq.Error); ok {
		f.Println("error: ", err)
		return db_err, CNC{err.Code, err.Constraint}.ErrorMessage()
	}

	for rows.Next() {
		f.Println(rows.Scan())
	}
	return !db_err, ER{}
}
