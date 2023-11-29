package product

import (
	"database/sql"
	"errors"

	"github.com/ncondezo/final/internal/domain"

	"github.com/go-sql-driver/mysql"
)

const (
	createUserQuery      = "INSERT INTO users (id, name, surname, email, password) VALUES (?, ?, ?, ?, ?)"
	findUserByEmailQuery = "SELECT * FROM users WHERE email = ?"
)

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserExists   = errors.New("user already exists")
)

type Repository interface {
	Create(product *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (repository *repository) Create(user *domain.User) (*domain.User, error) {
	var mysqlError *mysql.MySQLError
	stmt, err := repository.db.Prepare(createUserQuery)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		user.Id,
		user.Name,
		user.Surname,
		user.Email,
		user.Password,
	)
	ok := errors.As(err, &mysqlError)
	if ok {
		switch mysqlError.Number {
		case 1062:
			return nil, ErrorUserExists
		default:
			return nil, err
		}
	}
	return user, nil
}

func (repository *repository) FindByEmail(email string) (*domain.User, error) {
	stmt, err := repository.db.Prepare(findUserByEmailQuery)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	found := scanUser(stmt.QueryRow(email))
	if found.Id == "" {
		return nil, ErrorUserNotFound
	}
	return found, nil
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanUser(scanner scanner) *domain.User {
	userScanned := &domain.User{}
	_ = scanner.Scan(
		&userScanned.Id,
		&userScanned.Name,
		&userScanned.Surname,
		&userScanned.Email,
		&userScanned.Password,
	)
	return userScanned
}
