package ghutil

import (
	"context"
	"fmt"

	"github.com/google/go-github/v55/github"
)

func (org *Organization) ListMembers() {
	var allMembers []*github.User

	perPage := 100
	opts := &github.ListMembersOptions{
		ListOptions: github.ListOptions{
			PerPage: perPage,
		},
	}

	fmt.Print("Fetching Users")
	for {
		members, resp, err := org.client.Organizations.ListMembers(context.Background(), org.name, opts)
		must(err)
		fmt.Print(".")

		allMembers = append(allMembers, members...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	fmt.Println()

	for _, member := range allMembers {
		user, _, err := org.client.Users.Get(context.Background(), member.GetLogin())
		must(err)

		fmt.Printf("%s, %s\n", user.GetLogin(), user.GetName())
	}
}
