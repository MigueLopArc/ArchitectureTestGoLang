package services

import (
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
)

type INotesService interface {
	Create(note *DTOs.NoteDTO) (string, error)
	Update(id string, note *DTOs.NoteDTO) (*models.Note, error)
	Delete(id string) error
	List() ([]*models.Note, error)
	GetById(id string) (*models.Note, error)
}
