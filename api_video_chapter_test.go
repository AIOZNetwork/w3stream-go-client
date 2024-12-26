package w3streamsdk

import (
	"io"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testLang              = "en"
	testVideoIDForChapter = "fdcaf04d-db3d-41af-8d2a-42f272ed556b"
	chapterContent        = `WEBVTT

00:00:00.000 --> 00:01:00.000
Chapter 1

00:01:00.000 --> 00:02:00.000
Chapter 2`
)

func createTempVTTFile(t *testing.T) *os.File {
	tmpFile, err := os.CreateTemp("", "test-*.vtt")
	if err != nil {
		t.Fatal(err)
	}

	_, err = tmpFile.WriteString(chapterContent)
	if err != nil {
		t.Fatal(err)
	}

	tmpFile.Seek(0, 0)
	return tmpFile
}

func TestVideoChapterService_Create(t *testing.T) {
	notExistId := uuid.New().String()
	tmpFile := createTempVTTFile(t)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tests := []struct {
		name    string
		videoID string
		lang    string
		file    *os.File
		wantErr bool
	}{
		{
			name:    "Valid Create",
			videoID: testVideoIDForChapter,
			lang:    testLang,
			file:    tmpFile,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			videoID: "invalid-id",
			lang:    testLang,
			file:    tmpFile,
			wantErr: true,
		},
		{
			name:    "Invalid Language",
			videoID: testVideoIDForChapter,
			lang:    "invalid",
			file:    tmpFile,
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			videoID: notExistId,
			lang:    testLang,
			file:    tmpFile,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader
			var fileName string
			if tt.file != nil {
				reader = tt.file
				fileName = tt.file.Name()
			}

			resp, err := testClient.VideoChapter.Create(tt.videoID, tt.lang, fileName, reader)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestVideoChapterService_Get(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		videoID string
		wantErr bool
	}{
		{
			name:    "Get other",
			videoID: testVideoIDForChapter,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.VideoChapter.Get(tt.videoID, VideoChapterApiGetRequest{})
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		videoID string
		request VideoChapterApiGetRequest
		wantErr bool
		checkFn func(*testing.T, *GetVideoChaptersResponse)
	}{
		{
			name:    "Valid Get",
			videoID: testVideoIDForChapter,
			request: VideoChapterApiGetRequest{}.
				Limit(10).
				Offset(0),
			wantErr: false,
			checkFn: func(t *testing.T, resp *GetVideoChaptersResponse) {
				assert.NotNil(t, resp.Data)
			},
		},
		{
			name:    "Invalid Video ID",
			videoID: "invalid-id",
			request: VideoChapterApiGetRequest{},
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			videoID: notExistId,
			request: VideoChapterApiGetRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.VideoChapter.Get(tt.videoID, tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				if tt.checkFn != nil {
					tt.checkFn(t, resp)
				}
			}
		})
	}
}

func TestVideoChapterService_Delete(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		videoID string
		lang    string
		wantErr bool
	}{
		{
			name:    "Delete other",
			videoID: testVideoIDForChapter,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.VideoChapter.Delete(tt.videoID, tt.lang)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		videoID string
		lang    string
		wantErr bool
	}{
		{
			name:    "Valid Delete",
			videoID: testVideoIDForChapter,
			lang:    testLang,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			videoID: "invalid-id",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Invalid Language",
			videoID: testVideoIDForChapter,
			lang:    "invalid",
			wantErr: true,
		},
		{
			name:    "Empty Video ID",
			videoID: "",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Empty Language",
			videoID: testVideoIDForChapter,
			lang:    "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			videoID: notExistId,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.VideoChapter.Delete(tt.videoID, tt.lang)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
