package slack

import (
	"fmt"
	"time"

	"github.com/jncmaguire/release-notifier/internal/util"
)

type Client struct {
	APIURL    string
	APIToken  string
	ChannelID string
}

func (c *Client) getNewestSignificantRelease(text string) (message, error) {

	messages, err := c.conversationsHistory(nil) // we don't want to exhaust the API; if it's not in the first 100 results, w/e

	if err != nil {
		return message{}, err
	}

	for i := range messages {
		if messages[i].Text == text { // edgecase: someone else posts the exact same message
			return messages[i], nil
		}
	}

	return message{}, nil
}

func (c *Client) postSignificantRelease(text string) (message, error) {
	return c.chatPostMessage(text, nil)
}

func (c *Client) getCurrentParentReleaseNotification(repositoryServerURL string, repository string, prev util.Release, next util.Release) (message, error) {

	text := fmt.Sprintf("<%[1]s|%[2]s> release! :tada:  <%[1]s/releases/tag/%[3]v|%[3]v> :arrow_right: *v%[4]v.%[5]v.x*", repositoryServerURL+"/"+repository, repository, prev, next.Major, next.Minor)

	msg, err := c.getNewestSignificantRelease(text)

	if err != nil {
		return message{}, err
	}

	if (msg == message{}) {

		if prev.UpgradeType(next) == util.Major {

		}
		msg, err = c.postSignificantRelease(text)
	}

	return msg, err
}

func (c *Client) SendReleaseNotification(repositoryServerURL string, repository string, prev util.Release, next util.Release, text string) error {

	currentParent, err := c.getCurrentParentReleaseNotification(repositoryServerURL, repository, prev, next)

	if err != nil {
		return err
	}

	time.Sleep(1 * time.Millisecond)

	_, err = c.chatPostMessage(text, map[string]interface{}{
		"thread_ts":    currentParent.TS,
		`unfurl_media`: false,
		`unfurl_links`: false,
	})

	return err
}
