package routers

import (
	"database/sql"
	"log"

	"github.com/tonminhce/auth-microservice/handlers"
	"github.com/tonminhce/auth-microservice/models"
	"github.com/tonminhce/auth-microservice/token"

	"github.com/gin-gonic/gin"
)

type Login struct {
	logger       *log.Logger
	loginHandler *handlers.Login
	flags        *models.Flags
}

func NewRoute(l *log.Logger, f *models.Flags, db *sql.DB) *Login {
	loginHandler := handlers.NewLogin(l, f, db)
	token.Init()

	return &Login{
		logger:       l,
		loginHandler: loginHandler,
		flags:        f,
	}
}

func (r *Login) RegisterRoutes() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.POST("/login", r.loginHandler.Login)
	ginEngine.POST("/verifyToken", r.loginHandler.VerifyToken)
	return ginEngine
}
