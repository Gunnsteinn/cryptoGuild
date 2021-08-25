package sponsors

import (
	"encoding/hex"
	"github.com/Gunnsteinn/cryptoGuild/services"
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func getSponsorId(sponsorWalletParam string) (string, *errors.RestErr) {
	_, userErr := hex.DecodeString(sponsorWalletParam)
	if userErr != nil {
		return "", errors.NewBadRequestError("user id should be a hex.")
	}
	return sponsorWalletParam, nil
}

// Create create a new user.
func Create(c *gin.Context) {
	//var user users.User
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	restErr := errors.NewBadRequestError("invalid json body.")
	//	c.JSON(restErr.Status, restErr)
	//	return
	//}
	//
	//result, saveErr := services.UsersService.CreateUser(user)
	//if saveErr != nil {
	//	c.JSON(saveErr.Status, saveErr)
	//	return
	//}
	//c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get get an existing user.
func Get(c *gin.Context) {
	walletID, idErr := getSponsorId(c.Param("wallet_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.SponsorsService.GetUser(walletID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update update an exist user.
func Update(c *gin.Context) {
	//userID, idErr := getUserId(c.Param("user_id"))
	//if idErr != nil {
	//	c.JSON(idErr.Status, idErr)
	//	return
	//}
	//
	//var user users.User
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	restErr := errors.NewBadRequestError("invalid json body.")
	//	c.JSON(restErr.Status, restErr)
	//	return
	//}
	//
	//user.ID = userID
	//
	//// Decide if it's a full update (PUT) or a partial update (PATCH).
	//isPartial := c.Request.Method == http.MethodPatch
	//
	//result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	//if updateErr != nil {
	//	c.JSON(updateErr.Status, updateErr)
	//	return
	//}
	//c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Delete delete an exist user.
func Delete(c *gin.Context) {
	//userID, idErr := getUserId(c.Param("user_id"))
	//if idErr != nil {
	//	c.JSON(idErr.Status, idErr)
	//	return
	//}
	//
	//if deleteErr := services.UsersService.DeleteUser(userID); deleteErr != nil {
	//	c.JSON(deleteErr.Status, deleteErr)
	//	return
	//}
	//c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search find all the user by status
func Search(c *gin.Context) {
	//status := c.Query("status")
	//
	//users, searchErr := services.UsersService.Search(status)
	//if searchErr != nil {
	//	c.JSON(searchErr.Status, searchErr)
	//	return
	//}
	//c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
