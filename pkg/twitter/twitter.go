package twitter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

const (
	APIBaseURL = "https://api.twitter.com"
)

type Twitter struct {
	token string
}

var twitter Twitter

func Get() Twitter {
	return twitter
}

func Init() error {
	t := os.Getenv("TWITTER_BEARER_TOKEN")
	if t == "" {
		return errors.New("twitter token is not set to ENV")
	}

	twitter.token = t

	return nil
}

func (t *Twitter) call(endpoint string, query map[string]string) (string, error) {
	u, err := url.Parse(APIBaseURL)

	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, endpoint)
	q := u.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", t.token),
	)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
