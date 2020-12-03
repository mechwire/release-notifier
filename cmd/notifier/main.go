package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jncmaguire/labelle-release-notifier/internal/github"
	"github.com/jncmaguire/labelle-release-notifier/internal/slack"
	"github.com/jncmaguire/labelle-release-notifier/internal/util"
)

type args struct {
	gitHubAction github.Action
	gitHub       github.Client
	slack        slack.Client
}

func (a args) validate() (err error) {
	for _, val := range []string{
		a.gitHubAction.Actor,
		a.gitHubAction.Repository,
		a.gitHubAction.Ref,
		a.gitHubAction.Event,
		a.gitHubAction.ServerURL,
		a.gitHub.APIURL,
		a.gitHubAction.Activity,
		a.slack.APIURL,
		a.slack.APIToken,
		a.slack.Channel,
	} {
		if val == "" {
			err = errors.New("Value should not be empty")
		}
	}

	return err
}

func getEnvArgs() args {
	return args{
		gitHubAction: github.Action{
			Actor:      os.Getenv(`GITHUB_ACTOR`),
			Repository: os.Getenv(`GITHUB_REPOSITORY`),
			Ref:        os.Getenv(`GITHUB_REF`),
			Event:      os.Getenv(`GITHUB_EVENT`),
			Activity:   os.Getenv(`GITHUB_EVENT_ACTIVITY`), // set by user
			ServerURL:  os.Getenv(`GITHUB_SERVER_URL`),
		},
		gitHub: github.Client{
			APIURL: os.Getenv(`GITHUB_API_URL`),
		},
		slack: slack.Client{
			APIURL:   `https://slack.com/api/`,     // hardcoded
			APIToken: os.Getenv(`SLACK_API_TOKEN`), // set by user
			Channel:  os.Getenv(`SLACK_CHANNEL`),   // set by user
		},
	}
}

func main() {
	a := getEnvArgs()

	if err := a.validate(); err != nil {
		log.Fatalf("issue processing arguments %v", err)
	}

	// lazy cleanup githubRef
	strippedRef := a.gitHubAction.Ref[len(`refs/heads/`):]

	next, err := util.NewReleaseFromString(strippedRef)
	if err != nil {
		log.Fatalf("issue processing release: %v", err)
	}

	gitHubClient := a.gitHub

	prev, err := gitHubClient.GetPreviousRelease(a.gitHubAction.Repository, next)
	if err != nil {
		// exit
		log.Fatalf("issue with github: issue fetching previous release: %v", err)
	}

	slackClient := a.slack

	comment := fmt.Sprintf("%v - %s performed activity %q", next, a.gitHubAction.Actor, a.gitHubAction.Activity)

	err = slackClient.SendReleaseNotification(a.gitHubAction.ServerURL, a.gitHubAction.Repository, prev, next, comment)

	if err != nil {
		log.Fatalf("issue with slack: issue sending notification %v", err)
	}
}
