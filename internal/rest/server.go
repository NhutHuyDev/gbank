package rest

import (
	"github.com/NhutHuyDev/sgbank/internal/infra/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}
