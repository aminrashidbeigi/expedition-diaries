package endpoints

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EndpointsTestSuite struct {
	suite.Suite
	endpoints Router
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func (suite *EndpointsTestSuite) TestGetTravelsBySlugEndpointWithEmptySlug() {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	ctx.Request.Method = "GET"
	params := []gin.Param{
		{
			Key:   "slug",
			Value: "",
		},
	}

	ctx.Params = params
	suite.endpoints.GetTravelBySlug(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func TestEndpointTestSuite(t *testing.T) {
	suite.Run(t, new(EndpointsTestSuite))
}
