package gapi

import (
	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user database.User) *pb.User {
	return &pb.User{
		Username: user.Username,
		Email: user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}