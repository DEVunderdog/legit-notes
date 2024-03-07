package test

import (
	"context"
	"testing"
	"time"

	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUserAccount(t *testing.T) database.User {
	username, email := utils.RandomUser()
	password := utils.RandomPassword()

	arg := database.CreateUserParams{
		Username: username,
		Email: email,
		Password: password,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T){
	createRandomUserAccount(t)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomUserAccount(t)

	arg := database.UpdateAccountParams{
		ID: account1.ID,
		Password: utils.PasswordGenerator(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, arg.Password, account2.Password)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}