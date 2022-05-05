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

// Translate -.
func (uc *GithubUseCase) GetContributors(ctx context.Context, param model.Request) ([]model.Contributor, error) {
	contributors, err := uc.webAPI.GetContributors(param)
	if err != nil {
		return []model.Contributor{}, fmt.Errorf("GithubUseCase - GetContributors - s.webAPI.Translate: %w", err)
	}

	//err = uc.repo.Store(context.Background(), translation)
	//if err != nil {
	//	return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	//}

	return contributors, nil
}
