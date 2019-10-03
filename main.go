package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	greenColor  = "42;30m"
	yellowColor = "43;30m"
	resetColor  = "0m"
)

func sendIndexRequest(url string, target *BitbucketPullRequestIndex) error {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", "Bearer "+token)

	response, e := client.Do(request)
	if e != nil {
		return e
	}

	// fmt.Println("Status:", response.StatusCode)

	defer response.Body.Close()

	bodyBytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return e
	}

	e = json.Unmarshal(bodyBytes, target)
	if e != nil && e != io.EOF {
		return e
	}

	return nil
}

func colorText(text, color string) string {
	return fmt.Sprintf("\x1b[%s%s\x1b[%s", color, text, resetColor)
}

func formatReviewerNames(pullRequest BitbucketIndexValue) ([]string, []string) {
	var approvers []string
	var unapprovers []string

	for _, reviewer := range pullRequest.Reviewers {
		if reviewer.Role == "REVIEWER" {
			var coloredUsername string
			username := reviewer.User.Name
			u := strings.ToUpper(" " + string(username[0]) + string(username[3]) + " ")
			if reviewer.Approved {
				coloredUsername = colorText(u, greenColor)
				approvers = append(approvers, coloredUsername)
			} else {
				coloredUsername = colorText(u, yellowColor)
				unapprovers = append(unapprovers, coloredUsername)
			}

		}
	}
	return approvers, unapprovers
}

func main() {

	url := serverHostname + repoPath

	target := &BitbucketPullRequestIndex{}
	err := sendIndexRequest(url, target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, pr := range target.Values {

		approvers, unapprovers := formatReviewerNames(pr)

		fmt.Printf("%s [%d] %s [%s]\n", pr.FromRef.Repository.Name, pr.Id, pr.Title, pr.Author.User.DisplayName)
		fmt.Printf("%s (%2d) %s\n", "UNAPPROVED", len(unapprovers), strings.Join(unapprovers, " "))
		fmt.Printf("%s   (%2d) %s\n", "APPROVED", len(approvers), strings.Join(approvers, " "))
		isReady := len(approvers) > 1

		if isReady {
			fmt.Println(colorText("READY", greenColor))
		} else {
			fmt.Println(colorText("IN PROGRESS", yellowColor))
		}
		fmt.Println()

	}

}
