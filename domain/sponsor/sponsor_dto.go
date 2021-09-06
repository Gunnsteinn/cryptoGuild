package sponsor

import (
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"strings"
)

type Sponsor struct {
	//ID            primitive.ObjectID `bson:"_id" json:"id"`
	Name          string `bson:"name,omitempty" json:"name,omitempty"`
	LastName      string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email         string `bson:"email,omitempty" json:"email,omitempty"`
	WalletAddress string `bson:"wallet_address,omitempty" json:"wallet_address,omitempty"`
	Teams         []struct {
		TeamName      string  `bson:"team_name,omitempty" json:"team_name,omitempty"`
		WalletAddress string  `bson:"wallet_address,omitempty" json:"wallet_address,omitempty"`
		PoolPercent   float64 `bson:"pool_percent,omitempty" json:"pool_percent,omitempty"`
		Adventurer    struct {
			Name              string  `bson:"name,omitempty" json:"name,omitempty"`
			LastName          string  `bson:"last_name,omitempty" json:"last_name,omitempty"`
			WalletAddress     string  `bson:"wallet_address,omitempty" json:"wallet_address,omitempty"`
			ProfitSlp         int     `bson:"profit_slp,omitempty" json:"profit_slp,omitempty"`
			Performance       float64 `bson:"performance,omitempty" json:"performance,omitempty"`
			LastClaimedItemAt string  `bson:"last_claimed_item_at,omitempty" json:"last_claimed_item_at,omitempty"`
			PlayedDays        float64 `bson:"played_days,omitempty" json:"played_days,omitempty"`
		} `bson:"Adventurer,omitempty" json:"Adventurer,omitempty"`
	} `bson:"teams,omitempty" json:"teams,omitempty"`
}

type Sponsors []Sponsor

// Validate method implements User struct and validate if the email is OK.
func (sponsor *Sponsor) Validate() *errors.RestErr {
	sponsor.Name = strings.TrimSpace(sponsor.Name)
	sponsor.LastName = strings.TrimSpace(sponsor.LastName)

	sponsor.Email = strings.TrimSpace(strings.ToLower(sponsor.Email))
	if sponsor.Email == "" {
		return errors.NewBadRequestError("invalid email address.")
	}

	//TODO Need more validations
	return nil
}
