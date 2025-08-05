package handler

import (
	"encoding/json"
	"exception-reporter-agent/jira"
	"exception-reporter-agent/model"
	"exception-reporter-agent/pkg/cache"
	"exception-reporter-agent/pkg/filter"
	"exception-reporter-agent/pkg/fingerprint"
	"exception-reporter-agent/pkg/llm"
	"exception-reporter-agent/pkg/report"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

var store = cache.NewDuplicateStore(7 * 24 * time.Hour) // надо либо в файл выгружать или в редис но это накладно пока inmemmory

func HandleException(data []byte) {
	var exc model.ExceptionPayload

	if err := json.Unmarshal(data, &exc); err != nil {
		log.Printf("Invalid JSON: %v", err)
		return
	}

	fp := fingerprint.Generate(exc)
	if store.Seen(fp) {
		log.Printf("Duplicate exception (ignored): %s", fp)
		return
	}

	if filter.ShouldIgnore(exc) {
		log.Printf("Ignored by filter: %s — %s", fp, exc.Message)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	apiModel := os.Getenv("OPENAI_MODEL")

	if apiKey == "" {
		log.Fatal("❌ OPENAI_API_KEY is not set")
		return
	}
	if apiModel == "" {
		log.Fatal("❌ OPENAI_MODEL is not set")
		return
	}

	llmClient := llm.NewClient(apiKey, apiModel)
	llmResp, err := llmClient.Process(context.Background(), exc)
	if err != nil {
		log.Printf("LLM error: %v", err)
		return
	}

	bugReport := report.Generate(exc, llmResp)

	//log.Printf("[BUG REPORT]\nSummary: %s\nPriority: %s\nFingerprint: %s",
	//	bugReport.Summary, bugReport.Priority, bugReport.Fingerprint,
	//)

	jiraClient := jira.NewClientFromEnv()
	_ = jiraClient.CreateIssue(bugReport)
}
