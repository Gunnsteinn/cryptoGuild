package services

import (
	"github.com/Gunnsteinn/cryptoGuild/client"
	"github.com/Gunnsteinn/cryptoGuild/domain"
	"time"
)

func GetAdventurerDetail(userID string) (*domain.Adventurer, error) {
	result, err := client.FetchClaimInfo(userID)
	if err != nil {
		return nil, err
	}

	lastClaimedAt := time.Unix(0, (result.LastClaimedItemAt)*int64(time.Second)).Format("2006-01-02 15:04:05")
	ClaimedAt := time.Unix(0, (result.LastClaimedItemAt+1209600)*int64(time.Second)).Format("2006-01-02 15:04:05")

	adv := domain.Adventurer{ClientID: result.ClientID, Total: result.Total, ClaimableTotal: result.ClaimableTotal, LastClaimedItemAt: lastClaimedAt, ClaimedItemAt: ClaimedAt}

	return &adv, nil
}
