package webapi

import (
	"clean-gin-template/internal/model/github/dto"
	"fmt"
	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"time"
)

// GithubWebAPI -.
type GithubWebAPI struct {
	conf   githubWebAPIConfig
	client *req.Client
}

// githubWebAPIConfig is for client configurations.
type githubWebAPIConfig struct {
	UserAgent string
	Timeout   time.Duration
}

// New creates a new GithubWebAPI client.
func New() *GithubWebAPI {
	conf := &githubWebAPIConfig{
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
		Timeout:   10 * time.Second,
	}

	return &GithubWebAPI{
		client: req.C().
			SetUserAgent(conf.UserAgent).
			SetTimeout(conf.Timeout),
	}
}

// GetContributors -.
func (t *GithubWebAPI) GetContributors(param dto.ContributorRequest) ([]dto.ContributorResponse, error) {
	var contributors []dto.ContributorResponse
	//errMsg := &model.ErrorMessage{}

	resp, err := t.client.R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetPathParams(map[string]string{
			"owner": param.Owner,
			"repo":  param.Repo,
		}).
		Get("https://api.github.com/repos/{owner}/{repo}/contributors")

	if err != nil {
		log.Warn(err)
		return contributors, fmt.Errorf("GithubWebAPI - coudl't get contributors: %w", err)
	}

	if resp.IsSuccess() {
		err = resp.Unmarshal(&contributors)
		if err != nil {
			log.Warn(err)
		}
	} else {
		log.Warn("bad response:", resp)
	}

	return contributors, nil
}
