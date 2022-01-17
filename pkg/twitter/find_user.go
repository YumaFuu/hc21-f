package twitter

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	UsersResponse struct {
		Data []User `json:"data"`
	}
	UserResponse struct {
		Data User `json:"data"`
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

func (t *Twitter) GetUserByID(id string) (User, error) {
	q := map[string]string{}

	r, err := t.call(fmt.Sprintf("2/users/%s", id), q)
	if err != nil {
		return User{}, err
	}
	fmt.Println(r)

	u := UserResponse{}
	err = json.Unmarshal([]byte(r), &u)

	if err != nil {
		return User{}, err
	}

	return u.Data, nil
}
