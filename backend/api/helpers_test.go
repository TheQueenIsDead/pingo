// helpers_test.go contains helpers used for testing

package api

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http/httptest"
	"strings"
)

func SetupHttpRecorder(httpMethod, httpBody string) (echo.Context, *httptest.ResponseRecorder) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(httpMethod, "/", strings.NewReader(httpBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec

}

func SetupMockApplication() (*Application, sqlmock.Sqlmock) {

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
