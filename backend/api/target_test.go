package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"pingo/models"
	"strings"
	"testing"
)

type mockDB struct {
	mock.Mock
}

func (m mockDB) CreateTarget(ctx context.Context, t *models.Target) (models.Target, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockDB) ReadTarget(ctx context.Context, ID int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockDB) UpdateTarget(ctx context.Context, ID int, t models.Target) (models.Target, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockDB) DeleteTarget(ctx context.Context, ID int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func TestCreateTarget(t *testing.T) {

	var testCases = []struct {
		name         string
		whenURL      string
		whenBody     string
		expectBody   string
		expectStatus int
	}{
		{
			name:         "ok",
			whenURL:      "/api/target",
			whenBody:     `{"target_config":{"type":"x","source":"y"},"interval":{"frequency":1,"unit":"seconds"}}`,
			expectStatus: http.StatusOK,
			expectBody:   "{\"target_config\":{\"type\":\"x\",\"source\":\"y\"},\"interval\":{\"frequency\":1,\"unit\":\"seconds\"}}",
		},
		//{
		//	name:         "nok",
		//	whenURL:      "/123",
		//	whenBody:     `{"name":"John Doe","email":"x"}`,
		//	expectStatus: http.StatusBadRequest,
		//	expectBody:   "{\"message\":\"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag\"}\n",
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			e := echo.New()

			h := handler{}
			e.POST("/api/target", h.create)

			req := httptest.NewRequest(http.MethodPost, tc.whenURL, strings.NewReader(tc.whenBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectStatus, rec.Code)
			assert.Equal(t, tc.expectBody, rec.Body.String())
		})
	}
}

//
//func TestCreateTarget(t *testing.T) {
//
//	targetJSON := `{"target_config":{"type":"x","source":"y"},"interval":{"frequency":1,"unit":"seconds"}}`
//
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPost, "/api/target", strings.NewReader(targetJSON))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//
//	//c.SetParamNames("id")
//	//c.SetParamValues(tc.whenID)
//
//	mockedDB := new(mockDB)
//	//mockedDB.On("GetUsernameByID", mock.Anything, tc.resultID).Return(tc.resultUsername, tc.resultError)
//
//	targets := handler{mockedDB}
//	_ = targets.create(c)
//
//	//assert.Equal(t, tc.expect, rec.Body.String())
//	//if tc.expectErr != "" {
//	//	assert.EqualError(t, err, tc.expectErr)
//	//} else {
//	//	assert.NoError(t, err)
//	//}
//	mockedDB.AssertExpectations(t)
//
//	assert.Equal(t, http.StatusCreated, rec.Code)
//	assert.Equal(t, targetJSON, rec.Body.String())
//
//}

//package target
//
//import (
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"pingo"
//	"strings"
//	"testing"
//)
//
//var (
//	targetJSON = `{"target_config":{"type":"x","source":"y"},"interval":{"frequency":1,"unit":"seconds"}}`
//
//	mockDB = map[int]*types.Target{
//		0: {
//			TargetConfig: types.TargetConfig{Type: "x", Source: "y"},
//			Interval:     types.Interval{Frequency: 1, Unit: "z"},
//		},
//	}
//)
//
//func TestCreateTarget(t *testing.T) {
//	// Assemble
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPost, main.TargetEndpoint, strings.NewReader(targetJSON))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//
//	// Act
//	_ = CreateTarget(c)
//
//	// Assert
//	assert.Equal(t, http.StatusCreated, rec.Code)
//	assert.Equal(t, targetJSON+main.NewLineCharacter, rec.Body.String())
//}
//
//func TestReadTarget(t *testing.T) {
//	// Assemble
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPost, main.TargetEndpoint, strings.NewReader(targetJSON))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//
//	// Act
//	_ = ReadTarget(c)
//
//	// Assert
//	assert.Equal(t, http.StatusCreated, rec.Code)
//	assert.Equal(t, targetJSON+main.NewLineCharacter, rec.Body.String())
//}
