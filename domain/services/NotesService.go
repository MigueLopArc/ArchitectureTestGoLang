package services

import (
	"context"
	"strings"
	"time"

	db "github.com/MigueLopArc/ArchitectureTestGoLang/data"
	repos "github.com/MigueLopArc/ArchitectureTestGoLang/domain/dal/repositories"
	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
)

type NotesService struct {
	NotesRepository repos.INotesRepository
	context         context.Context
	UserId          string
}

// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
func NewNotesService(ctx context.Context, userId string) INotesService {
	var dbContext = db.New()

	var notesRepo repos.INotesRepository = repos.NewNotesRepo(dbContext)

	return &NotesService{
		NotesRepository: notesRepo,
		context:         ctx,
		UserId:          userId,
	}
}

func (notesService *NotesService) Create(request *DTOs.NoteDTO) (string, *responseCodes.ApiResponse) {

	var errors []responseCodes.CommonResponseDetail = []responseCodes.CommonResponseDetail{}

	if len(strings.TrimSpace(request.Title)) < 3 {
		errors = append(errors, *responseCodes.NoteTitleNotFound)
	}
	if len(strings.TrimSpace(request.Content)) == 0 {
		errors = append(errors, *responseCodes.NoteContentNotFound)
	}

	if len(errors) > 0 {
		apiResponse := responseCodes.BuildBadRequestMessage(errors)
		return "", &apiResponse
	}

	var note *models.Note = &models.Note{
		Title:   request.Title,
		Content: request.Content,
		UserId:  notesService.UserId,
	}

	id, err := notesService.NotesRepository.Create(notesService.context, note)

	return id, err
}

func (notesService *NotesService) Update(id string, request *DTOs.NoteDTO) (*models.Note, *responseCodes.ApiResponse) {

	oldNoteData, getErr := notesService.NotesRepository.GetById(notesService.context, id)

	if getErr != nil {
		return nil, getErr
	}

	var note *models.Note = &models.Note{
		Title:            request.Title,
		Content:          request.Content,
		ModificationDate: time.Now(),
	}

	if oldNoteData.UserId != notesService.UserId {
		return nil, &responseCodes.EntityDoesNotBelongToUser
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

func (notesService *NotesService) Delete(id string) *responseCodes.ApiResponse {
	oldNoteData, getErr := notesService.NotesRepository.GetById(notesService.context, id)

	if getErr != nil {
		return getErr
	}

	if oldNoteData.UserId != notesService.UserId {
		return &responseCodes.EntityDoesNotBelongToUser
	}

	err := notesService.NotesRepository.Delete(notesService.context, id)

	if err != nil {
		return err
	}

	return nil
}

func (notesService *NotesService) GetById(id string) (*models.Note, *responseCodes.ApiResponse) {

	result, err := notesService.NotesRepository.GetById(notesService.context, id)

	if err != nil {
		return nil, err
	}

	if result.UserId != notesService.UserId {
		return nil, &responseCodes.EntityDoesNotBelongToUser
	}

	return result, err
}

func (notesService *NotesService) GetUserNotes() ([]*models.Note, *responseCodes.ApiResponse) {
	result, err := notesService.NotesRepository.GetUserNotes(notesService.context, notesService.UserId)

	return result, err
}
