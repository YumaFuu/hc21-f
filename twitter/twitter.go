package twitter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

const (
	APIBaseURL = "https://api.twitter.com/2"
)

type twitter struct {
	token string
}

var Twitter twitter

func Init() error {
	t := os.Getenv("TWITTER_BEARER_TOKEN")
	if t == "" {
		log.Fatal("twitter token is not set")
	}

	Twitter.token = t

	return nil
}

func (t *twitter) call(endpoint string, query map[string]string) error {
	u, err := url.Parse(APIBaseURL)

	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, endpoint)
	q := u.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func (t *twitter) GetUserID(usernames []string) error {
	s := strings.Join(usernames, ",")

	q := map[string]string{
		"usernames": s,
	}
	return t.call("users/by", q)
}
