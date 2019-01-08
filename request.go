package gonnachan

import (
	"fmt"
	"strconv"
	"strings"
)

//PostRequest store data prepare the api querry
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

//APIrequest parse PostRequest to get the equivalent api query url
func (c *PostRequest) APIrequest() string {
	c.serverType = c.getServerKind()
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
		c.url = ServerGelbooru
	} else {
		c.url = c.TargetAPI
	}
	var endpoint string
	if c.serverType == typeKonachan {
		endpoint = endpointKonachan
	} else if c.serverType == typeGelboru {
		endpoint = endpointGelboru
	} else if c.serverType == typeSankaku {
		endpoint = endpointSankaku
	}
	uri := c.url + endpoint + query
	return uri
}

func (c *PostRequest) getServerKind() int {
	if contains(serversKonachan, c.TargetAPI) {
		return typeKonachan
	} else if contains(serversGelboru, c.TargetAPI) {
		return typeGelboru
	} else if contains(serversSankaku, c.TargetAPI) {
		return typeSankaku
	}
	return typeUnsupported
}

func contains(strList []string, str string) bool {
	for _, a := range strList {
		if a == str {
			return true
		}
	}
	return false
}
