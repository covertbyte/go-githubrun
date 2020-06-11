# go-githubrun

go-githubrun is a Go library for parsing the event payload and default environment variables available to the Github
workflow run instance.

## Usage

```go
import (
    "github.com/covertbyte/go-githubrun"
    "github.com/google/go-github/v32/github"
    "github.com/sethvargo/go-githubactions"
)

func main() {
    run, err := githubrun.ParseRun()
    if err != nil {
        githubactions.Fatalf("%w", err)
    }

    ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: run.Env.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

    pullRequest := run.Payload.(*github.PullRequestEvent).GetPullRequest()

    githubactions.DebugF("%s", pullRequest.GetBody())

    githubactions.DebugF("%s", run.Env.GithubEventName)
}
```

For API documentation see the [Go docs](https://godoc.org/github.com/covertbyte/go-githubrun).
