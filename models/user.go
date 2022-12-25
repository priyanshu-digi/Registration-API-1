package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique" ` //username should be unique
	Email    string `json:"email" gorm:"unique" `    //email should be unique
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

/*func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if u.Name == "" {
			return errors.New("required name parameter")
		}
		if u.Username == "" {
			return errors.New("required username parameter")
		}
		if u.Password == "" {
			return errors.New("required Password parameter")
		}
		if u.Email == "" {
			return errors.New("required Email parameter")
		}
		// validation of email parameter
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil
	}

}*/
