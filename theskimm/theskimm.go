package theskimm

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Tinee/newshub"
	"golang.org/x/net/html"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(url string) ([]newshub.Story, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	defer res.Body.Close()

	z := html.NewTokenizer(res.Body)

	var results []newshub.Story

loop:
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				// break the look because end of file indicates success
				break loop
			}
			return nil, z.Err()
		case html.StartTagToken:

			tn, hasAttr := z.TagName()
			k, v, _ := z.TagAttr()

			if string(tn) == "h3" || !hasAttr {
				if string(k) != "class" && string(v) != "heading--section--small" {
					continue
				}
				// Found an headline, let's extract it.
				headline, err := extractContent(z)
				if err != nil {
					return nil, err
				}
				results = append(results, newshub.Story{Headline: headline})

				continue
			}

			if string(tn) != "p" || !hasAttr {
				continue
			}

			if string(k) != "class" && string(v) != "copy" {
				continue
			}
			// Found the content, let's extract it.
			content, err := extractContent(z)
			if err != nil {
				return nil, err
			}

			latest := &results[len(results)-1]
			latest.Content = content

			//story := newshub.Story{}

			//ok := isElementWithClass(z, "h3", "heading--section--small")
			//if ok {
			//	if len(results) > 0 {
			//		latest := &results[len(results)-1]
			//		latest.Content = tempText
			//		tempText = ""
			//	}

			//	title, err := extractContent(z)
			//	if err != nil {
			//		return nil, err
			//	}
			//	story.Headline = title

			//	results = append(results, story)
			//}

			//ok = isElementWithClass(z, "p", "copy")
			//if ok {
			//	content, err := extractContent(z)
			//	if err != nil {
			//		return nil, err
			//	}
			//	tempText += content
			//}

		}
	}
	return results, nil
}

func isElementWithClass(t *html.Tokenizer, element, class string) bool {
	tn, hasAttr := t.TagName()
	if string(tn) != element || !hasAttr {
		return false
	}

	k, v, _ := t.TagAttr()
	if string(k) != "class" && string(v) != class {
		return false
	}

	return true
}

func extractContent(t *html.Tokenizer) (string, error) {
	var result string

	if t.Token().Type != html.StartTagToken {
		return "", fmt.Errorf("error: expected token to be %q but got %q", html.StartTagToken, t.Token())
	}
	for {
		tt := t.Next()

		if tt == html.EndTagToken {
			// Go to the next token because we have already seen this one.
			t.Next()
			break
		}

		if tt != html.TextToken {
			continue
		}
		result += string(t.Text())

	}
	return result, nil
}
