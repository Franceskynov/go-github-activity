package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Franceskynov/go-github-activity/utils"
)

func GetRawData(userName string) (io.ReadCloser, error) {

	// Interpolating the user into the URL
	url := fmt.Sprintf("https://api.github.com/users/%s/events", userName)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		return response.Body, err
	case http.StatusNotFound:
		return nil, errors.New("are you typing the correct user name?")
	default:
		return nil, fmt.Errorf("errors on the server with status: %d", response.StatusCode)
	}
}

/*
*
 */
func GetUserData(userName string) []utils.UserEvent {

	var userEvents []utils.UserEvent
	raw, err := GetRawData(userName)

	if err != nil {
		fmt.Println(err)
	} else {
		
		decoder := json.NewDecoder(raw)
		err = decoder.Decode(&userEvents)

		if err != nil {
			fmt.Println(err)
		}
	}

	return userEvents
}

func ShowUserEvents(userName string, event utils.UserEvent) {
	
	switch event.Type {
		case "ReleaseEvent":	
			fmt.Printf("> Tagged a version %s of: %s published at: %s name: %s\n", event.Payload.Release.TagName, event.Repo.Name, event.Payload.Release.PublishedAt, event.Payload.Release.Name)
		case "PushEvent":
			fmt.Printf("> Pushed %d commit(s) to: %s at: %s \n", len(event.Payload.Commits), event.Repo.Name, event.CreatedAt)
		case "WatchEvent":
			if event.Payload.Action == "started" {
				fmt.Printf("> User ⭐ to repo %s \n", event.Repo.Name)
			}
		case "IssueCommentEvent":
			fmt.Println("-------------------------------------")
			fmt.Printf("> Opened a new issue in: %s \nby: %s \ntitle: %s \nstatus: %s \n", event.Repo.Name, event.Payload.Issue.User.Login, event.Payload.Issue.Title, event.Payload.Issue.State)
			fmt.Println("-------------------------------------")
		case "PullRequestEvent":
			if event.Payload.Action == "opened" {
				fmt.Printf("> Opened a pull request in: %s created at: %s \n", event.Repo.Name, event.Payload.PullRequest.CreatedAt)
			}
			
			if event.Payload.Action == "closed" {
				fmt.Printf("> Closed a pull request in: %s closed at: %s \n", event.Repo.Name, event.Payload.PullRequest.ClosedAt)
			}
	}
}

func FormatUserData(userName string) {

	userEvents := GetUserData(userName)
	if len(userEvents) > 0 {
		for _, event := range userEvents {
			ShowUserEvents(userName, event)
		}
	}
}