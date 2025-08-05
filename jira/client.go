package jira

import (
	"exception-reporter-agent/model"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

type Client struct {
	baseURL   string
	authEmail string
	authToken string
	project   string
	issueType string
}

func NewClientFromEnv() *Client {
	return &Client{
		baseURL:   os.Getenv("JIRA_BASE_URL"),
		authEmail: os.Getenv("JIRA_EMAIL"),
		authToken: os.Getenv("JIRA_TOKEN"),
		project:   os.Getenv("JIRA_PROJECT_KEY"),
		issueType: os.Getenv("JIRA_ISSUE_TYPE"),
	}
}

func (c *Client) CreateIssue(bug *model.BugReport) error {
	req := model.JiraIssueRequest{}
	req.Fields.Project.Key = c.project
	req.Fields.Summary = bug.Summary
	req.Fields.Description = bug.Description
	req.Fields.Issuetype.Name = c.issueType
	req.Fields.Labels = bug.Labels
	//req.Fields.Priority = &model.JiraPriority{Name: bug.Priority} // вообще надо вынести в настройки проекта оказывается может не быть такого поля

	client := resty.New()
	resp, err := client.R().
		SetBasicAuth(c.authEmail, c.authToken).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post(c.baseURL + "/rest/api/2/issue")

	//fmt.Println(resp.String()) // отладка

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("jira error: %s", resp.Body())
	}

	//fmt.Println(resp.String()) //  отладка
	return nil
}
