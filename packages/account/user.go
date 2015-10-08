package account

import (
	"github.com/maderaka/goapp/packages/core"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

const (
	StatusPending   = "pending"
	StatusDeclined  = "declined"
	StatusConfirmed = "confirmed"
)

type User struct {
	Id             int32     `json:"id"`
	Username       string    `json:"username"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	PasswordHashed string    `json:"password"`
	PasswordSalt   string    `json:"-"`
	Avatar         string    `json:"avatar"`
	Bio            string    `json:"bio"`
	Status         string    `json:"status"`
	RegisteredAt   time.Time `json:"registered_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	LastLogin      time.Time `json:"last_login"`
	LastIp         string    `json:"last_ip"`
	TimeZone       string    `json:"timezone"`
}

func (user *User) Valid() []core.ResponseError {
	errors := make([]core.ResponseError, 0)
	if len(user.Username) == 0 {
		errors = append(errors, core.ResponseError{
			"username", core.ErrMissingField("username").Error(),
		})
	}

	if len(user.PasswordHashed) == 0 {
		errors = append(errors, core.ResponseError{
			"password", core.ErrMissingField("password").Error(),
		})
	}

	if len(user.Email) == 0 {
		errors = append(errors, core.ResponseError{
			"email", core.ErrMissingField("email").Error(),
		})
	}

	if len(user.FirstName) == 0 {
		errors = append(errors, core.ResponseError{
			"first_name", core.ErrMissingField("first_name").Error(),
		})
	}

	if len(user.LastName) == 0 {
		errors = append(errors, core.ResponseError{
			"last_name", core.ErrMissingField("last_name").Error(),
		})
	}
	return errors

}

func (user *User) HashPassword() (string, error) {
	return core.Encrypt(&user.PasswordHashed)
}

func (user *User) Fillable() []string {
	return []string{
		"username",
		"first_name",
		"last_name",
		"email",
		"password_hash",
		"password_salt",
		"bio",
		"status",
		"registered_at",
		"updated_at",
	}
}

type UserRepositoryInterface interface {
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (users *UserRepository) Create(db *sql.DB, user User) (int, error) {
	fields := strings.Join(user.Fillable(), ", ")
	now := time.Now().Format("2006-01-02 15:04:05")

	query := fmt.Sprintf(
		"INSERT INTO users("+fields+") VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') RETURNING id;",
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHashed,
		user.PasswordSalt,
		user.Bio,
		StatusPending,
		now,
		now,
	)

	var lastId int
	err := db.QueryRow(query).Scan(&lastId)
	if err != nil {
		return 0, err
	}
	return int(lastId), err

}

func (UserRepo *UserRepository) Collection(db *sql.DB, params map[string]int) ([]*User, error) {

	// Build query based on params
	query := "SELECT * FROM users"
	if v, ok := params["status"]; ok {
		query += " WHERE status = '%s'"
		query = fmt.Sprintf(query, v)
	}

	// Start to executing query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {

		user := new(User)
		err := rows.Scan(
			&user.Id, &user.FirstName, &user.LastName,
			&user.Avatar, &user.Bio,
			&user.LastLogin, &user.LastIp, &user.RegisteredAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (users *UserRepository) Update(db *sql.DB, user User) User {
	return User{}
}

func (users *UserRepository) FindById(db *sql.DB, id int) (User, error) {
	user := User{}
	fields := strings.Join(append(user.Fillable(), "id"), ", ")
	query := fmt.Sprintf("SELECT %s FROM users WHERE id = %d", fields, id)
	row := db.QueryRow(query)

	err := row.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHashed,
		&user.PasswordSalt,
		&user.Bio,
		&user.Status,
		&user.RegisteredAt,
		&user.UpdatedAt,
		&user.Id,
	)
	return user, err
}

func (users *UserRepository) FindByEmail(db *sql.DB, email string) (User, error) {
	user := User{}
	query := fmt.Sprintf("SELECT id, first_name FROM users WHERE email='%s'", email)
	err := db.QueryRow(query).Scan(&user.Id, &user.FirstName)
	return user, err
}

func (users *UserRepository) Delete(user User) bool {
	return false
}
