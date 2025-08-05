package report

import (
	"exception-reporter-agent/model"
	"exception-reporter-agent/pkg/fingerprint"
)

func Generate(exc model.ExceptionPayload, llmResp *model.LLMResponse) *model.BugReport {
	return &model.BugReport{
		Fingerprint: fingerprint.Generate(exc),
		App:         exc.App,
		Env:         exc.Env,
		Summary:     llmResp.Summary,
		Description: llmResp.Description,
		Priority:    llmResp.Priority,
		Timestamp:   exc.Timestamp,
		Labels: []string{
			"auto-created",
			"exception-reporter",
			exc.Env,
			exc.App,
		},
	}
}
