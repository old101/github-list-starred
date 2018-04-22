package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

func main() {
	options := &github.ActivityListStarredOptions{Sort: "created"}

	if len(os.Args) <= 1 {
		log.Fatalf("Usage: %s <GitHub username>", os.Args[0])
	}
	user := os.Args[1]

	client := newClient()

	for page := 1; ; page++ {
		time.Sleep(700 * time.Millisecond)
		options.Page = page
		options.PerPage = 100

		starredRepos, res, err := client.Activity.ListStarred(context.Background(), user, options)
		if err != nil {
			log.Fatalf("ListStarred: %s", err)
		}

		log.Printf("page: %d/%d", page, res.LastPage)
		for _, starredRepo := range starredRepos {
			fmt.Println(*starredRepo.Repository.HTMLURL)
		}

		if page >= res.LastPage {
			break
		}
	}
}

func newClient() *github.Client {
	githubToken := os.Getenv("GITHUB_TOKEN")

	if githubToken != "" {
		source := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
		return github.NewClient(oauth2.NewClient(context.Background(), source))
	}

	return github.NewClient(nil)
}
