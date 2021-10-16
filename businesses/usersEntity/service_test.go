package usersEntity_test

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/businesses/usersEntity"
	"TiBO_API/businesses/usersEntity/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var (
	mockUsersRepository mocks.Repository
	userServices usersEntity.Service
	hashedPassword string
	usersDomain usersEntity.Domain
	usersHashDomain usersEntity.Domain
)



func TestMain(m *testing.M){
	userServices = usersEntity.NewUserServices(&mockUsersRepository,&auth.ConfigJWT{},time.Second * 2)
	usersDomain = usersEntity.Domain{
		Name        : "Testing User",
		Username    : "testingUser",
		Email       : "testinguser@gmail.com",
		Password    : "testing password",
		PhoneNumber : "08123482342",
		Avatar: "images/avatar/avatar.png",
		Token:  "asiudgasigdaiusgdiausgdi",
	}

	usersHashDomain = usersEntity.Domain{
		Name        : "Testing User",
		Username    : "testingUser",
		Email       : "testinguser@gmail.com",
		Password    : "asdgajkdkjashdkjahsdkasdhjah",
		PhoneNumber : "08123482342",
		Token:  "asiudgasigdaiusgdiausgdi",
	}
	m.Run()
}

func TestUserServices_Register(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(usersEntity.Domain{}, nil).Once()
		mockUsersRepository.On("CreateNewUser",mock.Anything,mock.Anything).Return(&usersDomain, nil).Once()

		req := &usersEntity.Domain{
			Name        : "Testing User",
			Username    : "testingUser",
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
			PhoneNumber : "08123482342",
		}

		res, err := userServices.Register(context.Background(),req)

		assert.Nil(t, err)
		assert.Equal(t, usersDomain, *res)
	})
	t.Run("Invalid Test || duplicate email", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(usersDomain, nil).Once()
		mockUsersRepository.On("CreateNewUser",mock.Anything,mock.Anything).Return(&usersEntity.Domain{}, assert.AnError).Once()

		req := &usersEntity.Domain{
			Name        : "Testing User",
			Username    : "testingUser",
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
			PhoneNumber : "08123482342",
		}

		res, err := userServices.Register(context.Background(),req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Equal(t, usersDomain.Email, req.Email)
	})
	t.Run("Invalid Test || err not found", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(usersEntity.Domain{}, assert.AnError).Once()
		mockUsersRepository.On("CreateNewUser",mock.Anything,mock.Anything).Return(&usersEntity.Domain{}, assert.AnError).Once()

		req := &usersEntity.Domain{
			Name        : "Testing User",
			Username    : "testingUser",
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
			PhoneNumber : "08123482342",
		}

		res, err := userServices.Register(context.Background(),req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.NotEqual(t, usersDomain, req)
	})
	t.Run("Invalid Test || internal error", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(usersEntity.Domain{},nil).Once()
		mockUsersRepository.On("CreateNewUser",mock.Anything,mock.Anything).Return(&usersEntity.Domain{}, assert.AnError).Once()

		req := &usersEntity.Domain{
			Name        : "Testing User",
			Username    : "testingUser",
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
			PhoneNumber : "08123482342",
		}

		res, err := userServices.Register(context.Background(),req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.NotEqual(t, usersDomain, req)
	})
}

