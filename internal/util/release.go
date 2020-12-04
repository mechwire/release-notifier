package util

import (
	"fmt"
	"regexp"
	"strconv"
)

type upgradeType int

const (
	NoUpgrade upgradeType = iota
	Patch
	Minor
	Major
)

type Release struct {
	Major int
	Minor int
	Patch int
}

func (a Release) Less(b Release) bool {
	return a.UpgradeType(b) != NoUpgrade
}

func (r Release) UpgradeType(upgrade Release) upgradeType {
	if r.Major < upgrade.Major {
		return Major
	} else if r.Major > upgrade.Major {
		return NoUpgrade
	}

	if r.Minor < upgrade.Minor {
		return Minor
	} else if r.Minor > upgrade.Minor {
		return NoUpgrade
	}

	if r.Patch < upgrade.Patch {
		return Patch
	}

	return NoUpgrade
}

func (r Release) String() string {
	return fmt.Sprintf("v%v.%v.%v", r.Major, r.Minor, r.Patch)
}

func NewReleaseFromString(s string) (Release, error) {
	pattern, err := regexp.Compile(`^v(?P<major>\d+).(?P<minor>\d+).(?P<patch>\d+)$`)
	if err != nil {
		return Release{}, err
	}

	r := Release{}

	matches := pattern.FindAllStringSubmatch(s, -1)

	for i, name := range pattern.SubexpNames() {
		if i != 0 && name != "" {
			v, err := strconv.Atoi(matches[0][i])
			if err != nil {
				return Release{}, err
			}

			switch name {
			case `major`:
				r.Major = v
			case `minor`:
				r.Minor = v
			case `patch`:
				r.Patch = v
			}
		}
	}

	return r, nil
}
