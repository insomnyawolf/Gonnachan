package gonnachan

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	//APIurl Base address for api requests
	APIurl = "http://konachan.com/post.json?"

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
	RandOrder  bool
	Rating     string
	Height     int
	Width      int
	MaxResults int
}

//APIrequest parse KonachanPostRequest to get the equivalent api query url
func (c *KonachanPostRequest) APIrequest() string {

	tags := ""

	if c.RandOrder {
		tags += "order%3Arandom+"
	}

	for _, t := range c.Tags {
		tags += t + "+"
	}

	if c.Rating != "" {
		tags += fmt.Sprintf("rating%%3A%v", c.Rating)
	}

	if c.Height != 0 {
		tags += fmt.Sprintf("height%%3A%v+", c.Height)
	}

	if c.Width != 0 {
		tags += fmt.Sprintf("width%%3A%v", c.Width)
	}

	if c.MaxResults == 0 {
		c.MaxResults = 1
	}

	return fmt.Sprintf("%vlimit=%v&tags=%v", APIurl, strconv.Itoa(c.MaxResults), tags)
}

//GetResults runs the query obtained at APIrequest and returns KonachanPostResult
func (c *KonachanPostRequest) GetResults() []KonachanPostResult {
	URL := c.APIrequest()
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	thing, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	if bytes.Equal(thing, []byte("[]")) {
		return nil
	}

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

//KonachanPostResult has usefull data obtained from the API
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
