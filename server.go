package main

import (
	"github.com/MigueLopArc/ArchitectureTestGoLang/presentation/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	addRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func addRoutes(e *echo.Echo) {
	notesController := controllers.NewNotesController()
	e.POST("/notes", notesController.CreateNote)
	e.GET("/notes/:id", notesController.GetNote)
	e.PUT("/notes/:id", notesController.UpdateNote)
	e.DELETE("/notes/:id", notesController.DeleteNote)
	e.GET("/notes", notesController.GetNotes)
}
