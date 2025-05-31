package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Github struct {
	httpClient httpClient
}

type httpClient interface {
	Get(newURL string, params map[string]string) (*http.Response, error)
}

func NewGithub(httpClient httpClient) *Github {
	return &Github{
		httpClient: httpClient,
	}
}

func (g *Github) GetActivity(username string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := g.httpClient.Get(url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	githubActivities := []GithubActivity{}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(bytes, &githubActivities)
	if err != nil {
		return err
	}
	
	for _, activity := range githubActivities {
		g.displayActivity(activity)
	}
	
	return nil
}

func (g *Github) displayActivity(activity GithubActivity) {
	switch activity.Type {
	case CommitCommentEvent:
		fmt.Printf("Commented on commit %s\n", activity.Payload.Comment)
	case CreateEvent:
		fmt.Printf("Created a new %s\n", activity.Payload.RefType)
	case DeleteEvent:
		fmt.Printf("Deleted a %s\n", activity.Payload.RefType)
	case ForkEvent:
		fmt.Printf("Forked %s\n", activity.Repo.Name)
	case GollumEvent:
		fmt.Printf("Create / Edited a wiki page\n")
	case IssueCommentEvent:
		fmt.Printf("%s a comment in issue\n", activity.Payload.Action)
	case IssuesEvent:
		fmt.Printf("%s a issue\n", activity.Payload.Action)
	case MemberEvent:
		fmt.Printf("Added %s to %s\n", activity.Actor.Login, activity.Repo.Name)
	case PublicEvent:
		fmt.Printf("Made %s public\n", activity.Repo.Name)
	case PullRequestEvent:
		fmt.Printf("%s a pull request\n", activity.Payload.Action)
	case PullRequestReviewEvent:
		fmt.Printf("Reviewed a pull request\n")
	case PullRequestReviewCommentEvent:
		fmt.Printf("Commented on a pull request\n")
	case PullRequestReviewThreadEvent:
		fmt.Printf("%s on a pull request thread\n", activity.Payload.Action)
	case PushEvent:
		fmt.Printf("Pushed %d commits to %s\n", activity.Payload.Size, activity.Repo.Name)
	case ReleaseEvent:
		fmt.Printf("Published a new released for %s\n", activity.Repo.Name)
	case WatchEvent:
		fmt.Printf("Starred %s\n", activity.Repo.Name)
	default:
		fmt.Printf("Unknown activity type %s\n", activity.Type)
	}
}
