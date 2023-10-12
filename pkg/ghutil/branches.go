package ghutil

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v55/github"
)

func (org *Organization) listBranches(repository string) []*github.Branch {
	var allBranches []*github.Branch

	perPage := 100
	opts := &github.BranchListOptions{
		ListOptions: github.ListOptions{
			PerPage: perPage,
		},
	}

	fmt.Print("Fetching branches")
	for {
		branches, resp, err := org.client.Repositories.ListBranches(context.Background(), org.name, repository, opts)
		must(err)
		fmt.Print(".")

		allBranches = append(allBranches, branches...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	fmt.Println()

	return allBranches
}

func (org *Organization) CountBranches(repository string) {
	branches := org.listBranches(repository)
	fmt.Printf("Total branches: %d\n", len(branches))
}

func (org *Organization) DeleteStale(repository string) {
	branches := org.listBranches(repository)

	sixMonthAgo := time.Now().Add(-1 * 4 * 30 * 24 * time.Hour)

	fmt.Println("Deleting branches")
	for _, branch := range branches {
		branchDetails, _, err := org.client.Repositories.GetBranch(context.Background(), org.name, repository, *branch.Name, false)
		must(err)

		if branchDetails.Commit.Commit.Author.GetDate().Before(sixMonthAgo) {
			fmt.Println(branchDetails.GetName(), branchDetails.Commit.Commit.Author.GetDate())
			org.client.Git.DeleteRef(context.Background(), org.name, repository, "heads/"+branchDetails.GetName())
		}
	}
}
