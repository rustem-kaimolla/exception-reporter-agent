# Exception Reporter Agent

Exception Reporter Agent is a lightweight TCP server designed to receive and log exception reports from applications (e.g., Laravel) via a custom binary or JSON-based protocol. It is intended to be used in conjunction with a Laravel package (like laravel-exception-reporter) that handles the client-side reporting.

## Features

- Accepts incoming TCP connections on a configurable port
- Parses and stores incoming exception reports
- Lightweight, fast, and deployable via Docker
- Easy integration with Laravel or any custom application

## Example Use Case

You can use this agent together with [`laravel-exception-reporter`](https://github.com/rustem-kaimolla/laravel-exception-reporter) to automatically report Laravel exceptions to your Jira project.

## Setup

### 1. Build Docker image

```bash
docker build -t rustemkaimolla/exception-reporter-agent .
```

### 2. Run

```bash
docker run --env-file .env \
  -p 9000:9000 \
  rustemkaimolla/exception-reporter-agent
```

### 3. Send a test report

```bash
echo '{"project":"FM","env":"production","exception":{"class":"RuntimeException","message":"Something went wrong","file":"/app/Service.php","line":42,"trace":[]}}' | nc localhost 9000
```

## Environment Variables

| Variable           | Description                                  |
| ------------------ | -------------------------------------------- |
| `OPENAI_API_KEY`   | Open AI API key                              |
| `OPENAI_MODEL`     | Open AI model for analyze. Default gpt4o-mini|
| `JIRA_BASE_URL`    | Your Jira base URL (e.g., `https://...`)     |
| `JIRA_EMAIL`       | Jira username (email)                        |
| `JIRA_TOKEN`       | Jira API token                               |
| `JIRA_PROJECT_KEY` | Key of the target Jira project               |
| `JIRA_ISSUE_TYPE`  | Issue type to create (e.g., `Bug`)           |

## License
MIT License. See LICENSE for more details.

## Contributions
Pull requests are welcome. Please open issues for feature suggestions or bugs.
