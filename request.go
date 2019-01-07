package gonnachan

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	serversGelboru  = []string{ServerGelbooru, ServerSafebooru, ServerRule34}
	serversKonachan = []string{ServerKonachan, ServerYandere}
)

const (
	//ServerKonachan APIendPoint
	ServerKonachan = "http://konachan.com/"
	//ServerYandere APIendPoint
	ServerYandere = "https://yande.re/"
	//ServerGelbooru APIendPoint
	ServerGelbooru = "https://gelbooru.com/"
	//ServerSafebooru APIendPoint
	ServerSafebooru = "https://safebooru.org/"
	//ServerRule34 APIendPoint
	ServerRule34 = "https://rule34.xxx/"

	//Common query to get post result in similar pages
	endpointGelboru  = "index.php?page=dapi&s=post&q=index&json=1&"
	endpointKonachan = "post.json?"
	//Enum Alternative
	typeUnsupported = -1
	typeKonachan    = 0
	typeGelboru     = 1

	//Danbooru APIendPoint
	//Danbooru = "https://danbooru.donmai.us/posts/1.json"
	//Im too lazy to fix this at the moment.

	//ToDo
	//chan.sankakucomplex.com
	//idol.sankakucomplex.com

	//e-shuushuu.net

	//anime-pictures.net

	//danbooru.donmai.us

	//www.zerochan.net
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

const (
	//ErrNoResults No results found
	ErrNoResults = 1
	//ErrEmpty No answer from server
	ErrEmpty = 2
)

var (
	//URL Base address for the requests
	URL string
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
		URL = ServerGelbooru
	} else {
		URL = c.TargetAPI
	}
	var endpoint string
	if c.serverType == typeKonachan {
		endpoint = endpointKonachan
	} else if c.serverType == typeGelboru {
		endpoint = endpointGelboru
	}
	uri := URL + endpoint + query
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
