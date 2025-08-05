package model

type BugReport struct {
	Fingerprint string `json:"fingerprint"`
	App         string `json:"app"`
	Env         string `json:"env"`

	Summary     string `json:"summary"`
	Description string `json:"description"`
	Priority    string `json:"priority"`

	Timestamp string   `json:"timestamp"`
	Labels    []string `json:"labels,omitempty"`
}
