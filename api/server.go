package api

import (
	"fmt"

	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/token"
	"github.com/DEVunderdog/legit-notes/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     utils.Config
	store      database.Store
	tokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(config utils.Config, store *database.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      *store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {

	router := gin.Default()

	router.POST("/users", server.createUserAccount)
	router.POST("/users/login", server.loginUser)

	server.Router = router
}
