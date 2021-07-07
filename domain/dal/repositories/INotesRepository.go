package repositories

import (
	"context"

	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
)

type INotesRepository interface {
	Create(ctx context.Context, note *models.Note) (string, *responseCodes.ApiResponse)
	GetById(ctx context.Context, id string) (*models.Note, *responseCodes.ApiResponse)
	List(ctx context.Context, limit, offset uint) ([]*models.Note, error)
	Update(ctx context.Context, id string, note *models.Note) error
	Delete(ctx context.Context, id string) error
}
