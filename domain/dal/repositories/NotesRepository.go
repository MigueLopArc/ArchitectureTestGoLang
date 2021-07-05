package repositories

import (
	"context"

	db "github.com/MigueLopArc/ArchitectureTestGoLang/data"
	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
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

func (notesRepo *NotesRepository) List(ctx context.Context, limit, offset uint) ([]*models.Note, error) {
	q := `
        SELECT * FROM notes;
    `

	rows, err := notesRepo.DbContext.DB.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []*models.Note
	for rows.Next() {
		var note models.Note
		rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content,
			&note.CreationDate, &note.ModificationDate)
		notes = append(notes, &note)
	}

	return notes, nil
}

func (notesRepo *NotesRepository) GetById(ctx context.Context, id string) (*models.Note, error) {
	q := `
        SELECT * FROM notes WHERE id = $1;
    `

	row := notesRepo.DbContext.DB.QueryRowContext(ctx, q, id)

	var note models.Note

	err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.CreationDate,
		&note.ModificationDate)

	if err != nil {
		return &models.Note{}, err
	}

	return &note, nil
}

func (notesRepo *NotesRepository) Create(ctx context.Context, note *models.Note) (string, error) {
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
		return "", err
	}

	return note.Id, nil
}

func (notesRepo *NotesRepository) Update(ctx context.Context, id string, note *models.Note) error {
	q := `
    UPDATE notes SET title = $1, content = $2, modification_date = $3
        WHERE id=$4;
    `

	stmt, err := notesRepo.DbContext.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, note.Title, note.Content, note.ModificationDate, id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (notesRepo *NotesRepository) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM notes WHERE id=$1;`

	stmt, err := notesRepo.DbContext.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
