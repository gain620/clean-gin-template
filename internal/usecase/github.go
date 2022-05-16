package usecase

import (
	model "clean-gin-template/internal/model/github"
	"context"
	"fmt"
)

// MovieUseCase -
type GithubUseCase struct {
	webAPI GithubWebAPI
}

// NewGithubUseCse -
func New(w GithubWebAPI) *GithubUseCase {
	return &GithubUseCase{
		webAPI: w,
	}
}

// GetContributors -.
func (uc *GithubUseCase) GetContributors(ctx context.Context, param model.Request) ([]model.Contributor, error) {
	contributors, err := uc.webAPI.GetContributors(param)
	if err != nil {
		return []model.Contributor{}, fmt.Errorf("GithubUseCase - GetContributors : %w", err)
	}

	return contributors, nil
}
