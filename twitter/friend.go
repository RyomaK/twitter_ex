package twitter

import (
	"fmt"
	"log"
	"sort"

	"github.com/dghubble/go-twitter/twitter"
)

type Person struct {
	ScreenName string
	Name       string
	Followers  int
	Me         bool
	Image      string
}

type People []Person

//200*15 = 3000フォローまで取得
func GetFriends(client *twitter.Client) []twitter.User {
	users := []twitter.User{}
	var cursor int64 = 0
	for i := 0; i < 15; i++ {
		friends, resp, err := client.Friends.List(&twitter.FriendListParams{
			UserID:     GetClientData(client).ID,
			SkipStatus: twitter.Bool(false),
			Cursor:     cursor,
			Count:      200,
		})
		cursor = friends.NextCursor
		log.Println("resp in getfriends :", resp)

		if err != nil {
			fmt.Errorf("birthday.go! %v", err)
		}

		users = append(users, friends.Users...)

		if len(friends.Users) < 200 {
			break
		}
	}

	return users
}

func GetPeople(client *twitter.Client) People {
	users := GetFriends(client)
	me := GetClientData(client)
	people := People{}
	people = append(people, Person{me.ScreenName, me.Name, me.FollowersCount, true, me.ProfileImageURL})
	for _, v := range users {
		people = append(people, Person{v.ScreenName, v.Name, v.FollowersCount, false, v.ProfileImageURL})
	}
	return SortLine(people)
}

/*sort*/

func (p People) Len() int           { return len(p) }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p People) Less(i, j int) bool { return p[i].Followers < p[j].Followers }

func SortLine(p People) People {
	sort.Sort(sort.Reverse(p))
	return p
}
