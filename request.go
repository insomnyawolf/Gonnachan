package gonnachan

import (
	"fmt"
	"strconv"
	"strings"
)

//KonachanPostRequest store data prepare the api querry
type KonachanPostRequest struct {
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

//APIrequest parse KonachanPostRequest to get the equivalent api query url
func (c *KonachanPostRequest) APIrequest() string {
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
	}
	uri := c.url + endpoint + query
	return uri
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

func (c *KonachanPostRequest) getServerKind() int {
	if contains(serversKonachan, c.TargetAPI) {
		return typeKonachan
	} else if contains(serversGelboru, c.TargetAPI) {
		return typeGelboru
	} else if contains(serversGelboru, c.TargetAPI) {
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
