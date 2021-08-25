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
	GetUser(string) (*sponsor.Sponsor, *errors.RestErr)
	//CreateUser(sponsor.Sponsor) (*sponsor.Sponsor, *errors.RestErr)
	//UpdateUser(bool, sponsor.Sponsor) (*sponsor.Sponsor, *errors.RestErr)
	//DeleteUser(int64) *errors.RestErr
	//Search(string) (sponsor.Sponsors, *errors.RestErr)
}

func (s *sponsorService) GetUser(sponsorWallet string) (*sponsor.Sponsor, *errors.RestErr) {
	result := &sponsor.Sponsor{WalletAddress: sponsorWallet}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
