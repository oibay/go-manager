package usecase

import (
	"context"
	"github.com/abylq/learning-management-system/components/auth/models"
	"github.com/abylq/learning-management-system/components/auth/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)

	uc := NewAuthUseCase(repo, "salt", []byte("secret"),86400)

	var (
		username = "user"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87",
		}
	)

	//SignUp
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx,username, password)
	assert.NoError(t, err)

	//SignIn

	repo.On("GetUser", user.Username, user.Password).Return(user,nil)
	token, err := uc.SignIn(ctx, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	//Verify Token

	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
