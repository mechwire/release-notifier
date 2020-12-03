package slack

type message struct {
	ThreadTS float64 `json:"thread_ts"`
}

func (c *Client) searchMessages(query string, options map[string]interface{}) ([]message, error) {

	return []message{}, nil

}

func (c *Client) chatPostMessage(channel string, text string, options map[string]interface{}) (message, error) {

	return message{}, nil

}
