package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createNoteRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) createNote(ctx *gin.Context) {
	var req createNoteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := database.CreateNoteParams{
		UserID:      authPayload.Username,
		Title:       req.Title,
		Description: req.Description,
	}

	note, err := server.store.CreateNote(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, note)
}

type getNoteRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getNote(ctx *gin.Context) {

	var req getNoteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := database.GetNoteParams{
		UserID: authPayload.Username,
		ID: sql.NullInt32{
			Int32: int32(req.ID),
			Valid: true,
		},
	}

	note, err := server.store.GetNote(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if note.UserID != authPayload.Username {
		err := errors.New("note doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, note)

}

type listNoteRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset"`
}

func (server *Server) listNotes(ctx *gin.Context) {

	var req listNoteRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := database.ListNotesParams{
		UserID: authPayload.Username,
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	notes, err := server.store.ListNotes(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, notes)

}

type updateNoteParams struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) updateNotes(ctx *gin.Context) {

	var req updateNoteParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := database.UpdateNoteParams{
		ID: sql.NullInt32{
			Int32: int32(req.ID),
			Valid: true,
		},
		Title:       req.Title,
		Description: req.Description,
	}

	note, err := server.store.UpdateNote(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if note.UserID != authPayload.Username {
		err := errors.New("note doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, note)
}

type updateNoteDescriptionParams struct {
	ID          int32  `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) updateNoteDescription(ctx *gin.Context) {
	var req updateNoteDescriptionParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := database.UpdateNoteDescriptionParams{
		ID: sql.NullInt32{
			Int32: int32(req.ID),
			Valid: true,
		},
		Description: req.Description,
	}

	note, err := server.store.UpdateNoteDescription(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if note.UserID != authPayload.Username {
		err := errors.New("note doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, note)
}

type updateNoteTitleParams struct {
	ID    int32  `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
}

func (server *Server) updateNoteTitle(ctx *gin.Context) {
	var req updateNoteTitleParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := database.UpdateNoteTitleParams{
		ID: sql.NullInt32{
			Int32: int32(req.ID),
			Valid: true,
		},
		Title: req.Title,
	}

	note, err := server.store.UpdateNoteTitle(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if note.UserID != authPayload.Username {
		err := errors.New("note doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, note)
}
