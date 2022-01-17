package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	FriendsResponse struct {
		IDs    []int         `json:"ids"`
		Error  string        `json:"error"`
		Errors []interface{} `json:"errors"`
	}
)

var (
	UnauthorizedError = errors.New("twitter api unauthorized error")
	RateLimitError    = errors.New("twitter api rate limit error")
)

func (t *Twitter) GetFriends(id string) ([]int, error) {
	q := map[string]string{
		"user_id": id,
	}
	r, err := t.call("1.1/friends/ids.json", q)
	if err != nil {
		return nil, err
	}

	fs := FriendsResponse{}
	err = json.Unmarshal([]byte(r), &fs)
	if err != nil {
		return nil, err
	}

	if fs.Error != "" {
		fmt.Println("Unauthorized")
		return nil, UnauthorizedError
	}

	if len(fs.Errors) != 0 {
		fmt.Println("RateLimit")
		return nil, RateLimitError
	}

	return fs.IDs, nil
}
