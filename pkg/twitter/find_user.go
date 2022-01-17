package twitter

import (
	"encoding/json"
	"strings"
)

type (
	UsersResponse struct {
		Data []User `json:"data"`
	}
	User struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	}
)

func (t *Twitter) GetUserIDByUsernames(usernames []string) ([]User, error) {
	s := strings.Join(usernames, ",")

	q := map[string]string{
		"usernames": s,
	}
	r, err := t.call("2/users/by", q)
	if err != nil {
		return nil, err
	}

	us := UsersResponse{}
	err = json.Unmarshal([]byte(r), &us)

	if err != nil {
		return nil, err
	}

	return us.Data, nil
}
