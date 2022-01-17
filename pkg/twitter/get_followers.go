package twitter

import (
	"encoding/json"
)

type (
	FollowersResponse struct {
		IDs []FollowerID `json:"ids"`
	}
	FollowerID int
)

func (t *Twitter) GetFollowers(id string) ([]FollowerID, error) {
	q := map[string]string{
		"user_id": id,
	}
	r, err := t.call("1.1/followers/ids.json", q)
	if err != nil {
		return nil, err
	}
	// fmt.Println(r)

	fs := FollowersResponse{}
	err = json.Unmarshal([]byte(r), &fs)

	if err != nil {
		return nil, err
	}

	return fs.IDs, nil
}
