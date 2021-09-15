package sponsors

import (
	"encoding/hex"
	"github.com/Gunnsteinn/cryptoGuild/domain/sponsor"
	"github.com/Gunnsteinn/cryptoGuild/services"
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
	"strings"
)

// GetByWalletId get an existing user.
func GetByWalletId(c *gin.Context) {
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

// GetByQuery get an existing user.
func GetByQuery(c *gin.Context) {
	filterValue, idErr := validateUriQueryStrings(c.Query("filter"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	filterKey, idErr := validateUriParams(c.Param("nick_name"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.SponsorsService.GetSponsorByQuery(filterKey, filterValue)
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

// Validations
func getSponsorId(sponsorWalletParam string) (string, *errors.RestErr) {
	if _, addressHexErr := hex.DecodeString(sponsorWalletParam); addressHexErr == nil {
		return sponsorWalletParam, nil
	}
	if _, addressMailErr := mail.ParseAddress(sponsorWalletParam); addressMailErr == nil {
		return sponsorWalletParam, nil
	}
	return "", errors.NewBadRequestError("user id should be a hex or mail.")
}

// Validations
func validateUriQueryStrings(data string) (string, *errors.RestErr) {
	if len(data) < 1 {
		return "", errors.NewBadRequestError("Uri QueryStrings have empty values.")
	}

	return strings.ToLower(data), nil
}

// Validations
func validateUriParams(data string) (string, *errors.RestErr) {
	if len(data) < 1 {
		return "", errors.NewNotFoundError("Uri Params have empty values.")
	}

	return strings.ToLower(data), nil
}
