run: 
	go run main.go

install:
	go get github.com/gin-gonic/gin
	go get github.com/dghubble/oauth1/
	go get github.com/ikawaha/kagome/...
	go get github.com/dghubble/go-twitter/twitter
	go get github.com/PuerkitoBio/goquery
