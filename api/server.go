package api

import (
	"fmt"

	db "example.com/referralgen/db/sqlc"
	"example.com/referralgen/token"
	"example.com/referralgen/util"
	"github.com/gin-gonic/gin"
)

// this server will be used to handle all the requests from the client
type Server struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
	router     *gin.Engine
}

// Creates new server instance
func NewServer(config util.Config, store db.Store) (*Server, error) {
	maker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{store: store, tokenMaker: maker, config: config}

	server.setupRouter()

	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/hello", server.helloWorld)
	router.POST("/auth/register", server.CreateUser)
	router.POST("/auth/login", server.LoginUser)

	authRoutes := router.Group("/templates").Use(authMiddleware(server.tokenMaker))

	// template handlers
	authRoutes.POST("/", server.CreateTemplate)
	authRoutes.GET("/", server.GetTemplatesByUser)
	authRoutes.PATCH("/", server.UpdateTemplate)
	authRoutes.GET("/search", server.GetTemplateByName)
	authRoutes.DELETE("/:id", server.DeleteTemplateById)

	authRoutes.POST("/generate", server.GenerateReferralTemplate)

	server.router = router
}

func (server *Server) helloWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

// Starts the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
