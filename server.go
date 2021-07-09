package main

import (
	"github.com/MigueLopArc/ArchitectureTestGoLang/presentation/controllers"
	"github.com/MigueLopArc/ArchitectureTestGoLang/presentation/middlewares"
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
	notes := e.Group("/notes", middlewares.ValidateJwt) // Add middleware to all notes endpoints
	notes.POST("", notesController.CreateNote)
	//e.GET("/notes/:id", notesController.GetNote, middlewares.ValidateJwt)
	notes.GET("/:id", notesController.GetNote)
	notes.PUT("/:id", notesController.UpdateNote)
	notes.DELETE("/:id", notesController.DeleteNote)
	notes.GET("", notesController.GetUserNotes)

	authController := controllers.NewAuthController()
	e.POST("/sign-in", authController.SignIn)
	e.POST("/sign-up", authController.SignUp)
}
