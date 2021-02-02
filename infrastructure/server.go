package infrastructure

import (
	"server/ent"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Store  *ent.Client
	Router *gin.Engine
}
