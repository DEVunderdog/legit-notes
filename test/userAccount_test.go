package test

import (
	"context"
	"log"
	"testing"

	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUserAccount(t *testing.T) database.User {
	username, email := utils.RandomUser()
	password, err := utils.HashPassword("something")

	if err != nil {
		log.Fatalf("Some error: %v", err)
	}

	arg := database.CreateUserParams{
		Username: username,
		Email: email,
		HashedPassword: password,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T){
	createRandomUserAccount(t)
}

