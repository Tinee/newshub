package theskimm

import (
	"io"
	"net/http"
	"strings"

	"github.com/Tinee/newshub"
	"golang.org/x/net/html"
)

type Parser struct {
	baseURL string
}

func NewParser(baseURL string) *Parser {
	return &Parser{
		baseURL: baseURL,
	}
}

func (p *Parser) Parse() ([]newshub.Story, error) {
	res, err := http.Get(p.baseURL)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, err
	}
	defer res.Body.Close()

	z := html.NewTokenizer(res.Body)

	var results []newshub.Story

	story := newshub.NewStory()
	var lastSubTitle string

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return results, nil
			}

			return nil, z.Err()

		case html.StartTagToken:
			tn, hasAttr := z.TagName()
			k, v, _ := z.TagAttr()

			tag := string(tn)

			if tag == "h1" {
				z.Next()
				story.Headline = string(z.Text())
			}

			if string(tn) == "h3" || !hasAttr {
				if string(k) != "class" && string(v) != "heading--section--small" {
					continue
				}

				z.Next()
				lastSubTitle = string(z.Text())
			}

			if string(tn) == "p" && hasAttr {
				if string(k) != "class" && string(v) != "copy" {
					continue
				}
				for {
					tt := z.Next()
					tn, _ := z.TagName()
					if tt == html.EndTagToken && string(tn) == "p" {
						break
					}

					if tt != html.TextToken {
						continue
					}

					content := story.SubtitleToText[lastSubTitle]
					content += strings.Trim(string(z.Text()), "\u00a0")
					story.SubtitleToText[lastSubTitle] = content
				}
			}
		}
	}
}
