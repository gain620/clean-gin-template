package usecase

import (
	model "clean-gin-template/internal/model/github"
	"context"
	"fmt"
)

// GithubUseCase -
type GithubUseCase struct {
	webAPI GithubWebAPI
}

// New -
func New(w GithubWebAPI) *GithubUseCase {
	return &GithubUseCase{
		webAPI: w,
	}
}

// GetContributors -.
func (uc *GithubUseCase) GetContributors(ctx context.Context, param model.ContributorRequest) ([]model.ContributorResponse, error) {
	contributors, err := uc.webAPI.GetContributors(param)
	if err != nil {
		return []model.ContributorResponse{}, fmt.Errorf("GithubUseCase - GetContributors : %w", err)
	}

	return contributors, nil
}
