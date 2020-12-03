# Release Notifier

This project is a GitHub action that sends a message to the given Slack channel if one of the release events it listens for is triggered.

The top level messages should look like:

> [your-org/your-project](#) release! 🎉 [v1.2.3](https://github.com/your-org/your-project/releases/tag/v1.2.3) ➡️ **1.3.x**

The replies should look like:

> **v1.3.0** - **your-github-user-name** performed activity 'created'"

This repository was partially inspired by [the advocate of frequent releases][Patti LaBelle].

## Getting Started


```yaml
name: Release Notifier

on:
  release:
    types: [published, created, edited, deleted, prereleased, released]

jobs:
  release-notify-slack:
    name: Send notification to Slack about release
    runs-on: ubuntu-20.04
    steps:
     -  uses: jncmaguire/release-notifier@main
        run: notifier
        with:
          GITHUB_API_TOKEN: ${{ secrets.GITHUB_API_TOKEN }}
          SLACK_API_TOKEN: ${{ secrets.SLACK_API_TOKEN }}
          SLACK_CHANNEL_ID: "C1234567890"
```


<!-- References -->
[Slack Webhooks]: https://api.slack.com/messaging/webhooks
[Variables]: https://docs.github.com/en/free-pro-team@latest/actions/learn-github-actions/essential-features-of-github-actions#using-variables-in-your-workflows
[Environment Variables]: https://docs.github.com/en/free-pro-team@latest/actions/reference/environment-variables#default-environment-variables
[Events]: https://docs.github.com/en/free-pro-team@latest/actions/reference/events-that-trigger-workflows#release
[Patti LaBelle]: https://youtu.be/ROIYcZGbfH0