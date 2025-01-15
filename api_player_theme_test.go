package w3streamsdk

import (
	"io"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testPlayerIDForUpdateAndDeleteAndGet string
	playerName                           = "Test Player Theme"
	logoURL                              = "https://example.com/logo.png"
	testVideoForPlayer                   = "1f625e54-308d-401e-aa67-12f86eff6d1a"
	deletePlayerThemesLater              []string
)

func openTestImageFile(t *testing.T) (*os.File, error) {
	logoFile, err := os.Open("test-assets/logo.png")
	if err != nil {
		t.Fatal(err)
	}
	return logoFile, nil
}

func openInvalidFile(t *testing.T) *os.File {
	logoFile, err := os.Open("test-assets/invalid-file.txt")
	if err != nil {
		t.Fatal(err)
	}
	return logoFile
}

func TestPlayersService_Create(t *testing.T) {
	tests := []struct {
		name    string
		request CreatePlayerThemeRequest
		wantErr bool
	}{
		{
			name: "Valid Create with All Fields",
			request: CreatePlayerThemeRequest{
				Name:      stringPtr(playerName),
				IsDefault: boolPtr(true),
				Controls: &Controls{
					EnableApi:      boolPtr(true),
					EnableControls: boolPtr(true),
					ForceAutoplay:  boolPtr(false),
					ForceLoop:      boolPtr(false),
					HideTitle:      boolPtr(false),
				},
				Theme: &Theme{
					ControlBarBackgroundColor: stringPtr("#ffffff"),
					ControlBarHeight:          stringPtr("100px"),
					MainColor:                 stringPtr("#ffffff"),
					MenuBackgroundColor:       stringPtr("#ffffff"),
					MenuItemBackgroundHover:   stringPtr("#cccccc"),
					ProgressBarCircleSize:     stringPtr("10px"),
					ProgressBarHeight:         stringPtr("10px"),
					TextColor:                 stringPtr("#ffffff"),
					TextTrackBackground:       stringPtr("#ffffff"),
					TextTrackColor:            stringPtr("#ffffff"),
				},
			},
			wantErr: false,
		},
		{
			name: "Valid Create with Some Required Fields Only",
			request: CreatePlayerThemeRequest{
				Name:      stringPtr(playerName),
				IsDefault: boolPtr(true),
				Theme: &Theme{
					ControlBarBackgroundColor: stringPtr("#ffffff"),
					ControlBarHeight:          stringPtr("100px"),
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Color Code",
			request: CreatePlayerThemeRequest{
				Name: stringPtr(playerName),
				Theme: &Theme{
					TextColor: stringPtr("invalid-color"),
				},
			},
			wantErr: true,
		},
		{
			name: "Empty Name",
			request: CreatePlayerThemeRequest{
				Theme: &Theme{
					TextColor: stringPtr("#ffffff"),
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Size Format",
			request: CreatePlayerThemeRequest{
				Name: stringPtr(playerName),
				Theme: &Theme{
					ControlBarHeight: stringPtr("invalid-size"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.Create(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data.PlayerTheme.Id)
				testPlayerIDForUpdateAndDeleteAndGet = *resp.Data.PlayerTheme.Id
				deletePlayerThemesLater = append(deletePlayerThemesLater, *resp.Data.PlayerTheme.Id)
			}
		})
	}
}

func TestPlayersService_UploadLogo(t *testing.T) {
	logoFile, err := openTestImageFile(t)
	if err != nil {
		t.Fatal(err)
	}
	invalidFile := openInvalidFile(t)
	defer logoFile.Close()
	defer invalidFile.Close()

	tests := []struct {
		name     string
		playerID string
		file     *os.File
		link     string
		wantErr  bool
	}{
		{
			name:     "Valid Upload",
			playerID: testPlayerIDForUpdateAndDeleteAndGet,
			file:     logoFile,
			link:     logoURL,
			wantErr:  false,
		},
		{
			name:     "Invalid Player ID",
			playerID: "invalid-id",
			file:     logoFile,
			link:     logoURL,
			wantErr:  true,
		},
		{
			name:     "Invalid File Type",
			playerID: testPlayerIDForUpdateAndDeleteAndGet,
			file:     invalidFile,
			link:     logoURL,
			wantErr:  true,
		},
		{
			name:     "Invalid Upload Link",
			playerID: testPlayerIDForUpdateAndDeleteAndGet,
			file:     logoFile,
			link:     "invalid-link",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader
			if tt.file != nil {
				reader = tt.file
			}
			resp, err := testClient.Players.UploadLogo(tt.playerID, tt.link, "logo.png", reader)
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

func TestPlayersService_AddPlayer(t *testing.T) {
	notExistId := uuid.New().String()
	tests := []struct {
		name    string
		request AddPlayerThemesToVideoRequest
		wantErr bool
	}{
		{
			name: "Valid Add",
			request: AddPlayerThemesToVideoRequest{
				PlayerThemeId: stringPtr(testPlayerIDForUpdateAndDeleteAndGet),
				VideoId:       stringPtr(testVideoForPlayer),
			},
			wantErr: false,
		},
		{
			name: "Invalid Player ID",
			request: AddPlayerThemesToVideoRequest{
				PlayerThemeId: stringPtr("invalid-id"),
				VideoId:       stringPtr(testVideoForPlayer),
			},
			wantErr: true,
		},
		{
			name: "Empty Video ID",
			request: AddPlayerThemesToVideoRequest{
				PlayerThemeId: stringPtr(testPlayerIDForUpdateAndDeleteAndGet),
			},
			wantErr: true,
		},
		{
			name: "Not Exist ID",
			request: AddPlayerThemesToVideoRequest{
				PlayerThemeId: stringPtr(notExistId),
				VideoId:       stringPtr(testVideoForPlayer),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.AddPlayer(tt.request)
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

func TestPlayersService_Get(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testPlayerIDForUpdateAndDeleteAndGet,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Players.Get(tt.id)
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
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Get",
			id:      testPlayerIDForUpdateAndDeleteAndGet,
			wantErr: false,
		},
		{
			name:    "Invalid ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.Get(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data.PlayerTheme.Id)
			}
		})
	}
}

func TestPlayersService_Update(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		request UpdatePlayerThemeRequest
		wantErr bool
	}{
		{
			name: "Update other",
			id:   testPlayerIDForUpdateAndDeleteAndGet,
			request: UpdatePlayerThemeRequest{
				Name: stringPtr("Updated Player Theme"),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Players.Update(tt.id, tt.request)
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
		id      string
		request UpdatePlayerThemeRequest
		wantErr bool
	}{
		{
			name: "Valid Update",
			id:   testPlayerIDForUpdateAndDeleteAndGet,
			request: UpdatePlayerThemeRequest{
				Name: stringPtr("Updated Player Theme"),
				Theme: &Theme{
					TextColor: stringPtr("#000000"),
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Color Code",
			id:   testPlayerIDForUpdateAndDeleteAndGet,
			request: UpdatePlayerThemeRequest{
				Theme: &Theme{
					TextColor: stringPtr("invalid-color"),
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			id:   "invalid-id",
			request: UpdatePlayerThemeRequest{
				Name: stringPtr("Updated Player Theme"),
			},
			wantErr: true,
		},
		{
			name: "Not Exist ID",
			id:   notExistId,
			request: UpdatePlayerThemeRequest{
				Name: stringPtr("Updated Player Theme"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.Update(tt.id, tt.request)
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

func TestPlayersService_List(t *testing.T) {
	tests := []struct {
		name    string
		request PlayersApiListRequest
		wantErr bool
	}{
		{
			name:    "List All",
			request: PlayersApiListRequest{},
			wantErr: false,
		},
		{
			name: "With Search and Pagination",
			request: PlayersApiListRequest{}.
				Search("test").
				Limit(10).
				Offset(0).
				SortBy("created_at").
				OrderBy("desc"),
			wantErr: false,
		},
		{
			name: "Invalid Offset",
			request: PlayersApiListRequest{}.
				Offset(-1),
			wantErr: true,
		},
		{
			name: "Invalid Limit",
			request: PlayersApiListRequest{}.
				Limit(1001),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.List(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
			}
		})
	}
}

func TestPlayersService_DeleteLogo(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testPlayerIDForUpdateAndDeleteAndGet,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Players.DeleteLogo(tt.id)
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
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Delete Logo",
			id:      testPlayerIDForUpdateAndDeleteAndGet,
			wantErr: false,
		},
		{
			name:    "Invalid ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.DeleteLogo(tt.id)
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

func TestPlayersService_RemovePlayer(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		request RemovePlayerThemesFromVideoRequest
		wantErr bool
	}{
		{
			name: "Remove other",
			request: RemovePlayerThemesFromVideoRequest{
				PlayerThemeId: stringPtr(testPlayerIDForUpdateAndDeleteAndGet),
				VideoId:       stringPtr(testVideoForPlayer),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Players.RemovePlayer(tt.request)
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
		request RemovePlayerThemesFromVideoRequest
		wantErr bool
	}{
		{
			name: "Valid Remove",
			request: RemovePlayerThemesFromVideoRequest{
				PlayerThemeId: stringPtr(testPlayerIDForUpdateAndDeleteAndGet),
				VideoId:       stringPtr(testVideoForPlayer),
			},
			wantErr: false,
		},
		{
			name: "Invalid Player ID",
			request: RemovePlayerThemesFromVideoRequest{
				PlayerThemeId: stringPtr("invalid-id"),
				VideoId:       stringPtr(testVideoForPlayer),
			},
			wantErr: true,
		},
		{
			name: "Empty Video ID",
			request: RemovePlayerThemesFromVideoRequest{
				PlayerThemeId: stringPtr(testPlayerIDForUpdateAndDeleteAndGet),
			},
			wantErr: true,
		},
		{
			name: "Not Exist ID",
			request: RemovePlayerThemesFromVideoRequest{
				PlayerThemeId: stringPtr(notExistId),
				VideoId:       stringPtr(testVideoForPlayer),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.RemovePlayer(tt.request)
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
func TestPlayersService_Delete(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testPlayerIDForUpdateAndDeleteAndGet,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Players.Delete(tt.id)
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
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Delete",
			id:      testPlayerIDForUpdateAndDeleteAndGet,
			wantErr: false,
		},
		{
			name:    "Invalid ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Players.Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	for _, id := range deletePlayerThemesLater {
		testClient.Players.Delete(id)
	}
}
