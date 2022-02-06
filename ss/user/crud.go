package user

import (
	f "fmt"

	"nexignal/ss/config"

	"github.com/lib/pq"
)

var db = config.InitDB()

func (suf *SUForm) Create() (bool, ER) {
	var id int
	var db_err bool = true

	// The order of inputs(columns) is important
	query := f.Sprintf("SELECT db.signup('%s', '%s', '%s', '%s', '%s', '%s', %t)",
		suf.Name, suf.Lname, suf.Uname, suf.Email, suf.Password, suf.About, suf.Pub)

	err := db.QueryRow(query).Scan(&id)

	if err, ok := err.(*pq.Error); ok {
		return db_err, CNC{err.Code, err.Constraint}.ErrorMessage()
	}
	return !db_err, ER{}
}
