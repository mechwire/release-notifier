# Release Notifier

This project is a GitHub action that sends a message to the given Slack channel if one of the release events it listens for is triggered.

The top level messages should look like:

> [your-org/your-project](#) release! 🎉 [v1.2.3](#) ➡️ **1.3.x**

The replies should look like:

> **[v1.3.0](#)** - [your-github-user-name](#) performed activity 'activity'"

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

## Limitations

* Releases are only recognized if they follow the exact format `v<Major>.<Minor>.<Patch>`
* To avoid excessive queries, if the previous major or minor release in GitHub is more than 19 entries behind the triggering release, the most recent preceding release will be selected. (e.g., if you make release `v1.1.20`, it thinks the preceding release is `v1.1.19` instead of `v1.0.9`, because we never fetched the `v1.0.9` release.)
* Due to Slack API permissions for bots, if a message cannot be found in the last 100 messages of a channel, a new top-level message will be created.
  - Starring or pinning things can sometimes be done, but are not always supported by search, making it hard to do a specific search.
* Due to my laziness, if someone posts a message that looks identical to what the bot would want to post, it will respond to the message as if it was the post.

<!-- References -->
[Slack Webhooks]: https://api.slack.com/messaging/webhooks
[Variables]: https://docs.github.com/en/free-pro-team@latest/actions/learn-github-actions/essential-features-of-github-actions#using-variables-in-your-workflows
[Environment Variables]: https://docs.github.com/en/free-pro-team@latest/actions/reference/environment-variables#default-environment-variables
[Events]: https://docs.github.com/en/free-pro-team@latest/actions/reference/events-that-trigger-workflows#release
[Patti LaBelle]: https://youtu.be/ROIYcZGbfH0