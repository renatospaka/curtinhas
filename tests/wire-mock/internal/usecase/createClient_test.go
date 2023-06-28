package usecase_test

import (
	"testing"

	"github.com/renatospaka/tests/internal/entity"
	"github.com/renatospaka/tests/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientRepositoryMock struct {
	mock.Mock
}

func (c *ClientRepositoryMock) Save(client *entity.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

func TestCreateClientUsecase(t *testing.T) {
	mockRepo := new(ClientRepositoryMock)
	mockRepo.On("Save", mock.Anything).Return(nil)

	in := usecase.CreateClientInputDTO{
		Name:  "John Doe",
		Email: "jdoe@notemail.com",
	}
	createClientUsecase := usecase.NewCreateClientUsecase(mockRepo)
	out, err := createClientUsecase.Execute(in)
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", out.Name)
	assert.Equal(t, "jdoe@notemail.com", out.Email)
	assert.Equal(t, 0, out.Points)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "Save", 1)
}
