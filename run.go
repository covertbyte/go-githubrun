package githubrun

import (
	"github.com/google/go-github/v32/github"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Run represents a Github run.
//
// https://help.github.com/en/actions/reference/events-that-trigger-workflows
type Run struct {
	Env        Env
	Payload    interface{}
	Owner      string
	Repository string
}

// Env contains the default environment variables.
//
// https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
// https://help.github.com/en/actions/configuring-and-managing-workflows/authenticating-with-the-github_token
type Env struct {
	CI               bool
	Home             string
	GithubWorkflow   string
	GithubRunID      int64
	GithubRunNumber  int64
	GithubAction     string
	GithubActions    bool
	GithubActor      string
	GithubRepository string
	GithubEventName  string
	GithubEventPath  string
	GithubWorkspace  string
	GithubSHA        string
	GithubRef        string
	GithubHeadRef    string
	GithubBaseRef    string
	GithubToken      string
}

// ParseRun parses the event payload and default environment variables. A struct containing the env variables,
// and payload of the corresponding struct type will be returned.
func ParseRun() (Run, error) {
	r := new(Run)

	ci, err := strconv.ParseBool(os.Getenv("CI"))
	if err != nil {
		return Run{}, err
	}
	r.Env.CI = ci
	r.Env.Home = os.Getenv("HOME")
	r.Env.GithubWorkflow = os.Getenv("GITHUB_WORKFLOW")
	runID, err := strconv.ParseInt(os.Getenv("GITHUB_RUN_ID"), 10, 64)
	if err != nil {
		return Run{}, err
	}
	r.Env.GithubRunID = runID
	runNumber, err := strconv.ParseInt(os.Getenv("GITHUB_RUN_NUMBER"), 10, 64)
	if err != nil {
		return Run{}, err
	}
	r.Env.GithubRunNumber = runNumber
	r.Env.GithubAction = os.Getenv("GITHUB_ACTION")
	actions, err := strconv.ParseBool(os.Getenv("GITHUB_ACTIONS"))
	if err != nil {
		return Run{}, err
	}
	r.Env.GithubActions = actions
	r.Env.GithubActor = os.Getenv("GITHUB_ACTOR")
	r.Env.GithubRepository = os.Getenv("GITHUB_REPOSITORY")
	r.Env.GithubEventName = os.Getenv("GITHUB_EVENT_NAME")
	r.Env.GithubEventPath = os.Getenv("GITHUB_EVENT_PATH")
	r.Env.GithubWorkspace = os.Getenv("GITHUB_WORKSPACE")
	r.Env.GithubSHA = os.Getenv("GITHUB_SHA")
	r.Env.GithubRef = os.Getenv("GITHUB_REF")
	r.Env.GithubHeadRef = os.Getenv("GITHUB_HEAD_REF")
	r.Env.GithubBaseRef = os.Getenv("GITHUB_BASE_REF")
	r.Env.GithubToken = os.Getenv("GITHUB_TOKEN")

	// Split "owner/repository" string into more useful owner and repository values
	repository := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
	r.Owner = repository[0]
	r.Repository = repository[1]

	// Parse the webhook event payload, For recognized event types,
	// a value of the corresponding struct type will be returned.
	jsonFile, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return Run{}, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	payload, err := github.ParseWebHook(os.Getenv("GITHUB_EVENT_NAME"), byteValue)
	if err != nil {
		return Run{}, err
	}
	r.Payload = payload
	return *r, nil
}
