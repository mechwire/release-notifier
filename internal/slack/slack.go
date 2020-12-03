package slack

import (
	"fmt"

	"github.com/jncmaguire/release-notifier/internal/util"
)

type Client struct {
	APIURL   string
	APIToken string
	Channel  string
}

func (c *Client) getNewestSignificantRelease(text string) (message, error) {

	messages, err := c.searchMessages(fmt.Sprintf("from:@me in:#%s %q", c.Channel, text), map[string]interface{}{
		`count`: 5,
	})

	if err != nil {
		return message{}, err
	}

	if len(messages) == 0 {
		return message{}, nil
	}

	return messages[0], nil // this is probably sufficient?
}

func (c *Client) postSignificantRelease(text string) (message, error) {
	return c.chatPostMessage(text, nil)
}

func (c *Client) getCurrentParentReleaseNotification(repositoryServerURL string, repository string, prev util.Release, next util.Release) (message, error) {

	text := fmt.Sprintf("<%[1]s|%[2]s> release! :tada:  <%[1]s/releases/tag/%[3]v|%[3]v> :arrow_right: *v%[4]v.%[5]v.x*", repositoryServerURL+repository, repository, prev, next.Major, next.Minor)

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

	_, err = c.chatPostMessage(text, map[string]interface{}{
		"thread_ts": currentParent.TS,
	})

	return err
}
