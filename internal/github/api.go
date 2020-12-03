package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/jncmaguire/labelle-release-notifier/internal/util"
)

func (c *Client) request(method string, path string, pathArgs map[string]interface{}, body interface{}) (*http.Response, error) {

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.APIURL+path, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	v := url.Values{}

	for key, value := range pathArgs {
		v.Add(key, fmt.Sprintf("%v", value))
	}

	req.URL.RawQuery = v.Encode()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(`Accept`, `application/vnd.github.v3+json`)

	return http.DefaultClient.Do(req)
}

func (c *Client) getReleases(owner string, repo string, perPage int, page int) ([]util.Release, error) {
	response, err := c.request(http.MethodGet, fmt.Sprintf("/repos/%s/%s/releases", owner, repo), map[string]interface{}{
		`per_page`: perPage,
		`page`:     page,
	}, nil)

	if err != nil {
		return []util.Release{}, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []util.Release{}, err
	}

	objects := make([]struct {
		TagName string `json:"tag_name"`
	}, perPage)

	if err = json.Unmarshal(data, &objects); err != nil {
		return []util.Release{}, err
	}

	releases := make([]util.Release, 0, perPage)

	i := 0
	for i < perPage {
		release, err := util.NewReleaseFromString(objects[i].TagName)

		if err == nil {
			releases = append(releases, release)
		}
	}

	return releases, err
}
