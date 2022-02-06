package user

import "net/mail"

const (
	USER_NAME_LENGTH = "محدودیت مشخصات: حداقل ۳ و حداکثر ۳۰ حرف"
	PARSE_ADDRESS    = "ایمیل نامعتبر"
	PASSWORD_LENGTH  = "محدودیت گذرواژه: حداقل ۸ کاراکتر"
)

type SUForm struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Lname    string `form:"lname" json:"lname" binding:"required"`
	Uname    string `form:"uname" json:"uname" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	About    string `form:"about" json:"about" binding:"required" `
	Pub      bool   `form:"pub" json:"pub" binding:"required"`
}

type SIForm struct {
	Identifier string `form:"identifier" json:"identifier" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
}

func (su *SUForm) Validate() (bool, string) {
	var err bool = true
	if len(su.Name) < 3 || len(su.Name) > 30 {
		return err, USER_NAME_LENGTH
	} else if len(su.Lname) < 3 || len(su.Lname) > 30 {
		return err, USER_NAME_LENGTH
	} else if len(su.Uname) < 3 || len(su.Uname) > 30 {
		return err, USER_NAME_LENGTH
	} else if _, e := mail.ParseAddress(su.Email); e != nil {
		return err, PARSE_ADDRESS
	} else if len(su.Password) < 1 {
		return err, PASSWORD_LENGTH
	}
	return false, ""
}

func (si *SIForm) Validate() (bool, string) {
	var err bool = true
	if len(si.Identifier) < 3 {
		return err, USER_NAME_LENGTH
	} else if len(si.Password) < 1 {
		return err, PASSWORD_LENGTH
	}
	return false, ""
}
