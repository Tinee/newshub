package mobi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/766b/mobi"

	"github.com/Tinee/newshub"
)

func Convert(s *newshub.Story) (io.Reader, error) {
	fileName := fmt.Sprintf("output-%s.mobi", time.Now().Format("06-01-02"))

	m, err := mobi.NewWriter(fileName)
	if err != nil {
		return nil, err
	}

	m.Title("Newshub")
	m.Compression(mobi.CompressionNone)
	m.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	m.NewExthRecord(mobi.EXTH_AUTHOR, "Marcus Karlsson")

	m.NewChapter(s.Headline, []byte(toMobiText(s)))

	m.Write()

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = os.Remove(fileName)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bs), nil
}

func toMobiText(s *newshub.Story) string {
	var result string
	for title, content := range s.SubtitleToText {
		result += fmt.Sprintf(`<h3>%s<\h3><p>%s<\p>`, title, content)
	}

	return result
}
