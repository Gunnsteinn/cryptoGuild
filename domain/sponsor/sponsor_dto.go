package sponsor

import (
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Sponsor struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	LastName      string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	WalletAddress string             `json:"wallet_address,omitempty" bson:"wallet_address,omitempty"`
	Teams         []struct {
		TeamName      string `json:"team_name,omitempty" bson:"team_name,omitempty"`
		WalletAddress string `json:"wallet_address,omitempty" bson:"wallet_address,omitempty"`
		PoolPercent   int    `json:"pool_percent,omitempty" bson:"pool_percent,omitempty"`
		Adventurer    struct {
			Name              string    `json:"name,omitempty" bson:"name,omitempty"`
			LastName          string    `json:"last_name,omitempty" bson:"last_name,omitempty"`
			WalletAddress     string    `json:"wallet_address,omitempty" bson:"wallet_address,omitempty"`
			ProfitSlp         int       `json:"profit_slp,omitempty" bson:"profit_slp,omitempty"`
			Performance       int       `json:"performance,omitempty" bson:"performance,omitempty"`
			LastClaimedItemAt time.Time `json:"last_claimed_item_at,omitempty" bson:"last_claimed_item_at,omitempty"`
			PlayedDays        int       `json:"played_days,omitempty" bson:"played_days,omitempty"`
		} `bson:"Adventurer,omitempty"`
	} `json:"teams,omitempty" bson:"teams,omitempty"`
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
