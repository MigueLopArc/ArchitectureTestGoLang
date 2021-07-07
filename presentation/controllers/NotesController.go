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
	n := &DTOs.NoteDTO{
		UserId: "0029f464-ea55-4ab2-8549-990695c18e02", // This should be bind from JWT
	}

	if err := c.Bind(n); err != nil {
		return err
	}
	notesService := services.NewNotesService(c.Request().Context())
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
	n := &DTOs.NoteDTO{
		UserId: "0029f464-ea55-4ab2-8549-990695c18e02", // This should be bind from JWT
	}

	if err := c.Bind(n); err != nil {
		return err
	}
	notesService := services.NewNotesService(c.Request().Context())
	var result, _ = notesService.Update(c.Param("id"), n)

	return c.JSON(http.StatusOK, result)
}

func (notesController *NotesController) GetNote(c echo.Context) error {
	// return c.JSON(responseCodes.Test2BadRequest.HttpStatusCode, responseCodes.Test2BadRequest.Detail)
	notesService := services.NewNotesService(c.Request().Context())
	var result, err = notesService.GetById(c.Param("id"))
	if err != nil {
		return c.JSON(err.HttpStatusCode, err.Detail)
	}
	return c.JSON(http.StatusOK, result)
}

func (notesController *NotesController) DeleteNote(c echo.Context) error {
	notesService := services.NewNotesService(c.Request().Context())
	if err := notesService.Delete(c.Param("id")); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusNoContent)
}

func (notesController *NotesController) GetNotes(c echo.Context) error {
	notesService := services.NewNotesService(c.Request().Context())
	var result, _ = notesService.List()

	return c.JSON(http.StatusOK, result)
}
