package w3streamsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	// "strings"
	"testing"
)

var getVideoListJSONResponse = `
{
  "data": {
    "total": 0,
    "videos": [
      {
        "assets": {
          "hls_url": "string",
          "mp4_url": "string",
          "source_url": "string",
          "thumbnail_url": "string"
        },
        "chapters": [
          {
            "language": "string",
            "url": "string"
          }
        ],
        "created_at": "string",
        "description": "string",
        "duration": 0,
        "id": "string",
        "is_mp4": true,
        "is_panoramic": true,
        "is_public": true,
        "metadata": [
          {
            "key": "string",
            "value": "string"
          }
        ],
        "player_theme": {
          "asset": {
            "logo": "string",
            "logo_image_link": "string",
            "logo_link": "string"
          },
          "controls": {
            "enable_api": true,
            "enable_controls": true,
            "force_autoplay": true,
            "force_loop": true,
            "hide_title": true
          },
          "created_at": "string",
          "id": "string",
          "name": "string",
          "theme": {
            "captions_background": "string",
            "controls_background": "string",
            "main_color": "string",
            "main_color_active": "string",
            "menu_background": "string",
            "text_control_color": "string",
            "text_control_hover": "string",
            "text_menu_color": "string",
            "thumb_height": "string",
            "track_background": "string",
            "track_height": "string",
            "track_played": "string",
            "track_unplayed": "string",
            "video_background": "string"
          },
          "user_id": "string"
        },
        "player_theme_id": "string",
        "qualities": [
          {
            "name": "string",
            "status": "string",
            "type": "string"
          }
        ],
        "size": 0,
        "status": "string",
        "captions": [
          {
            "is_default": true,
            "language": "string",
            "url": "string"
          }
        ],
        "tags": [
          "string"
        ],
        "title": "string",
        "user_id": "string"
      }
    ]
  },
  "status": "string"
}
`

var getVideoJSONResponse = `
{
  "data": {
    "video": {
      "assets": {
        "hls_url": "string",
        "mp4_url": "string",
        "source_url": "string",
        "thumbnail_url": "string"
      },
      "chapters": [
        {
          "language": "string",
          "url": "string"
        }
      ],
      "created_at": "string",
      "description": "string",
      "duration": 0,
      "id": "string",
      "is_mp4": true,
      "is_panoramic": true,
      "is_public": true,
      "metadata": [
        {
          "key": "string",
          "value": "string"
        }
      ],
      "player_theme": {
        "asset": {
          "logo": "string",
          "logo_image_link": "string",
          "logo_link": "string"
        },
        "controls": {
          "enable_api": true,
          "enable_controls": true,
          "force_autoplay": true,
          "force_loop": true,
          "hide_title": true
        },
        "created_at": "string",
        "id": "string",
        "name": "string",
        "theme": {
          "captions_background": "string",
          "controls_background": "string",
          "main_color": "string",
          "main_color_active": "string",
          "menu_background": "string",
          "text_control_color": "string",
          "text_control_hover": "string",
          "text_menu_color": "string",
          "thumb_height": "string",
          "track_background": "string",
          "track_height": "string",
          "track_played": "string",
          "track_unplayed": "string",
          "video_background": "string"
        },
        "user_id": "string"
      },
      "player_theme_id": "string",
      "qualities": [
        {
          "name": "string",
          "status": "string",
          "type": "string"
        }
      ],
      "size": 0,
      "status": "string",
      "captions": [
        {
          "is_default": true,
          "language": "string",
          "url": "string"
        }
      ],
      "tags": [
        "string"
      ],
      "title": "string",
      "user_id": "string"
    }
  },
  "status": "string"
}
`

var getVideoTransCodingJSONResponse = `
{
  "data": {
    "is_enough": true,
    "price": 0
  },
  "status": "string"
}
`

