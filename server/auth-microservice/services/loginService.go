package services

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tonminhce/auth-microservice/models"
	"github.com/tonminhce/auth-microservice/repository"
	"github.com/tonminhce/auth-microservice/token"
)

type Login struct {
	logger          *log.Logger
	flags           *models.Flags
	loginRepository *repository.Login
}

func NewLogin(l *log.Logger, f *models.Flags, db *sql.DB) *Login {
	return &Login{
		logger:          l,
		flags:           f,
		loginRepository: repository.NewLogin(db),
	}
}

func (l *Login) GetToken(loginModel models.LoginRequest, origin string) (string, *models.ErrorDetail) {
	user, err := l.loginRepository.GetUserByUserName(loginModel.UserName, loginModel.Password)
	if err != nil {
		return "", err
	}
	var claims = &models.JwtClaims{
		ComapnyId: strconv.Itoa(user.Id),
		Username:  user.Name,
		Roles:     user.Roles,
		StandardClaims: jwt.StandardClaims{
			Audience: origin,
		},
	}
	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	return token.GenrateToken(claims, expirationTime)

}

func (*Login) VerifyToken(tokenString, referer string) (bool, *models.JwtClaims) {
	return token.VerifyToken(tokenString, referer)
}