func TestUserServices_Login(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.Anything, mock.Anything).Return(usersDomain, nil).Once()

		req := usersEntity.Domain{
			Email       : "testinguser@gmail.com",
			Password    : "asdgajkdkjashdkjahsdkasdhjah",
		}

		_, err := userServices.Login(context.Background(),req.Email,req.Password)
		token := "asiudgasigdaiusgdiausgdi"

		assert.Equal(t, token,usersHashDomain.Token)
		assert.Equal(t, req.Password,usersHashDomain.Password)
		assert.NotNil(t, err)
	})
	t.Run("Invalid Test | password and email null", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		req := usersEntity.Domain{
			Email       : "",
			Password    : "",
		}

		token := ""
		_, err := userServices.Login(context.Background(),req.Email,req.Password)

		assert.NotEqual(t, token,usersHashDomain.Token)
		assert.NotEqual(t, req.Password,usersHashDomain.Password)
		assert.NotNil(t, err)
	})
	t.Run("Invalid Test | email unregistered", func(t *testing.T){
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.AnythingOfType("string")).Return(usersEntity.Domain{}, assert.AnError).Once()

		req := usersEntity.Domain{
			Email       : "testinguser@gmail.com",
			Password    : "asdgajkdkjashdkjahsdkasdhjah",
		}

		_, err := userServices.Login(context.Background(),req.Email,req.Password)

		assert.NotEqual(t, req,usersEntity.Domain{})
		assert.NotNil(t, err)
	})
}

func TestUserServices_DeleteUser(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockUsersRepository.On("DeleteUserByUuid", mock.Anything,mock.AnythingOfType("string")).Return("Movie was Deleted", nil).Once()

		resp, err := userServices.DeleteUser(context.Background(),"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, "Movie was Deleted", resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockUsersRepository.On("DeleteUserByUuid", mock.Anything,mock.AnythingOfType("string")).Return("", assert.AnError).Once()

		resp, err := userServices.DeleteUser(context.Background(),"asd;sasjdlkajd")

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestUserServices_FindByUuid(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockUsersRepository.On("GetByUuid",  mock.Anything,mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		resp, err := userServices.FindByUuid(context.Background(),"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, resp, usersDomain)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockUsersRepository.On("GetByUuid",  mock.Anything,mock.AnythingOfType("string")).Return(usersEntity.Domain{}, assert.AnError).Once()

		resp, err := userServices.FindByUuid(context.Background(),"asd;sasjdlkajd")

		assert.NotNil(t, err)
		assert.Equal(t, usersEntity.Domain{}, resp)
	})
}

func TestUserServices_UpdateByIdt(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockUsersRepository.On("UpdateUser",mock.Anything, mock.AnythingOfType("string"),mock.Anything).Return(&usersDomain, nil).Once()

		req := usersEntity.Domain{
			Name        : "Testing User",
			Username    : "testingUser",
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
			PhoneNumber : "08123482342",
			Token:  "asiudgasigdaiusgdiausgdi",
		}
		resp, err := userServices.UpdateById(context.Background(),&req,"tasdasdasdad")

		assert.Nil(t, err)
		assert.Equal(t, usersDomain, *resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockUsersRepository.On("UpdateUser",mock.Anything, mock.AnythingOfType("string"),mock.Anything).Return(&usersEntity.Domain{}, assert.AnError).Once()

		req := usersEntity.Domain{
			Name        : "Testing User",
			Username    : "testingUser",
			Email       : "testinguser@gmail.com",
			Password    : "testing password",
			PhoneNumber : "08123482342",
			Token:  "asiudgasigdaiusgdiausgdi",
		}
		resp, err := userServices.UpdateById(context.Background(),&req,"tasdasdasdad")

		assert.NotNil(t, err)
		assert.Equal(t, usersEntity.Domain{}, *resp)
	})
}

func TestUserServices_UploadAvatar(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockUsersRepository.On("GetByUuid", mock.Anything,mock.AnythingOfType("string")).Return(usersDomain, nil).Once()
		mockUsersRepository.On("UploadAvatar",mock.Anything,mock.Anything,mock.Anything).Return(&usersDomain, nil).Once()

		req := &usersEntity.Domain{
			Avatar: "images/avatar/avatar.png",
		}

		res, err := userServices.UploadAvatar(context.Background(),"kajdnkajsdkjas",req.Avatar)

		assert.Nil(t, err)
		assert.Equal(t, usersDomain.Avatar, res.Avatar)
	})
}