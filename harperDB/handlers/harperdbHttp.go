package handlers

import (
	"log"
	"net/http"
	"regexp"

	harperdb "github.com/HarperDB-Add-Ons/sdk-go"
	"github.com/call-forwarding/database"
	"github.com/call-forwarding/harperDB/models"
	"github.com/labstack/echo/v4"
)

var (
	badRequestStatus  = http.StatusBadRequest
	badRequestCode    = "INVALID_ARGUMENT"
	badRequestMessage = "Client specified an invalid argument, request body or query param."

	internalServerStatus  = http.StatusInternalServerError
	internalServerCode    = "INTERNAL_SERVER_ERROR"
	internalServerMessage = "Unable to process the request as server is down."

	notFoundStatus  = http.StatusNotFound
	notFoundCode    = "NOT_FOUND"
	notFoundMessage = "The specified resource is not found."
)

type HarperHttpHandler struct {
	db database.Database
}

func NewHarperHttpHandler(db database.Database) HarperdbHandler {
	return &HarperHttpHandler{db: db}
}

func (h *HarperHttpHandler) CfStatus(c echo.Context) error {
	// Read body
	bd := models.CallForwardingBody{}
	err := c.Bind(&bd)
	if err != nil {
		log.Printf("Unable to parse request body. Error - [%s]", err.Error())
		return c.JSON(http.StatusBadRequest, models.CallForwardStatusResp{
			Status:  &badRequestStatus,
			Code:    &badRequestCode,
			Message: &badRequestMessage,
		})
	}

	// Check for valid phone number string
	pattern := `^\+\d{4,14}$`
	rex, err := regexp.Compile(pattern)
	if err != nil {
		log.Printf("error while compiling regular expression. error - [%s]", err.Error())
		return c.JSON(http.StatusBadRequest, models.CallForwardStatusResp{
			Status:  &badRequestStatus,
			Code:    &badRequestCode,
			Message: &badRequestMessage,
		})
	}
	if !rex.MatchString(bd.PhoneNumber) {
		log.Printf("Phone number does not match regular expression. Invalid phone number provided [%s]", bd.PhoneNumber)
		return c.JSON(http.StatusBadRequest, models.CallForwardStatusResp{
			Status:  &badRequestStatus,
			Code:    &badRequestCode,
			Message: &badRequestMessage,
		})
	}

	// Get details from the harperdb
	phoneDetails := []models.CallForwardStatusResp{}
	schemaName := "data"
	dbName := "callforwards"
	searchAttribute := "phoneNumber"
	searchValue := bd.PhoneNumber
	attributeList := []string{"*"}

	err = h.db.GetDb().SearchByValue(schemaName, dbName, &phoneDetails, harperdb.Attribute(searchAttribute), searchValue, harperdb.AttributeList(attributeList))
	if err != nil {
		log.Println("Error while reading phone details from database. Error - ", err.Error())
		return c.JSON(http.StatusInternalServerError, models.CallForwardStatusResp{
			Status:  &internalServerStatus,
			Code:    &internalServerCode,
			Message: &internalServerMessage,
		})
	}

	// Iterate over the list
	for _, phn := range phoneDetails {
		if phn.PhoneNumber != nil && *phn.PhoneNumber == bd.PhoneNumber && phn.Active != nil {
			phn.PhoneNumber = nil
			return c.JSON(http.StatusOK, phn)
		}
	}

	return c.JSON(http.StatusInternalServerError, models.CallForwardStatusResp{
		Status:  &notFoundStatus,
		Code:    &notFoundCode,
		Message: &notFoundMessage,
	})
}
