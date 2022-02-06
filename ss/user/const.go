package user

import (
	"net/http"

	"github.com/lib/pq"
)

// Error Code & Constraint
type CNC struct {
	code       pq.ErrorCode
	constraint string
}

// Error Response
type ER struct {
	code   int
	header string
}

// Corresponding message and error code for db errors
var ECNC = map[CNC]ER{
	{"23505", "user_email_key"}: {http.StatusConflict, "آدرس ایمیل تکراری است"},
	{"23505", "user_uname_key"}: {http.StatusConflict, "نام کاربری تکراری است"},
	{"23514", "e_domain_check"}: {http.StatusBadRequest, "ایمیل نامعتبر"},
}

func (c CNC) ErrorMessage() ER { return ECNC[c] }
