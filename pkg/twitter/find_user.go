package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type (
	UsersResponse struct {
		Data []User `json:"data"`
	}
	UserResponse struct {
		Data  User   `json:"data"`
		Title string `json:"title"`
	}
	User struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Username    string `json:"username"`
		Description string `json:"description"`
	}
)

var (
	TooManyRequestError = errors.New("too many request")
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
	q := map[string]string{
		"user.fields": "description",
	}

	time.Sleep(time.Second * 1)

	r, err := t.call(fmt.Sprintf("2/users/%s", id), q)
	if err != nil {
		return User{}, err
	}
	fmt.Println(r)

	u := UserResponse{}
	err = json.Unmarshal([]byte(r), &u)

	if u.Title != "" {
		return User{}, TooManyRequestError
	}

	if err != nil {
		return User{}, err
	}

	return u.Data, nil
}
