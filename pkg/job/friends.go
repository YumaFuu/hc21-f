package job

import (
	"encoding/csv"
	"fmt"
	"hc21f/pkg/twitter"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	FriendsCsvHeader   = []string{"ID"}
	FriendsCsvFilePath = "./data/friends.csv"
)

func (job *Job) SearchFriends() error {
	f, err := os.Open(FollowerCsvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if err != nil {
		return err
	}

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(row) == 0 {
			continue
		}

		id := row[0]
		fmt.Println("[START]")

		err = job.getFriends(id)
		if err != nil {
			log.Printf("[ ERROR ]: get %s's friends failed: %v", id, err)
			continue
		}

	}

	return nil
}

func (job *Job) getFriends(uid string) error {
	b, err := ioutil.ReadFile(FriendsCsvFilePath)
	if err != nil {
		return err
	}
	hasIDs := strings.Contains(string(b), uid)

	if hasIDs {
		fmt.Println("next")
		return nil
	}

	// time.Sleep(time.Second * 3)
	f, err := os.OpenFile(FriendsCsvFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	ids, err := job.twitter.GetFriends(uid)

	if err != nil {
		if err == twitter.UnauthorizedError {
			str := fmt.Sprintf("%s,UNAUTHORIZED\n", uid)
			_, err := f.Write([]byte(str))

			if err != nil {
				fmt.Println("eeeee", err)
				return err
			}
		}

		if err == twitter.RateLimitError {
			os.Exit(1)
		}
		return err
	}

	for _, id := range ids {
		fmt.Println(id)
		str := fmt.Sprintf("%s,%d\n", uid, id)
		_, err := f.Write([]byte(str))

		if err != nil {
			fmt.Println("dddd", err)
			return err
		}

	}

	return nil
}
