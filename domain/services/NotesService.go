package services

import (
	"context"
	"time"

	db "github.com/MigueLopArc/ArchitectureTestGoLang/data"
	repos "github.com/MigueLopArc/ArchitectureTestGoLang/domain/dal/repositories"
	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
)

type NotesService struct {
	NotesRepository repos.INotesRepository
	context         context.Context
}

// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
func NewNotesService(ctx context.Context) INotesService {
	var dbContext = db.New()

	var notesRepo repos.INotesRepository = repos.NewNotesRepo(dbContext)

	return &NotesService{
		NotesRepository: notesRepo,
		context:         ctx,
	}
}

func (notesService *NotesService) Create(request *DTOs.NoteDTO) (string, error) {

	var note *models.Note = &models.Note{
		Title:   request.Title,
		Content: request.Content,
		UserId:  request.UserId,
	}

	id, _ := notesService.NotesRepository.Create(notesService.context, note)

	return id, nil
}

func (notesService *NotesService) Update(id string, request *DTOs.NoteDTO) (*models.Note, error) {

	oldNoteData, _ := notesService.NotesRepository.GetById(notesService.context, id)

	var note *models.Note = &models.Note{
		Title:            request.Title,
		Content:          request.Content,
		ModificationDate: time.Now(),
	}

	err := notesService.NotesRepository.Update(notesService.context, id, note)

	if err != nil {
		return nil, err
	}

	oldNoteData.Title = note.Title
	oldNoteData.Content = note.Content
	oldNoteData.ModificationDate = note.ModificationDate

	return oldNoteData, nil
}

func (notesService *NotesService) Delete(id string) error {

	err := notesService.NotesRepository.Delete(notesService.context, id)

	if err != nil {
		return err
	}

	return nil
}

func (notesService *NotesService) GetById(id string) (*models.Note, error) {

	result, _ := notesService.NotesRepository.GetById(notesService.context, id)

	return result, nil
}

func (notesService *NotesService) List() ([]*models.Note, error) {
	result, _ := notesService.NotesRepository.List(notesService.context, 0, 0)

	return result, nil
}
