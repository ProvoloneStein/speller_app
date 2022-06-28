package handler

import (
	"Nexign/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) postText(c *gin.Context) {
	var input model.Speller
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	response, err := h.services.Speller.CreateOne(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"text": response,
	})
}

func (h *Handler) postMany(c *gin.Context) {
	var input model.Spellers

	//body, err := ioutil.ReadAll(c.Request.Body)
	//
	//fmt.Println(string(body))
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	response, err := h.services.Speller.CreateMany(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"texts": response,
	})
}
