package newshub

func NewStory() *Story {
	return &Story{
		SubtitleToText: make(map[string]string),
	}
}

type Story struct {
	Headline       string
	SubtitleToText map[string]string
}

type Source interface {
	Parse() (*Story, error)
}
