package main

import (
	"log"

	"hc21f/twitter"

	"github.com/joho/godotenv"
)

var (
	// メンバー
	twitterAccounts = []string{
		// QuizKnock/クイズノック
		"QuizKnock",
		// 伊沢拓司
		"tax_i_",
		// ふくらP（QuizKnockクイズノック）
		"fukura_p",
		// こうちゃん
		"Miracle_Fusion",
		// 山本 祥彰
		"quiz_yamamoto",
		// 河村・拓哉
		"kawamura_domo",
		// 須貝 駿貴
		"Sugai_Shunki",
		// 鶴崎 修功
		"Tsurusaki_H",
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := twitter.Init()
	if err != nil {
		log.Fatal(err)
	}
	t := twitter.Twitter

	t.GetUserID(twitterAccounts)

	// defer resp.Body.Close()

	// ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray))
}
