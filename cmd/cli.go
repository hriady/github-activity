package cmd

import "fmt"

type CLI struct {
	args           []string
	githubActivity githubActivity
}

type githubActivity interface {
	GetActivity(username string) error
}

func NewCLI(args []string, activity githubActivity) *CLI {
	return &CLI{
		args:           args,
		githubActivity: activity,
	}
}

func (c *CLI) Run() {
	if len(c.args) < 2 {
		fmt.Println("Not enough arguments")
		return
	}

	username := c.args[1]
	c.githubActivity.GetActivity(username)
}
