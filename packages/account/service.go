package account

import (
	"github.com/maderaka/goapp/packages/core/email"
	"database/sql"
	"errors"
	"golang.org/x/net/context"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (service *UserService) Create(ctx context.Context, user User) (User, error) {
	db := ctx.Value("db").(*sql.DB)
	users := NewUserRepository()
	u, _ := users.FindByEmail(db, user.Email)

	// Return somethings
	// If email address is already registered
	if u.Id != 0 {
		return User{}, errors.New("email is already registered")
	}

	// Create new user
	salt, e := user.HashPassword()
	user.PasswordSalt = salt
	if e != nil {
		return User{}, e
	}

	id, _ := users.Create(db, user)
	u, err := users.FindById(db, id)

	go sendEmail(u, ctx)
	return u, err

}

func (service *UserService) Activate(ctx context.Context, id int) User {
	return User{}
}

func sendEmail(user User, ctx context.Context) {
	emailEngine := ctx.Value("email").(*email.Engine)
	emailEngine.SendEmail(
		[]string{user.Email},
		"User Registration",
		"<html><body>Exception 1</body></html>Exception 1",
	)
}
