# Gonnachan

[![Go Report Card](https://goreportcard.com/badge/github.com/insomnyawolf/Gonnachan)](https://goreportcard.com/report/github.com/insomnyawolf/Gonnachan)

Konachan api written in go WORK IN PROGRESS

Working:
    PostSearch
        http://konachan.com/
        https://yande.re/
        https://gelbooru.com/"
        https://safebooru.org/"
        https://rule34.xxx/

#Usage
Imports
```go
import "github.com/insomnyawolf/Gonnachan"
```
Code
```go
Req := gonnachan.KonachanPostRequest{
		AfterID : 	int64,
		BeforeID: 	int64,
		Height: 	int,
		MaxResults: int,
		RandOrder: 	bool,
		Rating: 	gonnachan.Rating*,
		Tags: 		[]string{},
		TargetAPI: 	gonnachan.Server*,
		Width: 		int,
	}
	//String, Final Query
	Req.APIrequest()
	//Results Execute and process query.
	//Returns Struct with all the data needed
	res := Req.GetResults()
```