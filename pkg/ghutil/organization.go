package ghutil

import "github.com/google/go-github/v55/github"

type Organization struct {
	name   string
	client *github.Client
}

func CreateOrganization(token string, name string) *Organization {
	client := github.NewClient(nil).WithAuthToken(token)

	return &Organization{
		name:   name,
		client: client,
	}
}