var video = Video{
	Assets: &VideoAssets{
		HlsUrl:       PtrString("string"),
		Mp4Url:       PtrString("string"),
		SourceUrl:    PtrString("string"),
		ThumbnailUrl: PtrString("string"),
	},
	Chapters: &[]VideoChapter{
		{
			Language: PtrString("string"),
			Url:      PtrString("string"),
		},
	},
	CreatedAt:   PtrString("string"),
	Description: PtrString("string"),
	Duration:    PtrFloat32(0),
	Id:          PtrString("string"),
	IsMp4:       PtrBool(true),
	IsPublic:    PtrBool(true),
	Metadata: &[]Metadata{
		{
			Key:   PtrString("string"),
			Value: PtrString("string"),
		},
	},
	PlayerTheme: &PlayerTheme{
		Asset: &Asset{
			Logo:          PtrString("string"),
			LogoImageLink: PtrString("string"),
			LogoLink:      PtrString("string"),
		},
		Controls: &Controls{
			EnableApi:      PtrBool(true),
			EnableControls: PtrBool(true),
			ForceAutoplay:  PtrBool(true),
			ForceLoop:      PtrBool(true),
			HideTitle:      PtrBool(true),
		},
		CreatedAt: PtrString("string"),
		Id:        PtrString("string"),
		Name:      PtrString("string"),
		Theme: &Theme{
			CaptionsBackground: PtrString("string"),
			ControlsBackground: PtrString("string"),
			MainColor:          PtrString("string"),
			MainColorActive:    PtrString("string"),
			MenuBackground:     PtrString("string"),
			TextControlColor:   PtrString("string"),
			TextControlHover:   PtrString("string"),
			TextMenuColor:      PtrString("string"),
			ThumbHeight:        PtrString("string"),
			TrackBackground:    PtrString("string"),
			TrackHeight:        PtrString("string"),
			TrackPlayed:        PtrString("string"),
			TrackUnplayed:      PtrString("string"),
			VideoBackground:    PtrString("string"),
		},
		UserId: PtrString("string"),
	},
	PlayerThemeId: PtrString("string"),
	Qualities: &[]QualityObject{
		{
			Name:   PtrString("string"),
			Status: PtrString("string"),
			Type:   PtrString("string"),
		},
	},
	Size:   PtrInt32(0),
	Status: PtrString("string"),
	Captions: &[]VideoCaption{
		{
			IsDefault: PtrBool(true),
			Language:  PtrString("string"),
			Url:       PtrString("string"),
		},
	},
	Tags: &[]string{
		"string",
	},
	Title:  PtrString("string"),
	UserId: PtrString("string"),
}

var getVideoListResponse = GetVideoListResponse{
	Data: &GetVideoListData{
		Total: PtrInt32(0),
		Videos: &[]Video{
			video,
		},
	},
	Status: PtrString("string"),
}

var createVideoResponse = CreateVideoResponse{
	Data:   &Video{},
	Status: PtrString("string"),
}

var getVideoTransCodingResponse = GetTranscodeCostResponse{
	Data: &GetTranscodeCostData{
		IsEnough: PtrBool(true),
		Price:    PtrString("0"),
	},
	Status: PtrString("string"),
}

var getVideoDetailResponse = GetVideoDetailResponse{
	Data:   &video,
	Status: PtrString("string"),
}

func TestVideos_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, getVideoListJSONResponse)
	})

	resp, err := client.Video.GetVideoList(GetVideoListRequest{})
	if err != nil {
		t.Errorf("Videos.VideosApiListRequest error: %v", err)
	}

	expected := &getVideoListResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.VideosApiListRequest\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestVideos_create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, getVideoJSONResponse)
	})

	resp, err := client.Video.Create(CreateVideoRequest{})
	if err != nil {
		t.Errorf("Videos.Create error: %v", err)
	}

	expected := &createVideoResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.Create\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestVideos_getCost(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/cost", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, getVideoTransCodingJSONResponse)
	})

	resp, err := client.Video.GetCost("360p", 120.0)
	if err != nil {
		t.Errorf("Videos.GetCost error: %v", err)
	}

	expected := &getVideoTransCodingResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.GetCost\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestVideos_getDetails(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/videoId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, getVideoJSONResponse)
	})

	resp, err := client.Video.GetDetail("videoId")
	if err != nil {
		t.Errorf("Videos.GetStatus error: %v", err)
	}

	expected := &getVideoDetailResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.GetStatus\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestVideos_delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/videoId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Video.Delete("videoId")
	if err != nil {
		t.Errorf("Videos.Delete error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.Delete\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestVideos_getComplete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/videoId/complete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Video.UploadVideoComplete("videoId")
	if err != nil {
		t.Errorf("Videos.GetComplete error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.GetComplete\n got=%#v\nwant=%#v", resp, expected)
	}
}
func TestVideos_updateInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/videoId/info", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Video.Update("videoId", UpdateVideoInfoRequest{})
	if err != nil {
		t.Errorf("Videos.UpdateInfo error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.UpdateInfo\n got=%#v\nwant=%#v", resp, expected)
	}
}
func TestVideos_uploadVideoThumbnail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/videoId/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, successResponse)
	})
	reader := strings.NewReader("data")
	resp, err := client.Video.UploadThumbnail("videoId", "thumbnail.jpg", reader)
	if err != nil {
		t.Errorf("Videos.UploadThumbnail error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Videos.UploadThumbnail\n got=%#v\nwant=%#v", resp, expected)
	}
}
