package github

import (
	"strings"

	"github.com/jncmaguire/labelle-release-notifier/internal/util"
)

type Action struct {
	ServerURL  string
	Actor      string
	Repository string
	Ref        string
	Event      string
	Activity   string
}

type Client struct {
	APIURL string
}

func (c *Client) GetPreviousRelease(repository string, next util.Release) (release util.Release, err error) {

	reps := strings.Split(repository, "/")
	owner := reps[0]
	project := reps[1]

	releases, err := c.getReleases(owner, project, 5, 1)

	if err != nil {
		return release, err
	}

	for i := range releases {
		if releases[i].Less(next) && release.Less(releases[i]) { // grab the most recent release before "next"
			release = releases[i]
		}
	}

	return release, err
}
