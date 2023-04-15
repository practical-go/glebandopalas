package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Cat_Fact struct {
	Title string
	Text  string `json:"text"`
}

type Space_Fact struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		URL         string    `json:"url"`
		ImageURL    string    `json:"image_url"`
		NewsSite    string    `json:"news_site"`
		Summary     string    `json:"summary"`
		PublishedAt time.Time `json:"published_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Featured    bool      `json:"featured"`
		Launches    []struct {
			LaunchID string `json:"launch_id"`
			Provider string `json:"provider"`
		} `json:"launches"`
		Events []interface{} `json:"events"`
	} `json:"results"`
}

func parseCatFacts() []Cat_Fact {
	var facts []Cat_Fact
	response, err := http.Get("https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=10")
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(response.Body).Decode(&facts)
	for i := 0; i < len(facts); i++ {
		facts[i].Title = "Cat Fact 🐱"
	}
	return facts
}

func parseSpaceFacts() []Space_Fact {
	var facts []Space_Fact
	response, err := http.Get("https://api.spaceflightnewsapi.net/v4/articles/")
	if err != nil {
		log.Fatal(err)
	}
	error := json.NewDecoder(response.Body).Decode(&facts)
	if error != nil {
		log.Fatal(error)
	}
	for i := 0; i < len(facts); i++ {
		fmt.Print(facts[i])
	}

	return facts
}

// Get Cat Facts
func HandleNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "		")
	enc.Encode(parseSpaceFacts())
	enc.Encode(parseCatFacts())
}
