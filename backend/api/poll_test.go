package api

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var (
	createPollJSON = `{"target_ids":[1],"duration":5}`
	pollJSON       = `{"ID": 7,"CreatedAt": "2022-12-21T11:06:52.151093088+13:00","UpdatedAt": "2022-12-21T11:06:52.151093088+13:00","DeletedAt": null,"target_ids": [1],"duration": 5,"started": "0001-01-01T00:00:00Z"}`
)

func TestCreatePoll(t *testing.T) {

	// Setup
	c, rec := SetupHttpRecorderPOST("", createPollJSON)
	app, mock := SetupMockApplication()

	// Expect the following database activity
	mock.ExpectBegin()

	// TODO: Remove the insert into targets call where possible...

	// Create the poll
	mock.ExpectExec("INSERT INTO `polls`").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()). // 5 args
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Attempt to insert target
	mock.ExpectExec("INSERT INTO `targets`").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()). // 6 args
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Attempt to insert poll_targets relationship table
	mock.ExpectExec("INSERT INTO `poll_targets`").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()). // 2 args
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	// Assertions
	if assert.NoError(t, app.createPoll(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		//TODO: Figure out how to assert response when the time differs
		//assert.Equal(t, targetJSON, rec.Body.String())
	}
}

func TestGetPoll(t *testing.T) {

	t.SkipNow()
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

func TestGetPollById(t *testing.T) {

	t.SkipNow()
	// Setup
	c, rec := SetupHttpRecorderGET("/")
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	app, mock := SetupMockApplication()

	// Expect the following database activity
	mock.ExpectQuery("SELECT (.+) FROM `targets` WHERE `targets`.`id` = (.+) `targets`.`deleted_at` IS NULL").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "source", "frequency", "unit"}).
				AddRow("1", time.Now(), time.Now(), time.Now(), "https://example.com", rand.Int(), "SECONDS"))

	// Assertions
	if assert.NoError(t, app.getTarget(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: Figure out how to assert response when the time differs
		//assert.Equal(t, targetJSON, rec.Body.String())
	}
}

func TestDeletePoll(t *testing.T) {

	t.SkipNow()

	// Setup
	c, rec := SetupHttpRecorderDELETE("/")
	app, mock := SetupMockApplication()

	// Expect the following database activity
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `targets`").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Assertions
	if assert.NoError(t, app.deleteTarget(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: Figure out how to assert response when the time differs
		//assert.Equal(t, targetJSON, rec.Body.String())
	}
}

func TestDeletePollById(t *testing.T) {

	t.SkipNow()

	// Setup
	c, rec := SetupHttpRecorderDELETE("/")
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	app, mock := SetupMockApplication()

	// Expect the following database activity
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `targets`").
		WithArgs(sqlmock.AnyArg(), "1").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Assertions
	if assert.NoError(t, app.deleteTarget(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//TODO: Figure out how to assert response when the time differs
		//assert.Equal(t, targetJSON, rec.Body.String())
	}
}
