package test

import (
	"context"
	"testing"

	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/utils"
	"github.com/stretchr/testify/require"
)

func createRandomNote(t *testing.T) database.Note {

	user := createRandomUserAccount(t)
	title, description := utils.RandomNote()

	arg := database.CreateNoteParams{
		UserID: user.ID,
		Title: title,
		Description: description,
	}

	note, err := testQueries.CreateNote(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, note)
	require.Equal(t, arg.UserID, note.UserID)
	require.Equal(t, arg.Title, note.Title)
	require.Equal(t, arg.Description, note.Description)

	require.NotZero(t, note.ID)
	require.NotZero(t, note.CreatedAt)

	return note

}

func TestCreateNote(t *testing.T) {
	createRandomNote(t)
}

func TestNoteUpdate(t *testing.T) {
	note1 := createRandomNote(t)
	title, description := utils.RandomNote()

	arg := database.UpdateNoteParams{
		ID: note1.ID,
		Title: title,
		Description: description,
	}

	note2, err := testQueries.UpdateNote(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, note2)
	
	require.Equal(t, note1.ID, note2.ID)
	require.Equal(t, arg.Title, note2.Title)
	require.Equal(t, arg.Description, note2.Description)
	
	require.NotZero(t, note2.CreatedAt)
	require.NotZero(t, note2.UpdatedAt)
}

