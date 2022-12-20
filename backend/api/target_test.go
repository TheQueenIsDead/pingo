package api

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupMockApplication() (*Application, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(
			sqlmock.NewRows([]string{"sqlite_version()"}).
				AddRow(""))

	gdb, err := gorm.Open(&sqlite.Dialector{
		DriverName: sqlite.DriverName,
		DSN:        "",
		Conn:       db,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &Application{db: gdb}, mock
}

var (
	createTargetJSON = `{"source":"https://github.com","frequency":5,"unit":"SECONDS"}`
	targetJSON       = `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null""source":"https://github.com","frequency":5,"unit":"SECONDS"}`
)

func TestCreateTarget(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createTargetJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	app, mock := setupMockApplication()

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
