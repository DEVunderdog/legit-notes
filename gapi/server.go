package gapi

import (
	"fmt"

	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/pb"
	"github.com/DEVunderdog/legit-notes/token"
	"github.com/DEVunderdog/legit-notes/utils"
)

type Server struct {
	pb.UnimplementedLegitNotesServer
	config utils.Config
	store database.Store
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store database.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("error occured: %q", err)
	}

	server := &Server {
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}