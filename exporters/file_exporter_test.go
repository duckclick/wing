package exporters

import (
	"fmt"
	"github.com/duckclick/wing/config"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"encoding/base64"
	"github.com/duckclick/wing/trackentry"
)

func TestExport(t *testing.T) {
	exporterConfig := config.FileExporter{
		Folder: "/tmp/test/track_entries",
	}

	htmlSample := "<html><head></head><body></body></html>"
	htmlAsBase64 := base64.StdEncoding.EncodeToString([]byte(htmlSample))

	trackEntry := &trackentry.TrackEntry{
		CreatedAt: 123456,
		Markup:    htmlAsBase64,
	}

	recordId := uuid.NewV4().String()

	exporter := FileExporter{Config: exporterConfig}
	err := exporter.Export(trackEntry, recordId)
	assert.Nil(t, err, "export should succeed")

	recordPath := fmt.Sprintf("/tmp/test/track_entries/%s/%d.html", recordId, trackEntry.CreatedAt)
	if _, err := os.Stat(recordPath); os.IsNotExist(err) {
		fmt.Sprintf("FileExporter#Export failed to save track entry, expected:\n%v\n to exist", recordPath)
		t.FailNow()
	}
	defer os.Remove(recordPath)

	htmlBytes, _ := ioutil.ReadFile(recordPath)
	assert.Equal(t, "<html><head></head><body></body></html>", string(htmlBytes), "FileExporter#Export should save the expected content")
}
