package repositories

import (
	"context"

	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
)

type INotesRepository interface {
	Create(ctx context.Context, note *models.Note) (string, *responseCodes.ApiResponse)
	GetById(ctx context.Context, id string) (*models.Note, *responseCodes.ApiResponse)
	GetUserNotes(ctx context.Context, userId string) ([]*models.Note, *responseCodes.ApiResponse)
	Update(ctx context.Context, id string, note *models.Note) *responseCodes.ApiResponse
	Delete(ctx context.Context, id string) *responseCodes.ApiResponse
}
