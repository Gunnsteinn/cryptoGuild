package services

import (
	"github.com/Gunnsteinn/cryptoGuild/domain/sponsor"
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
)

var (
	SponsorsService sponsorsServiceInterface = &sponsorService{}
)

type sponsorService struct {
}

type sponsorsServiceInterface interface {
	GetSponsor(string) (*sponsor.Sponsor, *errors.RestErr)
	GetAllSponsor() (sponsor.Sponsors, *errors.RestErr)
	CreateSponsor(sponsor.Sponsor) (*sponsor.Sponsor, *errors.RestErr)
	UpdateSponsor(sponsor.Sponsor) (*sponsor.Sponsor, *errors.RestErr)
	DeleteSponsor(string) *errors.RestErr
}

// GetSponsor is the business logic to get user from the db.
func (s *sponsorService) GetSponsor(sponsorWallet string) (*sponsor.Sponsor, *errors.RestErr) {
	result := &sponsor.Sponsor{WalletAddress: sponsorWallet}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllSponsor is the business logic to get all user from the db.
func (s *sponsorService) GetAllSponsor() (sponsor.Sponsors, *errors.RestErr) {
	result := &sponsor.Sponsor{}
	return result.GetAll()
}

// CreateSponsor is the business logic to create a new user from the database.
func (s *sponsorService) CreateSponsor(sponsor sponsor.Sponsor) (*sponsor.Sponsor, *errors.RestErr) {
	if err := sponsor.Validate(); err != nil {
		return nil, err
	}

	if errDetail := getAdventureDetail(sponsor); errDetail != nil {
		return nil, errDetail
	}

	if err := sponsor.Create(); err != nil {
		return nil, err
	}
	return &sponsor, nil
}

// UpdateSponsor is the business logic to update a new user from the database.
func (s *sponsorService) UpdateSponsor(sponsor sponsor.Sponsor) (*sponsor.Sponsor, *errors.RestErr) {

	if errDetail := getAdventureDetail(sponsor); errDetail != nil {
		return nil, errDetail
	}

	if err := sponsor.Update(); err != nil {
		return nil, err
	}
	return &sponsor, nil
}

// DeleteSponsor is the business logic to delete user from the db.
func (s *sponsorService) DeleteSponsor(sponsorWallet string) *errors.RestErr {
	result := &sponsor.Sponsor{WalletAddress: sponsorWallet}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil
}

func getAdventureDetail(sponsor sponsor.Sponsor) *errors.RestErr {
	for i, team := range sponsor.Teams {
		result, err := GetAdventurerDetail(team.WalletAddress)
		if err != nil {
			return errors.NewBadRequestError(err.Error())
		}

		sponsor.Teams[i].Adventurer.ProfitSlp = result.Total
		sponsor.Teams[i].Adventurer.Performance = result.Performance
		sponsor.Teams[i].Adventurer.LastClaimedItemAt = result.LastClaimedItemAt
		sponsor.Teams[i].Adventurer.PlayedDays = result.PlayedDays
	}
	return nil
}
