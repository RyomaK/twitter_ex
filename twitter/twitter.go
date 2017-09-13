package twitter

import (
	"fmt"

	"github.com/RyomaK/treasure2017/mid/RyomaK/webapp/analysis"
	"github.com/dghubble/go-twitter/twitter"

	"github.com/RyomaK/treasure2017/mid/RyomaK/webapp/regexp"
	"github.com/dghubble/oauth1"
	"log"
)

//timelineに流すデータの構造体
type Twe struct {
	CreateAt string
	User     *twitter.User
	Text     string
}

//Tweのスライス
type Twes struct {
	Tweets []Twe
}

func TwitterClient(t Tokens) *twitter.Client {
	config := oauth1.NewConfig(t.ck, t.cs)
	token := oauth1.NewToken(t.At, t.As)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return twitter.NewClient(httpClient)
}

//timelineに流すtweetを返す
func Timeline(client *twitter.Client) []Twe {
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		fmt.Println(err)
		return []Twe{}
	}

	twes := &[]Twe{}
	for _, v := range tweets {
		*twes = append(*twes, Twe{v.CreatedAt, v.User, regexp.ChangeURL(v.Text)})

	}
	return *twes
}

//Tweetを返す
func SerachTweet(client *twitter.Client, str string) []twitter.Tweet {
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: str,
	})
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}
	tweets := search.Statuses
	return tweets
}

//userのtweetを返す
func UserTweet(client *twitter.Client, screenName string, num int) []Twe {
	tweets, resp, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count: num,
	})
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}

	twes := &[]Twe{}
	for _, v := range tweets {
		*twes = append(*twes, Twe{v.CreatedAt, v.User, regexp.ChangeURL(v.Text)})
	}
	log.Println(resp)
	return *twes

}

//Meのクライアントデータを返す
func GetClientData(client *twitter.Client) *twitter.User {
	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	})
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return user
}

func UpdateTweet(client *twitter.Client, forUser string, str string) {
	if forUser == "" {
		client.Statuses.Update(str+"=>\n"+analysis.GetRhyme(str), nil)
	} else {
		client.Statuses.Update("@"+forUser+" "+str+"=>\n"+analysis.GetRhyme(str), nil)
	}
}
