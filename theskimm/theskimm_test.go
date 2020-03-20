package theskimm_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tinee/newshub/theskimm"

	"github.com/matryer/is"
)

func TestParser(t *testing.T) {
	is := is.New(t)
	err := errors.New("HesakN")
	is.NoErr(err)

	p := theskimm.NewParser()

	srv := setupTheSkimmHTTPServer(strings.NewReader(``))
	srv.Start()
	p.Parse()
}

func setupTheSkimmHTTPServer(rd io.Reader) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, rd)
	}))
}
