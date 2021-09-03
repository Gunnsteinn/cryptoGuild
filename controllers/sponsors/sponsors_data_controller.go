package sponsors

import (
	"encoding/hex"
	"github.com/Gunnsteinn/cryptoGuild/domain/sponsor"
	"github.com/Gunnsteinn/cryptoGuild/services"
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
)

//
func getSponsorId(sponsorWalletParam string) (string, *errors.RestErr) {
	if _, addressHexErr := hex.DecodeString(sponsorWalletParam); addressHexErr == nil {
		return sponsorWalletParam, nil
	}
	if _, addressMailErr := mail.ParseAddress(sponsorWalletParam); addressMailErr == nil {
		return sponsorWalletParam, nil
	}
	return "", errors.NewBadRequestError("user id should be a hex or mail.")
}

// GetBy get an existing user.
func GetBy(c *gin.Context) {
	walletID, idErr := getSponsorId(c.Param("wallet_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.SponsorsService.GetSponsor(walletID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAll get all an existing user.
func GetAll(c *gin.Context) {
	user, getErr := services.SponsorsService.GetAllSponsor()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Create a new user.
func Create(c *gin.Context) {
	var sponsorModel sponsor.Sponsor
	if err := c.ShouldBindJSON(&sponsorModel); err != nil {
		restErr := errors.NewBadRequestError("invalid json body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.SponsorsService.CreateSponsor(sponsorModel)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Update a new user.
func Update(c *gin.Context) {
	var sponsorModel sponsor.Sponsor
	if err := c.ShouldBindJSON(&sponsorModel); err != nil {
		restErr := errors.NewBadRequestError("invalid json body.")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.SponsorsService.UpdateSponsor(sponsorModel)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Delete an existing user.
func Delete(c *gin.Context) {
	walletID, idErr := getSponsorId(c.Param("wallet_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	getErr := services.SponsorsService.DeleteSponsor(walletID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
