package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelease_UpgradeType(t *testing.T) {
	scenarios := []struct {
		name    string
		a       Release
		b       Release
		upgrade upgradeType
	}{

		{
			name:    "major",
			a:       Release{},
			b:       Release{Major: 1, Minor: 1, Patch: 1},
			upgrade: Major,
		},
		{
			name:    "minor",
			a:       Release{},
			b:       Release{Major: 0, Minor: 1, Patch: 1},
			upgrade: Minor,
		},
		{
			name:    "patch",
			a:       Release{},
			b:       Release{Major: 0, Minor: 0, Patch: 1},
			upgrade: Patch,
		},
		{
			name:    "less",
			a:       Release{Major: 3, Minor: 0, Patch: 0},
			b:       Release{Major: 1, Minor: 1, Patch: 1},
			upgrade: NoUpgrade,
		},
		{
			name:    "equal",
			a:       Release{},
			b:       Release{},
			upgrade: NoUpgrade,
		},
	}

	for i := range scenarios {
		t.Run(scenarios[i].name, func(t *testing.T) {
			assert.Exactly(t, scenarios[i].upgrade, scenarios[i].a.UpgradeType(scenarios[i].b))
		})
	}
}

func TestRelease_Less(t *testing.T) {
	scenarios := []struct {
		name   string
		a      Release
		b      Release
		result bool
	}{

		{
			name:   "major",
			a:      Release{},
			b:      Release{Major: 1, Minor: 1, Patch: 1},
			result: true,
		},
		{
			name:   "minor",
			a:      Release{},
			b:      Release{Major: 0, Minor: 1, Patch: 1},
			result: true,
		},
		{
			name:   "patch",
			a:      Release{},
			b:      Release{Major: 0, Minor: 0, Patch: 1},
			result: true,
		},
		{
			name:   "less",
			a:      Release{Major: 3, Minor: 0, Patch: 0},
			b:      Release{Major: 1, Minor: 1, Patch: 1},
			result: false,
		},
	}

	for i := range scenarios {
		t.Run(scenarios[i].name, func(t *testing.T) {
			assert.Exactly(t, scenarios[i].result, scenarios[i].a.Less(scenarios[i].b))
		})
	}
}
