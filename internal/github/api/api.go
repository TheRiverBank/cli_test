package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RepositoryResponse struct {
	StargazersCount int64 `json:"stargazers_count"`
}

func GetStarsCount(repositoryName string) (*int64, error) {
	response, err := http.Get("https://api.github.com/repos/" + repositoryName)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseStruct RepositoryResponse
	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct.StargazersCount, nil
}
