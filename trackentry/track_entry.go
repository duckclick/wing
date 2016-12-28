package trackentry

import (
	"encoding/base64"
	"github.com/pkg/errors"
)

type TrackEntry struct {
	CreatedAt int    `json:"created_at"`
	Origin    string `json:"origin"`
	Markup    string `json:"markup"`
}

func (trackEntry *TrackEntry) MarkupBytes() ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(trackEntry.Markup)
	return bytes, errors.Wrap(err, "Invalid base64 payload")
}