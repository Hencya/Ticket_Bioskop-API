package usersEntity

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/businesses"
	"TiBO_API/helpers"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type UserServices struct {
	UserRepository Repository
	jwtAuth        *auth.ConfigJWT
	ContextTimeout time.Duration
}

func NewUserServices(repoUser Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &UserServices{
		UserRepository: repoUser,
		jwtAuth:        auth,
		ContextTimeout: timeout,
	}
}

func (us *UserServices) Register(ctx context.Context, userDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	existedUser, err := us.UserRepository.GetByEmail(ctx, userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedUser != (Domain{}) {
		return nil, businesses.ErrDuplicateEmail
	}

	userDomain.Password, err = helpers.HashPassword(userDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.UserRepository.CreateNewUser(ctx, userDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *UserServices) Login(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	userDomain, err := us.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", businesses.ErrEmailNotRegistered
	}

	if !helpers.ValidateHash(password, userDomain.Password) {
		return "", businesses.ErrPassword
	}

	token := us.jwtAuth.GenerateToken(userDomain.Uuid.String(), userDomain.Role)

	return token, nil
}

func (us *UserServices) Logout(ctx echo.Context) error {
	cookie, err := auth.LogoutCookie(ctx)
	fmt.Println(cookie)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	result, err := us.UserRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (us *UserServices) UpdateById(ctx context.Context, userDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	passwordHash, err := helpers.HashPassword(userDomain.Password)
	if err != nil {
		panic(err)
	}

	userDomain.Password = passwordHash
	result, err := us.UserRepository.UpdateUser(ctx, id, userDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (us *UserServices) UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	user, err := us.UserRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	user.Avatar = fileLocation
	updateAvatar, err := us.UserRepository.UploadAvatar(ctx, id, &user)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}

func (us *UserServices) DeleteUser(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	res, err := us.UserRepository.DeleteUserByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundUser
	}
	return res, nil
}
