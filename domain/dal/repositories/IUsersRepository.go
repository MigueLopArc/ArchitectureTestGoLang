package repositories

import (
	"context"

	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
)

type IUsersRepository interface {
	Create(ctx context.Context, user *models.User) (string, *responseCodes.ApiResponse)
	GetByEmail(ctx context.Context, email string) (*models.User, *responseCodes.ApiResponse)
}
