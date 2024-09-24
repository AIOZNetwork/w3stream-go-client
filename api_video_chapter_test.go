package w3streamsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var getVideoChaptersJSONResponse = `
	{
  "data": {
    "total": 0,
    "video_chapters": [
      {
        "language": "string",
        "url": "string"
      }
    ]
  },
  "status": "string"
}
`

var createVideoChapterJSONResponse = `{
  "data": {
    "video_chapter": {
      "language": "string",
      "url": "string"
    }
  },
  "status": "string"
}`

var videoChapter = VideoChapter{
	Language: PtrString("string"),
	Url:      PtrString("string"),
}

var getVideosChaptersResponse = GetVideoChaptersResponse{
	Data: &GetVideoChaptersData{
		Total: PtrInt32(0),
		VideoChapters: &[]VideoChapter{
			videoChapter,
		},
	},
	Status: PtrString("string"),
}

var createVideoChapterResponse = CreateVideoChapterResponse{
	Data: &CreateVideoChapterData{
		VideoChapter: &videoChapter,
	},
	Status: PtrString("string"),
}

func TestVideoChapter_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/videoId/chapter/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, createVideoChapterJSONResponse)
	})

	reader := strings.NewReader("data")
	resp, err := client.VideoChapter.Create("videoId", "en", "chapter.vtt", reader)
	if err != nil {
		t.Errorf("Chapter.Create error: %v", err)
	}

	expected := &createVideoChapterResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Chapter.Create\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestVideoChapter_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/videoId/chapters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, getVideoChaptersJSONResponse)
	})

	resp, err := client.VideoChapter.Get("videoId", VideoChapterApiGetRequest{})
	if err != nil {
		t.Errorf("Chapter.List error: %v", err)
	}

	expected := &getVideosChaptersResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Chapter.List\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestVideoChapter_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/videoId/chapter/en", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.VideoChapter.Delete("videoId", "en")
	if err != nil {
		t.Errorf("Chapter.Delete error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Chapter.Delete\n got=%#v\nwant=%#v", resp, expected)
	}
}
