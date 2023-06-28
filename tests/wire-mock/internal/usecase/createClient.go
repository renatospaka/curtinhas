package usecase

import "github.com/renatospaka/tests/internal/entity"

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID     string
	Name   string
	Email  string
	Points int
}

type CreateClientUsecase struct {
	ClientRepository entity.ClientRepositoryInterface
}

func NewCreateClientUsecase(repo entity.ClientRepositoryInterface) *CreateClientUsecase {
	return &CreateClientUsecase{
		ClientRepository: repo,
	}
}

func (c *CreateClientUsecase) Execute(in CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(in.Name, in.Email)
	if err != nil {
		return nil, err
	}

	err = c.ClientRepository.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:     client.ID,
		Name:   client.Name,
		Email:  client.Email,
		Points: client.Points,
	}, nil
}
