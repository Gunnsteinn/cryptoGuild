package client

import (
	"encoding/json"
	"github.com/Gunnsteinn/cryptoGuild/domain"
	"log"
	"net/http"
)

func FetchClaimInfo(id string) (*domain.ClaimInfo, error) {
	URL := "https://game-api.skymavis.com/game-api/clients/" + id + "/items/1"
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cResp domain.ClaimInfo
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	//Invoke the text output function & return it with nil as the error value
	return &cResp, nil
}
