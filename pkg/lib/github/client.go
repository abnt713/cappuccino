package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang.org/x/oauth2"
)

const githubURL = "https://api.github.com"

// Client interacts with the github API.
type Client struct {
	httpClient          *http.Client
	defaultPollInterval int
}

// NewClient creates a new client.
func NewClient(ctx context.Context, token string, defaultPollInterval int) Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return Client{
		httpClient:          tc,
		defaultPollInterval: defaultPollInterval,
	}
}

// GetNotifications retrieves notifications from github.
func (c Client) GetNotifications(lastModified string, showAll bool) (NotificationsResult, error) {
	const endpoint = "/notifications"
	result := NotificationsResult{PollInterval: c.defaultPollInterval}
	url := githubURL + endpoint + fmt.Sprintf("?all=%t", showAll)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	if lastModified != "" {
		req.Header.Add("If-Modified-Since", lastModified)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return result, err
	}

	if res.StatusCode == http.StatusNotModified {
		result.NotModified = true
		return result, nil
	}

	if res.StatusCode != http.StatusOK {
		return result, fmt.Errorf("non ok status received: %d", res.StatusCode)
	}

	pollInterval, err := strconv.Atoi(res.Header.Get("X-Poll-Interval"))
	if err != nil {
		return result, err
	}

	result.PollInterval = pollInterval
	result.LastModified = res.Header.Get("Last-Modified")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result.Notifications)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Notification is a github notification.
type Notification struct {
	Reason string `json:"reason"`
	Unread bool   `json:"unread"`
}

// IsUnread tells if a notification has been already read.
func (n Notification) IsUnread() bool {
	return n.Unread
}

// GetReason returns the notification reason.
func (n Notification) GetReason() string {
	return n.Reason
}

// NotificationsResult returns the results of a notification request.
type NotificationsResult struct {
	PollInterval  int
	LastModified  string
	NotModified   bool
	Notifications []Notification
}
