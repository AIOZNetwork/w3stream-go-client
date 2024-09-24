package w3streamsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var playerTheresJSONResponse = `{
	"data": {
    "player_themes": [
      {
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
      }
    ],
    "total": 0
  },
  "status": "string"
}`

var playerThemeJSONResponse = `{
	"data": {
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
    }
  },
  "status": "string"
}`

var playTheme = PlayerTheme{
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
}

var createPlayerThemeResponse = CreatePlayerThemesResponse{
	Data: &CreatePlayerThemesData{
		PlayerTheme: &playTheme,
	},
	Status: PtrString("string"),
}

var getPlayerThemeResponse = GetPlayerThemeByIdResponse{
	Data: &GetPlayerThemeByIdData{
		PlayerTheme: &playTheme,
	},
	Status: PtrString("string"),
}

var updatePlayerThemeResponse = UpdatePlayerThemeResponse{
	Data: &GetPlayerThemeByIdData{
		PlayerTheme: &playTheme,
	},
	Status: PtrString("string"),
}

var getPlayerThemesResponse = GetPlayerThemeResponse{
	Data: &GetPlayerThemeData{
		PlayerThemes: &[]PlayerTheme{
			playTheme,
		},
		Total: PtrInt32(0),
	},
	Status: PtrString("string"),
}

func TestPlayerTheme_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, playerThemeJSONResponse)
	})

	resp, err := client.Players.Create(CreatePlayerThemeRequest{})
	if err != nil {
		t.Errorf("PlayerTheme.Create error: %v", err)
	}

	expected := &createPlayerThemeResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.Create\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestPlayerTheme_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, playerTheresJSONResponse)
	})

	resp, err := client.Players.List(PlayersApiListRequest{})
	if err != nil {
		t.Errorf("PlayerTheme.List error: %v", err)
	}

	expected := &getPlayerThemesResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.List\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestPlayerTheme_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/playerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, playerThemeJSONResponse)
	})

	resp, err := client.Players.Get("playerId")
	if err != nil {
		t.Errorf("PlayerTheme.Get error: %v", err)
	}

	expected := &getPlayerThemeResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.Get\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestPlayerTheme_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/playerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Players.Delete("playerId")
	if err != nil {
		t.Errorf("PlayerTheme.Delete error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.Delete\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestPlayerTheme_Update(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/playerId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, playerThemeJSONResponse)
	})

	resp, err := client.Players.Update("playerId", UpdatePlayerThemeRequest{})
	if err != nil {
		t.Errorf("PlayerTheme.Update error: %v", err)
	}

	expected := &updatePlayerThemeResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.Update\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestPlayerTheme_UploadLogo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/playerId/logo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, playerThemeJSONResponse)
	})

	reader := strings.NewReader("data")
	resp, err := client.Players.UploadLogo("playerId", "link", "photo.png", reader)
	if err != nil {
		t.Errorf("PlayerTheme.UploadLogo error: %v", err)
	}

	expected := &getPlayerThemeResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.UploadLogo\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestPlayerTheme_DeleteLogo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/playerId/logo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Players.DeleteLogo("playerId")
	if err != nil {
		t.Errorf("PlayerTheme.DeleteLogo error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.DeleteLogo\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestPlayerTheme_AddPlayer(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/add-player", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Players.AddPlayer(AddPlayerThemesToVideoRequest{})
	if err != nil {
		t.Errorf("PlayerTheme.AddPlayer error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("PlayerTheme.AddPlayer\n got=%#v\nwant=%#v", resp, expected)
	}
}
