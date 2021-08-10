package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {

	d := json.NewDecoder(r)

	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil

}
