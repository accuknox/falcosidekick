package outputs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/andygrunwald/go-jira"
)

// PrettyString formats the string
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

// SendMsgtoJira to send message to jira integration
func SendMsgtoJira(ctx context.Context, value string, result []byte) error {
	splitval := strings.Split(value, ",")

	sendAlert := string(result)

	resJson, err := PrettyString(sendAlert)
	if err != nil {
		fmt.Println("Err in creating pretty res_json in Jira " + err.Error())
		return errors.New("Err in creating pretty res_json in Jira")
	}
	if len(splitval) == 7 {

		jiraClient, err := createClient(splitval[4], splitval[5], splitval[1])
		if err != nil {
			fmt.Println("Err in creating a new Jira Client " + err.Error())
			return errors.New("Err in creating a new Jira Client ")
		}

		i := jira.Issue{
			Fields: &jira.IssueFields{
				Description: resJson,
				Type: jira.IssueType{
					Name: splitval[3],
				},
				Project: jira.Project{
					Key: splitval[2],
				},
				Summary: splitval[0],
			},
		}
		issue, _, err := jiraClient.Issue.Create(&i)
		if err != nil {
			fmt.Println("Failed to create Jira Ticket " + err.Error())
			return errors.New("Failed to create Jira Ticket ")
		}
		fmt.Println("Jira Ticket Created Successfully :> %v", issue)
		return nil
	}
	fmt.Println("unable to get the required value in send msg to jira ")
	return errors.New("unable to get the required value for sending message to jira ")
}

func createClient(userEmail, token, siteUrl string) (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: userEmail,
		Password: token,
	}

	jiraClient, err := jira.NewClient(tp.Client(), siteUrl)
	if err != nil {
		fmt.Println("Err in creating a new Jira Client " + err.Error())
		return nil, errors.New("Err in creating a new Jira Client ")
	}
	return jiraClient, nil
}

// ValidateToken Validate token that is being sent from UI
func ValidateToken(userEmail, token, siteUrl, project string) bool {
	jiraClient, err := createClient(userEmail, token, siteUrl)
	if err != nil {
		return false
	}
	if _, _, err := jiraClient.Project.Get(project); err != nil {
		return false
	}
	return true
}
