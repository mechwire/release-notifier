# Release Notifier

This project is a GitHub action that sends a message to the given Slack channel if one of the release events it listens for is triggered.

This repository is named after [the advocate of frequent releases][Patti LaBelle].\*

\* But not affiliated. Please don't sue me.

## Getting Started


```yaml
on:
  release:
    types: [published, created, edited, deleted, prereleased, released]

jobs:
  release-notify-slack:    
    name: Send notification to Slack about release
    runs-on: ubuntu-latest
    steps:
      - uses: actions/setup-go@v2
          with:
            go-version: '1.15.0'
       - run: go run cmd/notifier
         env:
           GITHUB_EVENT_ACTIVITY="${{ github.event.action }}"
           SLACK_API_URL: ""
           SLACK_API_TOKEN: ""
           SLACK_WEBHOOK: ""
           SLACK_CHANNEL: "my-release-channel"
           SLACK_USERNAME: "my-bot-name"
```


<!-- References -->
[Variables]: https://docs.github.com/en/free-pro-team@latest/actions/learn-github-actions/essential-features-of-github-actions#using-variables-in-your-workflows
[Environment Variables]: https://docs.github.com/en/free-pro-team@latest/actions/reference/environment-variables#default-environment-variables
[Events]: https://docs.github.com/en/free-pro-team@latest/actions/reference/events-that-trigger-workflows#release
[Patti LaBelle]: https://youtu.be/ROIYcZGbfH0