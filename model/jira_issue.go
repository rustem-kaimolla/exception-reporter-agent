package model

type JiraIssueRequest struct {
	Fields struct {
		Project     JiraProject   `json:"project"`
		Summary     string        `json:"summary"`
		Description string        `json:"description"`
		Issuetype   JiraIssueType `json:"issuetype"`
		Labels      []string      `json:"labels,omitempty"`
		Priority    *JiraPriority `json:"priority,omitempty"`
	} `json:"fields"`
}

type JiraProject struct {
	Key string `json:"key"`
}

type JiraIssueType struct {
	Name string `json:"name"`
}

type JiraPriority struct {
	Name string `json:"name"`
}
