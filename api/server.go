package api

import (
	db "github.com/davigiroux/travel-blog-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Users API
	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUsers)
	router.DELETE("/users/:id", server.listUsers)
	router.PATCH("/users/:id", server.updateUser)

	// Articles API
	router.POST("/article", server.createArticle)
	router.POST("/article", server.createArticle)
	router.POST("/article", server.createArticle)
	router.POST("/article", server.createArticle)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
