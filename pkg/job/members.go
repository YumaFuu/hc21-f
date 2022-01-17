package job

import (
	"fmt"
	"hc21f/pkg/csv"
)

var (
	Accounts = []string{
		"QuizKnock",
		// "tax_i_",
		// "fukura_p",
		// "Miracle_Fusion",
		// "quiz_yamamoto",
		// "kawamura_domo",
		// "Sugai_Shunki",
		// "Tsurusaki_H",
	}
	MemberCsvHeader   = []string{"ID", "Name", "Username"}
	MemberCsvFilePath = "./data/member.csv"
)

func (job *Job) FillMemberID() error {
	users, err := job.twitter.GetUserIDByUsernames(Accounts)
	if err != nil {
		return err
	}

	var userRows [][]string
	userRows = append(userRows, MemberCsvHeader)
	for _, u := range users {
		userRows = append(userRows, []string{u.ID, u.Name, u.Username})
	}

	err = csv.Add(MemberCsvFilePath, userRows)
	if err != nil {
		return err
	}

	fmt.Println("members successed")
	return nil
}
