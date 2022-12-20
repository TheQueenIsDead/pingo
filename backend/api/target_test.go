package api

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	createTargetJSON = `{"source":"https://github.com","frequency":5,"unit":"SECONDS"}`
	targetJSON       = `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null""source":"https://github.com","frequency":5,"unit":"SECONDS"}`
)

func TestCreateTarget(t *testing.T) {

	// Setup
	c, rec := SetupHttpRecorderPOST("", createTargetJSON)
	app, mock := SetupMockApplication()

	// Expect the following database activity
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `targets`").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()). // 6 args
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Assertions
	if assert.NoError(t, app.createTarget(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		//TODO: Figure out how to assert response when the time differs
		//assert.Equal(t, targetJSON, rec.Body.String())
	}
}

func TestGetTarget(t *testing.T) {

	// Setup
	c, rec := SetupHttpRecorderGET("")
	app, mock := SetupMockApplication()

	// Expect the following database activity
	mock.ExpectQuery("SELECT (.+) FROM `targets` WHERE `targets`.`deleted_at` IS NULL").
		WillReturnRows(sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "source", "frequency", "unit"}))

	// Assertions
	if assert.NoError(t, app.getTarget(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: Figure out how to assert response when the time differs
		//assert.Equal(t, targetJSON, rec.Body.String())
	}
}

func TestGetTargetById(t *testing.T) {

	// Setup
	c, rec := SetupHttpRecorderGET("/")
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	app, mock := SetupMockApplication()

	// Expect the following database activity
	mock.ExpectQuery("SELECT (.+) FROM `targets` WHERE `targets`.`id` = (.+) `targets`.`deleted_at` IS NULL").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "source", "frequency", "unit"}))

	// Assertions
	if assert.NoError(t, app.getTarget(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: Figure out how to assert response when the time differs
		//assert.Equal(t, targetJSON, rec.Body.String())
	}
}
