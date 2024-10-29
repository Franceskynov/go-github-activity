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
		break
	case http.StatusNotFound:
		return nil, errors.New("are you typing the correct user name?")
	default:
		return nil, fmt.Errorf("errors on the server with status: %d", response.StatusCode)
	}

	return response.Body, err
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

func FormatUserData(userName string) {

	userEvents := GetUserData(userName)
	if len(userEvents) > 0 {
		for _, event := range userEvents {

			pushEvents := 0
			issueCommentEvent := 0
			pullRequestEventOpen := 0
			pullRequestEventClosed := 0

			switch event.Type {
				// case "ReleaseEvent":	

				case "PushEvent":
					pushEvents = pushEvents + 1

				case "WatchEvent":
					if event.Payload.Action == "started" {
					
						fmt.Printf("> ⭐ to repo %s \n", event.Repo.Name)
					}
				case "IssueCommentEvent":
					fmt.Printf("> Opened a new issue in: %s \nby: %s \ntitle: %s \nstatus: %s \n", event.Repo.Name, event.Payload.Issue.User.Login, event.Payload.Issue.Title, event.Payload.Issue.State)
					fmt.Println("----------------------------------")
					issueCommentEvent = issueCommentEvent + 1
				case "PullRequestEvent":
					if event.Payload.Action == "opened" {
						pullRequestEventOpen = pullRequestEventOpen + 1
						fmt.Printf("> User has opened a pull request in: %s \ncreated at: %s \n", event.Repo.Name, event.Payload.PullRequest.CreatedAt)
						fmt.Println("----------------------------------")
					}
					
					if event.Payload.Action == "closed" {
						pullRequestEventClosed = pullRequestEventClosed + 1
						fmt.Printf("> User has closed a pull request in: %s \nclosed at: %s \n", event.Repo.Name, event.Payload.PullRequest.ClosedAt)
						fmt.Println("----------------------------------")
					}
			}

			// if pushEvents > 0 {
			// 	fmt.Printf("User push events %d to repo %s \n", pushEvents, event.Repo.Name)
			// }

	

			// if pullRequestEventOpen > 0 {
			// 	fmt.Printf("Number of pull request opened: %d \n", pullRequestEventOpen)
			// }

			// if pullRequestEventClosed > 0 {
			// 	fmt.Printf("Number of pull request closed: %d \n", pullRequestEventClosed)
			// }
		}
	}


}