package main

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
	APIurl = "http://konachan.com/post.json?"
)

type KonachanPostRequest struct {
	tags       []string
	randOrder  bool
	rating     string
	height     int
	width      int
	maxResults int
}

func (c *KonachanPostRequest) APIrequest() string {

	tags := ""

	if c.randOrder {
		tags += "order%3Arandom+"
	}

	for _, t := range c.tags {
		tags += t + "+"
	}

	if c.rating != "" {
		tags += fmt.Sprintf("rating%%3A%v", c.rating)
	}

	if c.height != 0 {
		tags += fmt.Sprintf("height%%3A%v+", c.height)
	}

	if c.width != 0 {
		tags += fmt.Sprintf("width%%3A%v", c.width)
	}

	if c.maxResults == 0 {
		c.maxResults = 1
	}

	return fmt.Sprintf("%vlimit=%v&tags=%v", APIurl, strconv.Itoa(c.maxResults), tags)
}

type KonachanPostResult struct {
	id  string
	url string
}

//Search a pic in Image Booru sites api
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

	for x := 0; x < c.maxResults; x++ {
		id := gjson.GetBytes(thing, fmt.Sprintf("%v.id", x)).String()
		url := gjson.GetBytes(thing, fmt.Sprintf("%v.file_url", x)).String()
		//Fix url for konachan sites wich doesn't start with http: on the file url
		if !strings.HasPrefix(url, "http") && url != "" {
			url = "http:" + url
		}
		results = append(results, KonachanPostResult{id: id, url: url})
	}

	return results
}

func (c *KonachanPostResult) SearchID(id int) []KonachanPostResult {
	tag := []string{fmt.Sprintf("id:%v", id)}
	req := KonachanPostRequest{tags: tag}
	return req.GetResults()
}
