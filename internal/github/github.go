package github

import (
	"log"
	"strings"

	"github.com/jncmaguire/release-notifier/internal/util"
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
	APIURL   string
	APIToken string
}

func (c *Client) GetPreviousNonPatchRelease(repository string, next util.Release) (release util.Release, err error) {

	reps := strings.Split(repository, "/")
	owner := reps[0]
	project := reps[1]

	releases, err := c.getReleases(owner, project, 20, 1)

	if err != nil {
		return release, err
	}

	previousMajorOrMinor := util.Release{}
	previousPatch := util.Release{}
	for i := range releases {
		switch ut := releases[i].UpgradeType(next); ut {
		case util.Major, util.Minor:
			if previousMajorOrMinor.Less(releases[i]) {
				previousMajorOrMinor = releases[i]
			}
		case util.Patch:
			if previousPatch.Less(releases[i]) {
				previousPatch = releases[i]
			}
		}
	}

	log.Println(releases, previousMajorOrMinor, previousPatch)

	if (previousMajorOrMinor != util.Release{}) {
		release = previousMajorOrMinor
	} else {
		release = previousPatch // even if this is 0.0.0, this is fine
	}

	return release, err
}
