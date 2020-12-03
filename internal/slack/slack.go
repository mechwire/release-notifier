package slack

import (
	"errors"
	"fmt"

	"github.com/jncmaguire/labelle-release-notifier/internal/util"
)

type Config struct {
	Webhook   string
	Username  string
	ChannelID string
}

type Client struct {
	Webhook   string
	Username  string
	ChannelID string
}

func (c *Client) getNewestSignificantRelease(text string) (message, error) {

	messages, err := c.searchMessages(fmt.Sprintf("in:%s from:v%s %s", c.ChannelID, c.Username, text), nil)

	// then loop through to find parent

	// if it doesn't exist, create it
	// if it does exist, get the parent
	// if it has a thread_ts, it's threaded
	// if the trhread_ts is the same as the ts value, it's the parent

	// return the post info

	if err != nil {
		return message{}, err
	}

	if len(messages) == 0 {
		return message{}, nil
	}

	for _ = range messages {
		continue

	}

	return message{}, nil
}

func (c *Client) postSignificantRelease(text string) (message, error) {
	return c.chatPostMessage(c.ChannelID, text, nil)
}

func (c *Client) getCurrentReleaseNotification(repositoryServerURL string, repository string, prev util.Release, next util.Release) (message, error) {

	text := fmt.Sprintf("<%[1]s|%[2]s> release! :tada:  <%[1]s/releases/tag/%[3]v|%[3]v> :arrow_right: *v%[4]v.%[5]v.x*", repositoryServerURL+repository, repository, prev, next.Major, next.Minor)

	msg, err := c.getNewestSignificantRelease(text)

	if err != nil {
		return message{}, err
	}

	if (msg == message{}) {
		msg, err = c.postSignificantRelease(text)
	}

	return msg, err
}

func (c *Client) SendReleaseNotification(repositoryServerURL string, repository string, prev util.Release, next util.Release, text string) error {

	msg, err := c.getCurrentReleaseNotification(repositoryServerURL, repository, prev, next)

	if err != nil {
		return err
	}

	if (msg == message{}) {
		return errors.New("issue")
	}

	_, err = c.chatPostMessage(c.ChannelID, text, map[string]interface{}{
		"thread_ts": msg.ThreadTS,
	})

	return err
}
