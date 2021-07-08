import (
    "encoding/json"
    "fmt"
    "os"
)

type AutoGenerated struct {
	Data struct {
		Axies struct {
			Total   int `json:"total"`
			Results []struct {
				ID         string `json:"id"`
				Image      string `json:"image"`
				Class      string `json:"class"`
				Name       string `json:"name"`
				Genes      string `json:"genes"`
				Owner      string `json:"owner"`
				Stage      int    `json:"stage"`
				Title      string `json:"title"`
				BreedCount int    `json:"breedCount"`
				Level      int    `json:"level"`
				Parts      []struct {
					ID           string        `json:"id"`
					Name         string        `json:"name"`
					Class        string        `json:"class"`
					Type         string        `json:"type"`
					SpecialGenes interface{}   `json:"specialGenes"`
					Stage        int           `json:"stage"`
					Abilities    []interface{} `json:"abilities"`
					Typename     string        `json:"__typename"`
				} `json:"parts"`
				Stats struct {
					Hp       int    `json:"hp"`
					Speed    int    `json:"speed"`
					Skill    int    `json:"skill"`
					Morale   int    `json:"morale"`
					Typename string `json:"__typename"`
				} `json:"stats"`
				Auction struct {
					StartingPrice     string `json:"startingPrice"`
					EndingPrice       string `json:"endingPrice"`
					StartingTimestamp string `json:"startingTimestamp"`
					EndingTimestamp   string `json:"endingTimestamp"`
					Duration          string `json:"duration"`
					TimeLeft          string `json:"timeLeft"`
					CurrentPrice      string `json:"currentPrice"`
					CurrentPriceUSD   string `json:"currentPriceUSD"`
					SuggestedPrice    string `json:"suggestedPrice"`
					Seller            string `json:"seller"`
					ListingIndex      int    `json:"listingIndex"`
					State             string `json:"state"`
					Typename          string `json:"__typename"`
				} `json:"auction"`
				Typename string `json:"__typename"`
			} `json:"results"`
			Typename string `json:"__typename"`
		} `json:"axies"`
	} `json:"data"`
}



