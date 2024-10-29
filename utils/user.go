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
		Commits [] struct {

		} `json:"commits"`
	} `json:"payload"`

	CreatedAt string `json:"created_at"`
}


func ArgsChecker(args []string) bool {
	return (len(args) > 1) && args[1] != ""
}