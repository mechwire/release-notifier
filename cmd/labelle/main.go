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
	GitHubAction github.Action
	GitHubConfig github.Config
	SlackConfig  slack.Config
}

func (a args) validate() (err error) {
	for _, val := range []string{
		a.GitHubAction.Actor,
		a.GitHubAction.Repository,
		a.GitHubAction.Ref,
		a.GitHubAction.Event,
		a.GitHubAction.ServerURL,
		a.GitHubConfig.APIURL,
		a.GitHubAction.Activity,
		a.SlackConfig.Webhook,
		a.SlackConfig.ChannelID,
		a.SlackConfig.Username,
	} {
		if val == "" {
			err = errors.New("Value should not be empty")
		}
	}

	return err
}

func getEnvArgs() args {
	return args{
		GitHubAction: github.Action{
			Actor:      os.Getenv(`GITHUB_ACTOR`),
			Repository: os.Getenv(`GITHUB_REPOSITORY`),
			Ref:        os.Getenv(`GITHUB_REF`),
			Event:      os.Getenv(`GITHUB_EVENT`),
			Activity:   os.Getenv(`GITHUB_EVENT_ACTIVITY`), // set by user
			ServerURL:  os.Getenv(`GITHUB_SERVER_URL`),
		},
		GitHubConfig: github.Config{
			APIURL: os.Getenv(`GITHUB_API_URL`),
		},
		SlackConfig: slack.Config{
			Webhook:   os.Getenv(`SLACK_WEBHOOK`),    // set by user
			ChannelID: os.Getenv(`SLACK_CHANNEL_ID`), // set by user
			Username:  os.Getenv(`SLACK_USERNAME`),   // set by user
		},
	}
}

func main() {
	a := getEnvArgs()

	if err := a.validate(); err != nil {
		log.Fatalf("issue processing arguments %v", err)
	}

	// lazy cleanup githubRef
	strippedRef := a.GitHubAction.Ref[len(`refs/heads/`):]

	next, err := util.NewReleaseFromString(strippedRef)
	if err != nil {
		log.Fatalf("issue processing release: %v", err)
	}

	gh := github.Client{}
	prev, err := gh.GetPreviousRelease(a.GitHubAction.Repository, next)
	if err != nil {
		// exit
		log.Fatalf("issue with github: issue fetching previous release: %v", err)
	}

	comment := fmt.Sprintf("%v - %s performed activity %q", next, a.GitHubAction.Actor, a.GitHubAction.Activity)

	s := slack.Client{}

	err = s.SendReleaseNotification(a.GitHubAction.ServerURL, a.GitHubAction.Repository, prev, next, comment)

	if err != nil {
		log.Fatalf("issue with slack: issue sending notification %v", err)
	}
}
