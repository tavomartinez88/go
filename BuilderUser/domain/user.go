//domain is a package that contains all types representant a object real life
package domain

import (
	"github.com/go-errors/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// User is a type that represent a Person
type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}

func (user *User) ValidateEmail() (bool, error){
	if len(user.Email) == 0 {
		return false, errors.New("email should not be empty")
	} else {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		return re.MatchString(user.Email), nil
	}
}

func (user *User) Encode() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password=string(hashedPassword)
	return err
}