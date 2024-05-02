package repository

import (
	"database/sql"
	"errors"

	"github.com/tonminhce/auth-microservice/models"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	db *sql.DB
}

func NewLogin(db *sql.DB) *Login {
	return &Login{db: db}
}

func (l *Login) GetUserByUserName(userName, password string) (models.User, *models.ErrorDetail) {
	var user models.User
	var hashedPassword string
	err := l.db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", userName).Scan(&user.Id, &user.UserName, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, &models.ErrorDetail{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: "UserName not found",
			}
		} else {
			return models.User{}, &models.ErrorDetail{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: "Database error",
			}
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.User{}, &models.ErrorDetail{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: "Password is incorrect",
			}
		} else {
			return models.User{}, &models.ErrorDetail{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: "Internal server error",
			}
		}
	}
	return user, nil
}

func (l *Login) CreateUser(user models.User) *models.ErrorDetail {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &models.ErrorDetail{
			ErrorType:    models.ErrorTypeError,
			ErrorMessage: "Error hashing password",
		}
	}

	_, err = l.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.UserName, hashedPassword)
	if err != nil {
		return &models.ErrorDetail{
			ErrorType:    models.ErrorTypeError,
			ErrorMessage: "Error creating user",
		}
	}

	return nil
}
