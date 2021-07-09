package controllers

import (
	"fmt"
	"net/http"

	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/services"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
	"github.com/labstack/echo/v4"
)

type NotesController struct{}

func NewNotesController() NotesController {
	return NotesController{}
}

// Handler
func (notesController *NotesController) CreateNote(c echo.Context) error {
	userId := c.Get(echo.HeaderAuthorization).(string)
	n := &DTOs.NoteDTO{}

	if err := c.Bind(n); err != nil {
		return err
	}
	notesService := services.NewNotesService(c.Request().Context(), userId)
	var resultingId, apiResponse = notesService.Create(n)

	if apiResponse != nil {
		return c.JSON(apiResponse.HttpStatusCode, apiResponse.Detail)
	}

	c.Response().Header().Set("location", fmt.Sprintf("notes/%s", resultingId))
	// c.Response().WriteHeader(http.StatusCreated)
	return c.NoContent(http.StatusCreated)
}

// Handler
func (notesController *NotesController) UpdateNote(c echo.Context) error {
	userId := c.Get(echo.HeaderAuthorization).(string)
	n := &DTOs.NoteDTO{}

	if err := c.Bind(n); err != nil {
		return err
	}
	notesService := services.NewNotesService(c.Request().Context(), userId)
	var result, err = notesService.Update(c.Param("id"), n)

	if err != nil {
		return c.JSON(err.HttpStatusCode, err.Detail)
	}

	return c.JSON(http.StatusOK, result)
}

func (notesController *NotesController) GetNote(c echo.Context) error {
	// return c.JSON(responseCodes.Test2BadRequest.HttpStatusCode, responseCodes.Test2BadRequest.Detail)
	userId := c.Get(echo.HeaderAuthorization).(string)
	noteId := c.Param("id")
	notesService := services.NewNotesService(c.Request().Context(), userId)
	var result, err = notesService.GetById(noteId)
	if err != nil {
		return c.JSON(err.HttpStatusCode, err.Detail)
	}
	return c.JSON(http.StatusOK, result)
}

func (notesController *NotesController) DeleteNote(c echo.Context) error {
	userId := c.Get(echo.HeaderAuthorization).(string)
	noteId := c.Param("id")
	notesService := services.NewNotesService(c.Request().Context(), userId)
	if err := notesService.Delete(noteId); err != nil {
		return c.JSON(err.HttpStatusCode, err.Detail)
	}

	return c.NoContent(http.StatusNoContent)
}

func (notesController *NotesController) GetUserNotes(c echo.Context) error {
	userId := c.Get(echo.HeaderAuthorization).(string)
	notesService := services.NewNotesService(c.Request().Context(), userId)
	var result, err = notesService.GetUserNotes()
	if err != nil {
		return c.JSON(err.HttpStatusCode, err.Detail)
	}
	return c.JSON(http.StatusOK, result)
}
