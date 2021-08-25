package mongoDb

import "github.com/Gunnsteinn/cryptoGuild/domain/sponsor"

type Repository interface {
	FindOne(id int) *sponsor.Sponsor
}
