package services

import (
	"fmt"
	"github.com/Gunnsteinn/cryptoGuild/client"
	"github.com/Gunnsteinn/cryptoGuild/domain"
	"time"
)

func GetAdventurerDetail(userID string) (*domain.Adventurer, error) {
	result, err := client.FetchClaimInfo(userID)
	if err != nil {
		return nil, err
	}

	lastClaimedAt := time.Unix(0, (result.LastClaimedItemAt)*int64(time.Second)).Format(time.RFC3339)
	ClaimedAt := time.Unix(0, (result.LastClaimedItemAt+1209600)*int64(time.Second)).Format(time.RFC3339)
	days, performance := performanceCalculator(lastClaimedAt, result.Total)

	adv := domain.Adventurer{ClientID: result.ClientID, Total: result.Total, Performance: performance, ClaimableTotal: result.ClaimableTotal, LastClaimedItemAt: lastClaimedAt, PlayedDays: days, ClaimedItemAt: ClaimedAt}

	return &adv, nil
}

func GetAllAdventurerDetail(listOfUser domain.BodyRequestListOfAdventurer) (*[]domain.ListOfAdventurer, error) {

	var items []domain.ListOfAdventurer

	for _, userData := range listOfUser {
		result, err := GetAdventurerDetail(userData.ClientID)
		if err != nil {
			return nil, err
		}

		auxListOfAdventurer := domain.ListOfAdventurer{User: userData.User, Data: result}
		items = append(items, auxListOfAdventurer)
	}

	return &items, nil
}

func performanceCalculator(lastClaimedItemAt string, total int) (float64, float64) {
	t, err := time.Parse(time.RFC3339, lastClaimedItemAt)
	if err != nil {
		fmt.Println(err)
	}

	days := time.Now().Sub(t).Hours() / 24
	performance := float64(total) / days

	return days, performance
}
