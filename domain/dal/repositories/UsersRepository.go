package repositories

import (
	"context"
	"database/sql"

	db "github.com/MigueLopArc/ArchitectureTestGoLang/data"
	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
	"github.com/lib/pq"
)

type UsersRepository struct {
	DbContext *db.DatabaseContext
}

// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
func NewUsersRepo(dbContext *db.DatabaseContext) IUsersRepository {
	return &UsersRepository{
		DbContext: dbContext,
	}
}

func (notesRepo *UsersRepository) GetByEmail(ctx context.Context, email string) (*models.User, *responseCodes.ApiResponse) {
	q := `
        SELECT * FROM users WHERE email = $1;
    `

	row := notesRepo.DbContext.DB.QueryRowContext(ctx, q, email)

	var user models.User

	err := row.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.Password,
		&user.CreationDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &responseCodes.UserNotFound
		}
		return nil, &responseCodes.UnknownError
	}

	return &user, nil
}

func (notesRepo *UsersRepository) Create(ctx context.Context, user *models.User) (string, *responseCodes.ApiResponse) {
	q := `
        INSERT INTO users (email, first_name, last_name, password)
            VALUES ($1, $2, $3, $4)
        RETURNING id;
    `

	row := notesRepo.DbContext.DB.QueryRowContext(
		ctx, q, &user.Email, &user.Firstname, &user.Lastname, &user.Password,
	)

	err := row.Scan(&user.Id)

	if err != nil {
		/// Real errors should be catch and rethrow as an Unknown error for security reasons
		if err.(*pq.Error).Code == "23505" { // Duplicate index violation (In this case the email)
			return "", &responseCodes.UserAlreadyExists
		}
		return "", &responseCodes.UnknownError
	}

	return user.Id, nil
}
