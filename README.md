# Release Notifier

This project is a GitHub action that sends a message to the given Slack channel if one of the release events it listens for is triggered.

This repository was partially inspired by [the advocate of frequent releases][Patti LaBelle].

## Getting Started


```yaml
on:
  release:
    types: [published, created, edited, deleted, prereleased, released]

jobs:
  release-notify-slack:    
    name: Send notification to Slack about release
    uses: jncmaguire/release-notifier@master
    with:
        SLACK_API_TOKEN: "my-api-token"
        SLACK_CHANNEL: "my-release-channel"
```


<!-- References -->
[Slack Webhooks]: https://api.slack.com/messaging/webhooks
[Variables]: https://docs.github.com/en/free-pro-team@latest/actions/learn-github-actions/essential-features-of-github-actions#using-variables-in-your-workflows
[Environment Variables]: https://docs.github.com/en/free-pro-team@latest/actions/reference/environment-variables#default-environment-variables
[Events]: https://docs.github.com/en/free-pro-team@latest/actions/reference/events-that-trigger-workflows#release
[Patti LaBelle]: https://youtu.be/ROIYcZGbfH0