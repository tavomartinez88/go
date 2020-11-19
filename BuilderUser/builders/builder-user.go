// builders is a package container to BuilderUser's contructor, it is responsable to build an domain.User
package builders

import "github.com/tavomartinez88/go/BuilderUser/domain"

// IConstructorUser is an interface, this declare a method Build to return a domain.User
type IConstructorUser interface {
	Build(fields map[string]string) *domain.User
}

// BuilderUser implement a constructor concrete with respect to design pattern builder
type BuilderUser struct {}

// Build is a method to return a domain.User, where it recive a map with data (firstname, lastname, etc..)
func (c BuilderUser) Build(fields map[string]string) *domain.User {
	user := &domain.User{
		FirstName: fields["firstname"],
		LastName:  fields["lastname"],
		Email:     fields["email"],
		Password:  fields["password"],
		Phone:     fields["phone"],
	}

	user.Encode()
	return user
}
