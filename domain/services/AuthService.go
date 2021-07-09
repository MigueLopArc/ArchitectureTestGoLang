package services

import (
	"context"
	"fmt"
	"net/mail"
	"strings"

	db "github.com/MigueLopArc/ArchitectureTestGoLang/data"
	repos "github.com/MigueLopArc/ArchitectureTestGoLang/domain/dal/repositories"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/helpers"
	models "github.com/MigueLopArc/ArchitectureTestGoLang/domain/models"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/auth"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/responseCodes"
	DTOs "github.com/MigueLopArc/ArchitectureTestGoLang/presentation/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UsersRepository repos.IUsersRepository
	context         context.Context
	TokenGenerator  helpers.TokenGenerator
}

// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
func NewAuthService(ctx context.Context) IAuthService {
	var dbContext = db.New()

	var usersRepo repos.IUsersRepository = repos.NewUsersRepo(dbContext)

	return &AuthService{
		UsersRepository: usersRepo,
		context:         ctx,
		TokenGenerator:  *helpers.NewTokenGenerator(),
	}
}

func (authService *AuthService) SignUp(request *DTOs.SignUpModel) (*auth.JsonWebToken, *responseCodes.ApiResponse) {

	var errors []responseCodes.CommonResponseDetail = []responseCodes.CommonResponseDetail{}

	if !emailIsValid(request.Email) {
		errors = append(errors, *responseCodes.InvalidEmail)
	}
	if len(strings.TrimSpace(request.Password)) < 6 {
		errors = append(errors, *responseCodes.InvalidPassword)
	}
	if len(strings.TrimSpace(request.FirstName)) == 0 {
		errors = append(errors, *responseCodes.UserFirstNameNotFound)
	}
	if len(strings.TrimSpace(request.LastName)) == 0 {
		errors = append(errors, *responseCodes.UserLastNameNotFound)
	}

	if len(errors) > 0 {
		apiResponse := responseCodes.BuildBadRequestMessage(errors)
		return nil, &apiResponse
	}

	hash, hashErr := hashPassword(request.Password)

	if hashErr != nil {
		return nil, &responseCodes.UnknownError
	}

	var user *models.User = &models.User{
		Email:     request.Email,
		Password:  hash,
		Firstname: request.FirstName,
		Lastname:  request.LastName,
	}

	id, err := authService.UsersRepository.Create(authService.context, user)
	user.Id = id

	token, tokenErr := authService.TokenGenerator.GenerateJwtToken(&auth.JwtUserIdentity{
		UserId:   user.Id,
		UserName: fmt.Sprintf("%s %s", user.Firstname, user.Lastname),
		Email:    user.Email,
	})

	if tokenErr != nil {
		return nil, &responseCodes.UnknownError
	}

	return token, err
}

func (authService *AuthService) SignIn(request *DTOs.SignInModel) (*auth.JsonWebToken, *responseCodes.ApiResponse) {
	var errors []responseCodes.CommonResponseDetail = []responseCodes.CommonResponseDetail{}

	if !emailIsValid(request.Email) {
		errors = append(errors, *responseCodes.InvalidEmail)
	}

	if len(strings.TrimSpace(request.Password)) < 6 {
		errors = append(errors, *responseCodes.InvalidPassword)
	}

	if len(errors) > 0 {
		apiResponse := responseCodes.BuildBadRequestMessage(errors)
		return nil, &apiResponse
	}

	user, err := authService.UsersRepository.GetByEmail(authService.context, request.Email)

	if err != nil {
		return nil, err
	}

	if !checkPassword(user.Password, request.Password) {
		return nil, &responseCodes.WrongPassword
	}

	token, tokenErr := authService.TokenGenerator.GenerateJwtToken(&auth.JwtUserIdentity{
		UserId:   user.Id,
		UserName: fmt.Sprintf("%s %s", user.Firstname, user.Lastname),
		Email:    user.Email,
	})

	if tokenErr != nil {
		return nil, &responseCodes.UnknownError
	}

	return token, err
}

func emailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

// checkPassword compares HashPassword with the password and returns true if they match.
func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
