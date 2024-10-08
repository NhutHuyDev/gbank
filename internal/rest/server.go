package rest

import (
	"net/http"

	"github.com/NhutHuyDev/sgbank/internal/infra/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()

	_ = router.SetTrustedProxies([]string{"192.168.1.1"})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", currencyValidator)
	}

	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OKE",
		})
	})
	router.POST("/users", server.createUserHandler)

	router.GET("/accounts/:id", server.getAccountHandler)
	router.GET("/accounts", server.listAccountsHandler)
	router.POST("/accounts", server.createAccountHandler)

	router.POST("/transfers", server.transferHandler)

	server.router = router

	return server
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
