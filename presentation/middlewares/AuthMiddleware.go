package middlewares

import (
	"errors"
	"strings"

	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/helpers"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
	"github.com/labstack/echo/v4"
)

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("Autorization is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("Invalid autorization format")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("Invalid autorization format")
	}

	return l[1], nil
}

// ServerHeader middleware adds a `Server` header to the response.
func ValidateJwt(next echo.HandlerFunc) echo.HandlerFunc {
	tokenGenerator := helpers.NewTokenGenerator()
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
		// tokenString := c.Request().Header[echo.HeaderAuthorization][0]
		tokenString, err := tokenFromAuthorization(authHeader)
		if err != nil {
			return c.JSON(responseCodes.AuthorizarionMissing.HttpStatusCode, responseCodes.AuthorizarionMissing.Detail)
		}

		claims, err := tokenGenerator.GetClaimsFromToken(tokenString)
		if err != nil {
			return c.JSON(responseCodes.AuthorizationFailed.HttpStatusCode, responseCodes.AuthorizationFailed.Detail)
		}
		c.Set(echo.HeaderAuthorization, claims.Subject)
		// c.Request().ctx = context.WithValue(c.Request().Context(), echo.HeaderAuthorization, claims.Subject)
		return next(c)
	}
}
