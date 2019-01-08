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
	Req := gonnachan.GonachanPostRequest{}
```
PostRequest:
```go
type PostRequest struct {
	Tags       []string
	BeforeID   int64
	AfterID    int64
	RandOrder  bool
	Rating     string
	Height     int
	Width      int
	MaxResults int
	TargetAPI  string
	serverType int
	url        string
}
```
Methods:
```go
	//String, Final Query
	Req.APIrequest()
	//Results Execute and process query.
	//Returns Struct with all the data needed
	res := Req.GetResults()
```
PostResult:
```go
type PostResult struct {
	ID       int64  `json:"id"`
	Tags     string `json:"tags"`
	Author   string `json:"author"`
	Source   string `json:"source"`
	Score    int64  `json:"score"`
	Md5      string `json:"md5"`
	FileSize int64  `json:"file_size"`
	FileURL  string `json:"file_url"`
	Rating   string `json:"rating"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	PostURL  string
}
```