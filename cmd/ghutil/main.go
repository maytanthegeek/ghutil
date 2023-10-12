package main

import (
	"fmt"
	"os"
	"time"

	ghutil "github.com/maytanthegeek/ghutil/pkg/ghutil"
	"github.com/urfave/cli/v2"
)

func init() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Fprintf(cCtx.App.Writer, "%s\n", cCtx.App.Version)
	}
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		fmt.Println("No auth token was passed; exiting.")
		os.Exit(0)
	}

	var githubOrg *ghutil.Organization
	var organization string
	var repository string

	app := &cli.App{
		Name:     "ghutil",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "Tanmay",
			},
		},
		Usage: "swiss knife for common GitHub tasks",
		Commands: []*cli.Command{
			{
				Name:        "delete-stale-branches",
				Aliases:     []string{"ds"},
				Category:    "branch",
				Description: "deletes branches older than 6 months",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "org", Aliases: []string{"o"}, Destination: &organization, Required: true},
					&cli.StringFlag{Name: "repo", Aliases: []string{"r"}, Destination: &repository, Required: true},
				},
				Before: func(cCtx *cli.Context) error {
					githubOrg = ghutil.CreateOrganization(token, organization)
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					githubOrg.DeleteStale(repository)
					return nil
				},
			},
			{
				Name:        "count-branches",
				Aliases:     []string{"cb"},
				Category:    "branch",
				Description: "count all branches",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "org", Aliases: []string{"o"}, Destination: &organization, Required: true},
					&cli.StringFlag{Name: "repo", Aliases: []string{"r"}, Destination: &repository, Required: true},
				},
				Before: func(cCtx *cli.Context) error {
					githubOrg = ghutil.CreateOrganization(token, organization)
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					githubOrg.CountBranches(repository)
					return nil
				},
			},
			{
				Name:        "list-members",
				Aliases:     []string{"lm"},
				Category:    "users",
				Description: "list all members of organization",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "org", Aliases: []string{"o"}, Destination: &organization, Required: true},
				},
				Before: func(cCtx *cli.Context) error {
					githubOrg = ghutil.CreateOrganization(token, organization)
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					githubOrg.ListMembers()
					return nil
				},
			},
		},
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("\n%s\n", err)
	}
}
