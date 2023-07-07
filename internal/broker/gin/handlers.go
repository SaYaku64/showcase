/*
If you are using swagger - comments can be like that
*/

package gin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	SetInfoInput struct {
		Name string `json:"name"`
		Age  uint8  `json:"age"`
	}

	SetInfoOutput struct {
		Status string `json:"status"`
	}
)

type obj map[string]any

// GetInfo - handler that gets info, sometimes with key param
// @Summary     Request that gets info, sometimes with key param
// @Tags        v1
// @ID          info-get
// @Accept      json
// @Produce     json
// @Param  key   body string false   "example"
// @Failure  400 {object}     SetInfoOutput
// @Success  200 {object}      SetInfoOutput
// @Router /api/v1/info/get [get]
func (gm *ginModule) GetInfo(c *gin.Context) {
	param := c.Query("key")
	if param != "" {
		// some logic here

		resp := "GetInfo parameter is not empty" + param

		c.JSON(http.StatusBadRequest, obj{"error": resp})
		return
	}

	c.Status(http.StatusOK)
}

// SetInfo - handler that sets info
// @Summary     Request that gets info
// @Tags        v1
// @ID          info-set
// @Accept      json
// @Produce     json
// @Param  inputData body   SetInfoInput  true   "body of request"
// @Failure  400 {object}     obj
// @Success  200 {object}      obj "returns key"
// @Router /api/v1/info/set [post]
func (gm *ginModule) SetInfo(c *gin.Context) {
	input := SetInfoInput{}

	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, SetInfoOutput{"failed to get body"})
		return
	}

	if err := json.Unmarshal(data, &input); err != nil {
		c.JSON(http.StatusBadRequest, SetInfoOutput{"failed to unmarshal"})
		return
	}

	// do some manipulations

	c.JSON(http.StatusOK, SetInfoOutput{"success"})
}

// ========================= v2 =========================

// GetUserByID - handler that gets user by ID
// @Summary     Request that gets user by ID
// @Tags        v2
// @ID          user-id
// @Accept      json
// @Produce     json
// @Param  userID   body string false   "123"
// @Failure  400 {object}     obj
// @Success  200 {object}      obj
// @Router /api/v2/user/:userID [get]
func (gm *ginModule) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil || userID == 0 {
		c.JSON(http.StatusBadRequest, obj{"error": "Wrong userID"})
		return
	}

	// some logic here

	c.Status(http.StatusOK)
}
