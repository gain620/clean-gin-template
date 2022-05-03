package usecase

import (
	"context"
	"fmt"
)

// MovieUseCase -.
// separate to entity?
type MovieUseCase interface {
	Translate(context.Context, entity.Translation) (entity.Translation, error)
	History(context.Context) ([]entity.Translation, error)
}

// MovieUseCase -
type movieUseCase struct {
	//repo   MovieRepo
	webAPI MovieWebAPI
}

// NewMovieUseCase -
func NewMovieUseCase(w MovieWebAPI) *MovieUseCase {
	return &movieUseCase{
		//repo:   r,
		webAPI: w,
	}
}

//// History - getting translate history from store.
//func (uc *movieUseCase) History(ctx context.Context) ([]entity.Translation, error) {
//	translations, err := uc.repo.GetHistory(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
//	}
//
//	return translations, nil
//}
//
//// Translate -.
//func (uc *movieUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
//	translation, err := uc.webAPI.Translate(t)
//	if err != nil {
//		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
//	}
//
//	err = uc.repo.Store(context.Background(), translation)
//	if err != nil {
//		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
//	}
//
//	return translation, nil
//}

// GetScore -.
func (uc *movieUseCase) GetScore(ctx context.Context, t entity.Translation) (entity.Translation, error) {
	translation, err := uc.webAPI.Translate(t)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
	}

	err = uc.repo.Store(context.Background(), translation)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return translation, nil
}
