package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jncmaguire/release-notifier/internal/util"
)

type message struct {
	Username string
	TS       float64
}

func (c *Client) request(method string, baseURL string, path string, pathArgs map[string]interface{}, body interface{}) ([]byte, error) {
	request, err := util.BuildRequest(http.MethodGet, c.APIURL, path, pathArgs, body)

	if err != nil {
		return nil, err
	}

	request.Header.Add(`Authorization`, fmt.Sprintf(` Bearer %s`, c.APIToken))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	object := struct {
		Error string
		Ok    bool
	}{}

	if err = json.Unmarshal(data, &object); err != nil {
		return nil, err
	}

	if !object.Ok {
		return nil, errors.New(object.Error)
	}

	return data, nil
}

func (c *Client) searchMessages(query string, options map[string]interface{}) ([]message, error) {

	data, err := c.request(http.MethodGet, c.APIURL, "/search.messages", map[string]interface{}{
		`query`: query,
	}, nil)
	if err != nil {
		return []message{}, err
	}

	object := struct {
		Messages struct {
			Matches []message
		}
	}{}

	if err = json.Unmarshal(data, &object); err != nil {
		return nil, err
	}

	return object.Messages.Matches, nil
}

func (c *Client) chatPostMessage(text string, options map[string]interface{}) (message, error) {

	data, err := c.request(http.MethodPost, c.APIURL, "/chat.postMessage", map[string]interface{}{
		`channel`: c.Channel,
		`text`:    text,
	}, nil)

	if err != nil {
		return message{}, err
	}

	object := struct {
		CreatedMessage message `json:"message"`
	}{}

	if err = json.Unmarshal(data, &object); err != nil {
		return message{}, err
	}

	return object.CreatedMessage, nil
}
