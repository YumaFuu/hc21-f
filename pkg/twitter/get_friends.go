package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	FriendsResponse struct {
		IDs    []int   `json:"ids"`
		Error  string  `json:"error"`
		Errors []Error `json:"errors"`
	}
)

var (
	UnauthorizedError = errors.New("twitter api unauthorized error")
	RateLimitError    = errors.New("twitter api rate limit error")
)

func (t *Twitter) GetFriends(id string) ([]int, error) {
	fmt.Println("call twitter api")
	q := map[string]string{
		"user_id": id,
	}
	r, err := t.call("1.1/friends/ids.json", q)
	if err != nil {
		return nil, err
	}
	// fmt.Println(r)

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
		if fs.Errors[0].Code == 88 {
			fmt.Println("RateLimit")
			return nil, RateLimitError
		} else {
			return nil, nil
		}
	}

	return fs.IDs, nil
}
