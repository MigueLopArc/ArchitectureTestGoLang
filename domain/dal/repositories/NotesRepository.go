package repositories

import (
	"context"
	"database/sql"

	db "github.com/MigueLopArc/ArchitectureTestGoLang/data"
	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
)

type NotesRepository struct {
	DbContext *db.DatabaseContext
}

// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
func NewNotesRepo(dbContext *db.DatabaseContext) INotesRepository {
	return &NotesRepository{
		DbContext: dbContext,
	}
}

func (notesRepo *NotesRepository) GetUserNotes(ctx context.Context, userId string) ([]*models.Note, *responseCodes.ApiResponse) {
	q := `
        SELECT * FROM notes WHERE user_id = $1;
    `

	rows, err := notesRepo.DbContext.DB.QueryContext(ctx, q, userId)

	if err != nil {
		return nil, &responseCodes.UnknownError
	}

	defer rows.Close()

	var notes []*models.Note = []*models.Note{}
	for rows.Next() {
		var note models.Note
		rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content,
			&note.CreationDate, &note.ModificationDate)
		notes = append(notes, &note)
	}

	return notes, nil
}

func (notesRepo *NotesRepository) GetById(ctx context.Context, id string) (*models.Note, *responseCodes.ApiResponse) {
	q := `
        SELECT * FROM notes WHERE id = $1;
    `

	row := notesRepo.DbContext.DB.QueryRowContext(ctx, q, id)

	var note models.Note

	err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.CreationDate,
		&note.ModificationDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &responseCodes.EntityNotFound
		}
		return nil, &responseCodes.UnknownError
	}

	return &note, nil
}

func (notesRepo *NotesRepository) Create(ctx context.Context, note *models.Note) (string, *responseCodes.ApiResponse) {
	q := `
        INSERT INTO notes (title, content, user_id)
            VALUES ($1, $2, $3)
        RETURNING id;
    `

	row := notesRepo.DbContext.DB.QueryRowContext(
		ctx, q, &note.Title, &note.Content, &note.UserId,
	)

	err := row.Scan(&note.Id)

	if err != nil {
		/// Real errors should be catch and rethrow as an Unknown error for security reasons
		return "", &responseCodes.UnknownError
	}

	return note.Id, nil
}

func (notesRepo *NotesRepository) Update(ctx context.Context, id string, note *models.Note) *responseCodes.ApiResponse {
	q := `
    UPDATE notes SET title = $1, content = $2, modification_date = $3
        WHERE id=$4;
    `

	stmt, err := notesRepo.DbContext.DB.PrepareContext(ctx, q)
	if err != nil {
		return &responseCodes.UnknownError
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, note.Title, note.Content, note.ModificationDate, id,
	)

	if err != nil {
		return &responseCodes.UnknownError
	}

	return nil
}

func (notesRepo *NotesRepository) Delete(ctx context.Context, id string) *responseCodes.ApiResponse {
	q := `DELETE FROM notes WHERE id=$1;`

	stmt, err := notesRepo.DbContext.DB.PrepareContext(ctx, q)
	if err != nil {
		return &responseCodes.UnknownError
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return &responseCodes.UnknownError
	}

	return nil
}
