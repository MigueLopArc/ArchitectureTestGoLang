package services

import (
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
)

type IAuthService interface {
	SignIn(user *DTOs.SignInModel) (*models.User, *responseCodes.ApiResponse)
	SignUp(user *DTOs.SignUpModel) (*models.User, *responseCodes.ApiResponse)
}
