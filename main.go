package main

import (
	"net/http"

	"html/template"
	"log"

	"github.com/RyomaK/twitter_ex/twitter"

	"github.com/gin-gonic/gin"
)

func main() {
	//初期設定　client,tokenをセット
	tokens := twitter.Token()
	client := twitter.TwitterClient(*tokens)

	r := gin.Default()

	r.LoadHTMLGlob("view/*")
	//ルーティング
	r.GET("/", func(c *gin.Context) {
		if tokens.At != "" {
			client = twitter.TwitterClient(*tokens)
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"Me": twitter.GetClientData(client),
			})
		} else {
			c.HTML(http.StatusOK, "top.tmpl", gin.H{})
		}
	})

	r.GET("/login", func(c *gin.Context) {
		if tokens.At != "" {
			client = twitter.TwitterClient(*tokens)
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"Me": twitter.GetClientData(client),
			})
		} else {
			type Data struct {
				URL          string
				RequestToken string
			}
			data := &Data{}
			data.RequestToken, data.URL = twitter.GetRequestToken(twitter.GetConfig())
			c.HTML(http.StatusOK, "pin.tmpl", gin.H{
				"data": data,
			})
		}
	})

	r.POST("/login", func(c *gin.Context) {
		c.Request.ParseForm()
		pincode := c.PostForm("pin")
		rt := c.PostForm("token")

		accessToken, err := twitter.ReceivePIN(twitter.GetConfig(), rt, pincode)
		if err != nil {
			log.Fatalf("Access Token Phase: %s", err.Error())
		}
		tokens.At = accessToken.Token
		tokens.As = accessToken.TokenSecret

		client = twitter.TwitterClient(*tokens)

		if err != nil {
			log.Fatal("%v", err.Error())
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Me": twitter.GetClientData(client),
		})
	})

	r.GET("/ranking", func(c *gin.Context) {
		if tokens.At != "" {
			people := twitter.GetPeople(client)
			c.HTML(http.StatusOK, "ranking.tmpl", gin.H{
				"add":    func(a, b int) int { return a + b },
				"people": people,
				"Me":     twitter.GetClientData(client),
			})
		} else {
			c.HTML(http.StatusOK, "top.tmpl", gin.H{})
		}
	})

	r.GET("/home", func(c *gin.Context) {
		if tokens.At != "" {
			client = twitter.TwitterClient(*tokens)
			tweets := twitter.Timeline(client)

			c.HTML(http.StatusOK, "home.tmpl", gin.H{
				"safe":   func(text string) template.HTML { return template.HTML(text) },
				"Me":     twitter.GetClientData(client),
				"tweets": tweets,
			})
		} else {
			c.HTML(http.StatusOK, "top.tmpl", gin.H{})
		}
	})

	r.GET("/mypage", func(c *gin.Context) {
		if tokens.At != "" {
			client = twitter.TwitterClient(*tokens)
			me := twitter.GetClientData(client)
			tweets := twitter.UserTweet(client, me.ScreenName, 30)

			c.HTML(http.StatusOK, "mypage.tmpl", gin.H{
				"safe":   func(text string) template.HTML { return template.HTML(text) },
				"Me":     me,
				"tweets": tweets,
			})
		} else {
			c.HTML(http.StatusOK, "top.tmpl", gin.H{})
		}
	})

	r.GET("/index", func(c *gin.Context) {
		if tokens.At != "" {
			client = twitter.TwitterClient(*tokens)

			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"Me": twitter.GetClientData(client),
			})
		} else {
			c.HTML(http.StatusOK, "top.tmpl", gin.H{})
		}
	})

	r.POST("/tweet", func(c *gin.Context) {
		client = twitter.TwitterClient(*tokens)

		twitter.UpdateTweet(client, c.PostForm("for"), c.PostForm("writing"))
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Me": twitter.GetClientData(client),
		})
	})

	r.Run(":8080")
}
