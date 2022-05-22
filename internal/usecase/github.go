package usecase

import (
	"clean-gin-template/internal/model/github/dto"
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
func (uc *GithubUseCase) GetContributors(ctx context.Context, param dto.ContributorRequest) ([]dto.ContributorResponse, error) {
	contributors, err := uc.webAPI.GetContributors(param)
	if err != nil {
		return []dto.ContributorResponse{}, fmt.Errorf("GithubUseCase - GetContributors : %w", err)
	}

	return contributors, nil
}
