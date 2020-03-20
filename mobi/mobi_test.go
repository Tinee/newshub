package mobi_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Tinee/newshub"
	"github.com/Tinee/newshub/mobi"
	"github.com/matryer/is"
)

func TestConvert(t *testing.T) {
	is := is.New(t)

	story := newshub.Story{
		Headline: "Test Headline",
		SubtitleToText: map[string]string{
			"subtitle1": "content1",
			"subtitle2": "content2",
		},
	}
	rd, err := mobi.Convert(&story)
	is.NoErr(err)

	bs, err := ioutil.ReadAll(rd)
	is.NoErr(err)

	is.True(bytes.Contains(bs, []byte("Test Headline")))
	is.True(bytes.Contains(bs, []byte("subtitle1")))
	is.True(bytes.Contains(bs, []byte("content1")))
	is.True(bytes.Contains(bs, []byte("subtitle2")))
	is.True(bytes.Contains(bs, []byte("content2")))
	is.True(!bytes.Contains(bs, []byte("This should not be in here.")))
}
