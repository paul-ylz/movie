package location

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type handlerTestSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *handlerTestSuite) SetupTest() {
	r := http.NewServeMux()
	r.HandleFunc("/search", Search())
	s.server = httptest.NewServer(r)
}

func (s *handlerTestSuite) TearDownTest() {
	s.server.Close()
}

func (s *handlerTestSuite) TestReturnsOKGivenAValidLocation() {
	names := []string{"Bangalore", "Mumbai"}
	for _, name := range names {
		url := s.server.URL + "/search?location=" + name

		resp, err := http.Get(url)

		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		s.NoError(err)
		s.Contains(buf.String(), name, fmt.Sprintf("expected to find %q in %q", name, buf.String()))
	}
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
