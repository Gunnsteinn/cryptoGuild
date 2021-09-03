package client

import (
	"encoding/json"
	"github.com/Gunnsteinn/cryptoGuild/domain"
	"io"
	"log"
	"net/http"
)

func FetchClaimInfo(id string) (*domain.ClaimInfo, error) {
	URL := "https://game-api.skymavis.com/game-api/clients/" + id + "/items/1"
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var cResp domain.ClaimInfo
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return nil, err
	}
	//Invoke the text output function & return it with nil as the error value
	return &cResp, nil
}
