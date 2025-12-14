package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}
func (h Handler) RefreshAccessToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh-token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "refresh token not found",
		})
		return
	}

	jwtToken, err := ValidateToken(refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	// get the user from token
	userID, err := jwtToken.Claims.GetSubject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	tokens, err := GenerateTokens(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.SetCookie("refresh-token", tokens.RefreshToken, 24, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokens.AccessToken,
	})

}
