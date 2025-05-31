package github

type GithubActivity struct {
	ID        string             `json:"id"`
	Type      GithubActivityType `json:"type"`
	Actor     GithubActor        `json:"actor"`
	Repo      GithubRepository   `json:"repo"`
	Payload   GithubPayload      `json:"payload"`
	Public    bool               `json:"public"`
	CreatedAt string             `json:"created_at"`
}

type GithubActivityType string

const (
	CommitCommentEvent            GithubActivityType = "CommitCommentEvent"
	CreateEvent                   GithubActivityType = "CreateEvent"
	DeleteEvent                   GithubActivityType = "DeleteEvent"
	ForkEvent                     GithubActivityType = "ForkEvent"
	GollumEvent                   GithubActivityType = "GollumEvent"
	IssueCommentEvent             GithubActivityType = "IssueCommentEvent"
	IssuesEvent                   GithubActivityType = "IssuesEvent"
	MemberEvent                   GithubActivityType = "MemberEvent"
	PublicEvent                   GithubActivityType = "PublicEvent"
	PullRequestEvent              GithubActivityType = "PullRequestEvent"
	PullRequestReviewEvent        GithubActivityType = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent GithubActivityType = "PullRequestReviewCommentEvent"
	PullRequestReviewThreadEvent  GithubActivityType = "PullRequestReviewThreadEvent"
	PushEvent                     GithubActivityType = "PushEvent"
	ReleaseEvent                  GithubActivityType = "ReleaseEvent"
	WatchEvent                    GithubActivityType = "WatchEvent"
)

type GithubActor struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type GithubRepository struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GithubPayload struct {
	RepositoryID int64            `json:"repository_id"`
	PushID       int64            `json:"push_id"`
	Size         int64            `json:"size"`
	DistinctSize int64            `json:"distinct_size"`
	Forkee       GithubRepository `json:"forkee"`
	Ref          string           `json:"ref"`
	RefType      string           `json:"ref_type"`
	MasterBranch string           `json:"master_branch"`
	Description  string           `json:"description"`
	PusherType   string           `json:"pusher_type"`
	Head         string           `json:"head"`
	Before       string           `json:"before"`
	Commits      []GithubCommit   `json:"commits"`
	Action       string           `json:"action"`
	Comment      string           `json:"comment"`
	Pages        []GithubPage     `json:"pages"`
}

type GithubPage struct {
	PageName string `json:"page_name"`
	Title    string `json:"title"`
	Action   string `json:"action"`
	Sha      string `json:"sha"`
	HtmlURL  string `json:"html_url"`
}

type GithubCommit struct {
	Sha      string       `json:"sha"`
	Author   GithubAuthor `json:"author"`
	Message  string       `json:"message"`
	Distinct bool         `json:"distinct"`
	URL      string       `json:"url"`
}

type GithubAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
