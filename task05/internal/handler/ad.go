package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
)

func (h *Handler) createAd(c *gin.Context) {

	var input model.Ad
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.Name) > 200 {
		newErrorResponse(c, http.StatusBadRequest, "name should be no more than 200 symbols")
		return
	}

	if len(input.Description) > 1000 {
		newErrorResponse(c, http.StatusBadRequest, "description should be no more than 1000 symbols")
		return
	}

	id, err := h.services.Ad.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) listAds(c *gin.Context) {

	ads, err := h.services.Ad.List()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ads)
}

func (h *Handler) getAd(c *gin.Context) {

	adId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}

	item, err := h.services.Ad.Get(adId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
