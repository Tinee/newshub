package theskimm_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/Tinee/newshub/theskimm"

	"github.com/matryer/is"
)

func TestParser(t *testing.T) {
	is := is.New(t)
	srv := setupTheSkimmHTTPServer(t, "theskimm.html")
	defer srv.Close()

	p := theskimm.NewParser(srv.URL)

	stories, err := p.Parse()
	is.NoErr(err)

	is.Equal(stories.Headline, "Daily Skimm: COVID-19 Funding, Climate Change, and the Return of the Dixie Chicks")
	is.Equal(len(stories.SubtitleToText), 12)
}

func setupTheSkimmHTTPServer(t *testing.T, path string) *httptest.Server {
	path, err := filepath.Abs(fmt.Sprintf("./testdata/%s", path))
	if err != nil {
		t.Fatal(err)
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		if _, err := io.Copy(w, f); err != nil {
			t.Fatal(err)
		}
	}))
}
