package gonnachan

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

//GetResults runs the query obtained at APIrequest and returns PostResult
func (c *PostRequest) GetResults() ([]PostResult, error) {
	URL := c.APIrequest()
	URL = strings.Replace(URL, " ", "%20", -1)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	thing, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	APIresult := string(thing)
	if APIresult == "[]" {
		return nil, errors.New("No results")
	}
	if APIresult == "" {
		return nil, errors.New("Empty Response")
	}

	//Single alocation
	var results []PostResult
	for x := 0; x < c.MaxResults; x++ {
		id := gjson.GetBytes(thing, fmt.Sprintf("%v.id", x))
		if !id.Exists() {
			break
		}
		r := PostResult{
			ID:       id.Int(),
			Tags:     gjson.GetBytes(thing, fmt.Sprintf("%v.tags", x)).String(),
			Author:   gjson.GetBytes(thing, fmt.Sprintf("%v.author", x)).String(),
			Source:   gjson.GetBytes(thing, fmt.Sprintf("%v.source", x)).String(),
			Score:    gjson.GetBytes(thing, fmt.Sprintf("%v.score", x)).Int(),
			FileSize: gjson.GetBytes(thing, fmt.Sprintf("%v.file_size", x)).Int(),
			Rating:   gjson.GetBytes(thing, fmt.Sprintf("%v.rating", x)).String(),
			Width:    gjson.GetBytes(thing, fmt.Sprintf("%v.width", x)).Int(),
			Height:   gjson.GetBytes(thing, fmt.Sprintf("%v.height", x)).Int(),
			FileURL:  gjson.GetBytes(thing, fmt.Sprintf("%v.file_url", x)).String(),
			Md5:      gjson.GetBytes(thing, fmt.Sprintf("%v.md5", x)).String(),
		}

		r.PostURL = fmt.Sprintf("%vpost/show/%v", c.TargetAPI, r.ID)

		//Server Specific Code
		//To be impoved
		if c.serverType == typeKonachan {
			//No specific code due to taking this as "default"
		} else if c.serverType == typeGelboru {
			//ImageUrl
			if c.TargetAPI != ServerGelbooru {
				image := gjson.GetBytes(thing, fmt.Sprintf("%v.image", x)).String()
				directory := gjson.GetBytes(thing, fmt.Sprintf("%v.directory", x)).String()
				r.FileURL = fmt.Sprintf("%s/images/%s/%s", c.TargetAPI, directory, image)
			}
			//Md5
			r.Md5 = gjson.GetBytes(thing, fmt.Sprintf("%v.hash", x)).String()
		} else if c.serverType == typeSankaku {
			//ToDo
		}

		//Fix url for konachan sites which doesn't start with http: on the file url
		if !strings.HasPrefix(r.FileURL, "https") && r.FileURL != "" {
			r.FileURL = "https:" + r.FileURL
		}
		results = append(results, r)
	}
	return results, nil
}

//RatingString Returns the human-readable sting for rating values
func (c *PostResult) RatingString() string {
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

//PostResult has useful data obtained from the API
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
