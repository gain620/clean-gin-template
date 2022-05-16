// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	model "clean-gin-template/internal/model/github"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Github -.
	Github interface {
		GetContributors(context.Context, model.Request) ([]model.Contributor, error)
	}

	// GithubWebApi -.
	GithubWebAPI interface {
		GetContributors(model.Request) ([]model.Contributor, error)
	}
)
