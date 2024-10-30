package utils

type UserEvent struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		Id           int64 `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
	} `json:"actor"`

	Repo struct {
		Id   int64 `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`

	Payload struct {
		RepositoryID int64 `json:"repository_id"`
		PushID int64 `json:"push_id"`
		Action string `json:"action"`

		Commits [] struct {
			Message string `json:"message"`
		} `json:"commits"`

		PullRequest struct {
			CreatedAt string `json:"created_at"`
			ClosedAt string `json:"closed_at"`
			MergedAt string `json:"merged_at"`
		} `json:"pull_request"`

		Issue struct {
			Title string `json:"title"`
			User struct {
				Login string `json:"login"`
			}
			State string `json:"state"`
		} `json:"issue"`

		Release struct {
			Name string `json:"name"`
			TagName string `json:"tag_name"`
			PublishedAt string `json:"published_at"`
		} `json:"release"`
	} `json:"payload"`

	CreatedAt string `json:"created_at"`
}


func ArgsChecker(args []string) bool {
	return (len(args) > 1) && args[1] != ""
}