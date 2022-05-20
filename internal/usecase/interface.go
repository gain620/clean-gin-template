// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"clean-gin-template/internal/model/github/dto"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Github -.
	Github interface {
		GetContributors(context.Context, dto.ContributorRequest) ([]dto.ContributorResponse, error)
	}

	// GithubWebAPI -.
	GithubWebAPI interface {
		GetContributors(request dto.ContributorRequest) ([]dto.ContributorResponse, error)
	}

	// UserRepo -.
	UserRepo interface {
		GetUsers() error
	}
)
