package domain

import "testing"

func TestUser_ValidateEmail_When_Is_Valid(t *testing.T) {
	user := &User{
		FirstName: "Gustavo",
		LastName:  "Martinez",
		Email:     "ggg@gmail.com",
		Password:  "Germ1=&",
		Phone:     "123456",
	}

	isValidEmail, _ := user.ValidateEmail()
	if !isValidEmail || len(user.Email) == 0 {
		t.Error("Expected", true, "Got", false)
	}
}

func TestUser_ValidateEmail_When_Not_Is_Valid(t *testing.T) {
	user := &User{
		FirstName: "Gustavo",
		LastName:  "Martinez",
		Email:     "ggg_gmail.com",
		Password:  "Germ1=&",
		Phone:     "123456",
	}

	isValidEmail, _ := user.ValidateEmail()
	if isValidEmail || len(user.Email) == 0 {
		t.Error("Expected", false, "Got", true)
	}
}

func TestUser_ValidateEmail_When_Not_Is_Empty(t *testing.T) {
	user := &User{
		FirstName: "Gustavo",
		LastName:  "Martinez",
		Email:     "",
		Password:  "asdasd",
		Phone:     "123456",
	}

	_, err := user.ValidateEmail()
	if err == nil {
		t.Error("Expected err!=nil", "Got", err)
	}
}

func TestUser_Encode(t *testing.T) {
	user := &User{
		FirstName: "Gustavo",
		LastName:  "Martinez",
		Email:     "ggg@gmail.com",
		Password:  "Germ1=&",
		Phone:     "123456",
	}

	user.Encode()

	if user.Password == "Germ1=&" {
		t.Error("Not Expect", "Germ1=&", "Got", user.Password)
	}
}
