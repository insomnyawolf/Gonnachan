package gonnachan

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	//Konachan APIendPoint
	Konachan = "http://konachan.com/post.json?"
	//Gelbooru APIendPoint
	Gelbooru = "https://gelbooru.com/index.php?page=dapi&s=post&q=index&json=1&"
)

var (
	//URL Base address for the requests
	URL string
)

const (
	//Rating

	//RatingSafe PG
	RatingSafe = "s"
	//RatingQuestionable +16
	RatingQuestionable = "q"
	//RatingExplicit +18
	RatingExplicit = "e"
)

//KonachanPostRequest store data prepare the api querry
type KonachanPostRequest struct {
	Tags       []string
	BeforeID   int64
	AfterID    int
	RandOrder  bool
	Rating     string
	Height     int
	Width      int
	MaxResults int
	TargetAPI  string
}

//Close EOL
func (c *KonachanPostRequest) Close() {
	c = nil
}

//APIrequest parse KonachanPostRequest to get the equivalent api query url
func (c *KonachanPostRequest) APIrequest() string {

	tags := strings.Join(c.Tags, "+")

	if tags != "" {
		tags += "+"
	}

	if c.BeforeID != 0 {
		tags += fmt.Sprintf("id:<%v+", c.BeforeID)
	}

	if c.AfterID != 0 {
		tags += fmt.Sprintf("id:>%v+", c.AfterID)
	}

	if c.RandOrder {
		tags += "order%3Arandom+"
	}

	if c.Rating != "" {
		tags += fmt.Sprintf("rating%%3A%v+", c.Rating)
	}

	if c.Height != 0 {
		tags += fmt.Sprintf("height%%3A%v+", c.Height)
	}

	if c.Width != 0 {
		tags += fmt.Sprintf("width%%3A%v+", c.Width)
	}

	if c.MaxResults == 0 {
		c.MaxResults = 1
	}

	query := fmt.Sprintf("limit=%v&tags=%v", strconv.Itoa(c.MaxResults), tags)
	if c.TargetAPI == "" {
		URL = Konachan
	} else {
		URL = c.TargetAPI
	}
	uri := URL + query
	return uri
}

//GetResults runs the query obtained at APIrequest and returns KonachanPostResult
func (c *KonachanPostRequest) GetResults() []KonachanPostResult {
	URL := c.APIrequest()
	res, err := http.Get(URL)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	thing, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	if string(thing) == "[]" {
		return nil
	}

	//Single alocation
	var results []KonachanPostResult
	for x := 0; x < c.MaxResults; x++ {
		r := KonachanPostResult{
			ID:       gjson.GetBytes(thing, fmt.Sprintf("%v.id", x)).Int(),
			Tags:     gjson.GetBytes(thing, fmt.Sprintf("%v.tags", x)).String(),
			Author:   gjson.GetBytes(thing, fmt.Sprintf("%v.author", x)).String(),
			Source:   gjson.GetBytes(thing, fmt.Sprintf("%v.source", x)).String(),
			Score:    gjson.GetBytes(thing, fmt.Sprintf("%v.score", x)).Int(),
			Md5:      gjson.GetBytes(thing, fmt.Sprintf("%v.md5", x)).String(),
			FileSize: gjson.GetBytes(thing, fmt.Sprintf("%v.file_size", x)).Int(),
			FileURL:  gjson.GetBytes(thing, fmt.Sprintf("%v.file_url", x)).String(),
			Rating:   gjson.GetBytes(thing, fmt.Sprintf("%v.rating", x)).String(),
			Width:    gjson.GetBytes(thing, fmt.Sprintf("%v.width", x)).Int(),
			Height:   gjson.GetBytes(thing, fmt.Sprintf("%v.height", x)).Int(),
		}
		//Fix url for konachan sites which doesn't start with http: on the file url
		if !strings.HasPrefix(r.FileURL, "http") && r.FileURL != "" {
			r.FileURL = "http:" + r.FileURL
		}
		results = append(results, r)
	}
	return results
}

//KonachanPostResult has useful data obtained from the API
type KonachanPostResult struct {
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
}

//Close EOL
func (c *KonachanPostResult) Close() {
	c = nil
}

//GetPostURL returns Konachan post url
func (c *KonachanPostResult) GetPostURL() string {
	return fmt.Sprintf("%v/post/show/%v", URL, c.ID)
}

//RatingString Returns the human-readable sting for rating values
func (c *KonachanPostResult) RatingString() string {
	switch c.Rating {
	case RatingSafe:
		return "Safe"
	case RatingQuestionable:
		return "Questionable"
	case RatingExplicit:
		return "Explicit"
	default:
		return ""
	}
}
