package main

import (
	"github-activity/cmd"
	"github-activity/github"
	"github-activity/http_client"
	"os"
)

func main() {
	httpClient := http_client.NewHttpClient()
	githubActivity := github.NewGithub(httpClient)
	cli := cmd.NewCLI(os.Args, githubActivity)
	cli.Run()
}
