package controllers

import (
	"net/http"

	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/services"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
	"github.com/labstack/echo/v4"
)

type AuthController struct{}

func NewAuthController() AuthController {
	return AuthController{}
}

// Handler
func (authController *AuthController) SignUp(c echo.Context) error {
	signUpRequest := &DTOs.SignUpModel{}

	if err := c.Bind(signUpRequest); err != nil {
		return err
	}
	authService := services.NewAuthService(c.Request().Context())
	var result, apiResponse = authService.SignUp(signUpRequest)

	if apiResponse != nil {
		return c.JSON(apiResponse.HttpStatusCode, apiResponse.Detail)
	}
	return c.JSON(http.StatusOK, result)
}

// Handler
func (authController *AuthController) SignIn(c echo.Context) error {
	signInRequest := &DTOs.SignInModel{}

	if err := c.Bind(signInRequest); err != nil {
		return err
	}
	authService := services.NewAuthService(c.Request().Context())
	var result, apiResponse = authService.SignIn(signInRequest)

	if apiResponse != nil {
		return c.JSON(apiResponse.HttpStatusCode, apiResponse.Detail)
	}
	return c.JSON(http.StatusOK, result)
}
