package services

import (
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/auth"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
)

type IAuthService interface {
	SignIn(user *DTOs.SignInModel) (*auth.JsonWebToken, *responseCodes.ApiResponse)
	SignUp(user *DTOs.SignUpModel) (*auth.JsonWebToken, *responseCodes.ApiResponse)
}
