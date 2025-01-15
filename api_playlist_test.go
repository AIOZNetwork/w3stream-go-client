package w3streamsdk

import (
	"io"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testPlaylistID       string
	testVideoIDOne       = "98cdf5b1-991a-4b39-803f-509f580e90d5"
	testVideoIDTwo       = "0cacff1a-df3b-4867-b4f4-c4c4b45b7c25"
	testVideoIDThree     = "48250f3f-ba51-441b-8629-d6d3cc0ca192"
	testName             = "Test Playlist"
	testCurrentId        string
	testNextId           string
	testPreviousId       string
	deletePlaylistsLater []string
)

func TestPlaylistService_CreatePlaylist(t *testing.T) {
	tests := []struct {
		name    string
		request CreatePlaylistRequest
		wantErr bool
	}{
		{
			name: "Valid Create Request",
			request: CreatePlaylistRequest{
				Name: stringPtr(testName),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.CreatePlaylist(tt.request)
			testPlaylistID = *resp.Data.Playlist.Id
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				deletePlaylistsLater = append(deletePlaylistsLater, *resp.Data.Playlist.Id)
			}
		})
	}
}
func TestPlaylistService_GetPlaylists(t *testing.T) {
	tests := []struct {
		name    string
		request GetPlaylistListRequest
		wantErr bool
	}{
		{
			name:    "Valid Request",
			request: GetPlaylistListRequest{},
			wantErr: false,
		},
		{
			name: "Valid Request with Filter",
			request: GetPlaylistListRequest{
				Limit:   int32Ptr(10),
				Offset:  int32Ptr(0),
				SortBy:  stringPtr("created_at"),
				OrderBy: stringPtr("desc"),
			},
			wantErr: false,
		},
		{
			name: "Invalid SortBy",
			request: GetPlaylistListRequest{
				SortBy: stringPtr("invalid"),
			},
			wantErr: true,
		},
		{
			name: "Invalid OrderBy",
			request: GetPlaylistListRequest{
				OrderBy: stringPtr("invalid"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.GetPlaylists(tt.request)
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

func TestPlaylistService_UpdatePlaylist(t *testing.T) {
	notExistId := uuid.New().String()
	thumbnailFileForAnonymous, err := openTestImageFile(t)
	if err != nil {
		t.Fatal(err)
	}
	if thumbnailFileForAnonymous == nil {
		t.Fatal("thumbnailFileForAnonymous is nil")
	}
	defer thumbnailFileForAnonymous.Close()

	name := "Test Playlist"
	anonymousTest := []struct {
		name    string
		id      string
		Name    *string
		wantErr bool
		file    *os.File
	}{
		{
			name:    "Update other",
			id:      testPlaylistID,
			Name:    stringPtr(name),
			wantErr: true,
			file:    thumbnailFileForAnonymous,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader
			if tt.file != nil {
				reader = tt.file
			}
			resp, err := testAnonymousClient.Playlist.UpdatePlaylist(tt.id, nil, tt.Name, nil, "logo.png", reader)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	thumbnailFile, err := openTestImageFile(t)
	if err != nil {
		t.Fatal(err)
	}
	if thumbnailFile == nil {
		t.Fatal("thumbnailFile is nil")
	}
	invalidFile := openInvalidFile(t)
	defer invalidFile.Close()
	defer thumbnailFile.Close()

	tests := []struct {
		name    string
		id      string
		Name    *string
		wantErr bool
		file    *os.File
	}{
		{
			name:    "Valid Update Request With Name and Thumbnail",
			id:      testPlaylistID,
			Name:    stringPtr(name),
			file:    thumbnailFile,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
			file:    thumbnailFile,
		},
		{
			name:    "Invalid Thumbnail",
			id:      testPlaylistID,
			file:    invalidFile,
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			Name:    stringPtr(name),
			file:    thumbnailFile,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader
			if tt.file != nil {
				reader = tt.file
			}
			resp, err := testClient.Playlist.UpdatePlaylist(
				tt.id,
				nil,
				tt.Name,
				nil,
				"logo.png",
				reader,
			)
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

func TestPlaylistService_GetPlaylistPublicInfo(t *testing.T) {
	notExistId := uuid.New().String()
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Request",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
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
			resp, err := testClient.Playlist.GetPlaylistPublicInfo(tt.id)
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

func TestPlaylistService_AddVideoToPlaylist(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		request AddVideoToPlaylistRequest
		wantErr bool
	}{
		{
			name: "Add other",
			id:   testPlaylistID,
			request: AddVideoToPlaylistRequest{
				VideoId: stringPtr(testVideoIDOne),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.AddVideoToPlaylist(tt.id, tt.request)
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
		payload AddVideoToPlaylistRequest
		wantErr bool
	}{
		{
			name: "Valid Add First Video Request",
			id:   testPlaylistID,
			payload: AddVideoToPlaylistRequest{
				VideoId: stringPtr(testVideoIDOne),
			},
			wantErr: false,
		},
		{
			name: "Valid Add Second Video Request",
			id:   testPlaylistID,
			payload: AddVideoToPlaylistRequest{
				VideoId: stringPtr(testVideoIDTwo),
			},
			wantErr: false,
		},
		{
			name: "Valid Add Third Video Request",
			id:   testPlaylistID,
			payload: AddVideoToPlaylistRequest{
				VideoId: stringPtr(testVideoIDThree),
			},
			wantErr: false,
		},
		{
			name: "Missing Video ID",
			id:   testPlaylistID,
			payload: AddVideoToPlaylistRequest{
				VideoId: stringPtr(""),
			},
			wantErr: true,
		},
		{
			name:    "Empty Request",
			id:      testPlaylistID,
			payload: AddVideoToPlaylistRequest{},
			wantErr: true,
		},
		{
			name: "Not Exist ID",
			id:   notExistId,
			payload: AddVideoToPlaylistRequest{
				VideoId: stringPtr(testVideoIDOne),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.AddVideoToPlaylist(tt.id, tt.payload)
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

func TestPlaylistService_GetPlaylistByID(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testPlaylistID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.GetPlaylistById(tt.id, PlaylistApiGetPlaylistByIdRequest{})
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
		request PlaylistApiGetPlaylistByIdRequest
		wantErr bool
	}{
		{
			name:    "Valid Playlist ID",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name: "Valid Playlist ID with sortBy and orderBy",
			id:   testPlaylistID,
			request: PlaylistApiGetPlaylistByIdRequest{
				sortBy:  stringPtr("created_at"),
				orderBy: stringPtr("desc"),
			},
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Empty Request",
			id:      testPlaylistID,
			request: PlaylistApiGetPlaylistByIdRequest{},
			wantErr: false,
		},
		{
			name: "Invalid SortBy",
			id:   testPlaylistID,
			request: PlaylistApiGetPlaylistByIdRequest{
				sortBy: stringPtr("invalid"),
			},
			wantErr: true,
		},
		{
			name: "Invalid orderBy",
			id:   testPlaylistID,
			request: PlaylistApiGetPlaylistByIdRequest{
				orderBy: stringPtr("invalid"),
			},
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
			resp, err := testClient.Playlist.GetPlaylistById(tt.id, tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
				assert.NotNil(t, resp.Data.Playlist)

				videoItems := *resp.Data.Playlist.VideoItems
				assert.NotNil(t, videoItems)

				if len(videoItems) > 0 {
					firstVideoItem := videoItems[1]
					testCurrentId = *firstVideoItem.Id
					testNextId = *firstVideoItem.NextId
					testPreviousId = *firstVideoItem.PreviousId
					return
				}
			}
		})
	}

}

func TestPlaylistService_MoveVideoInPlaylist(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		payload MoveVideoInPlaylistRequest
		wantErr bool
	}{
		{
			name: "Move other",
			id:   testPlaylistID,
			payload: MoveVideoInPlaylistRequest{
				CurrentId:  stringPtr(testNextId),
				NextId:     stringPtr(testCurrentId),
				PreviousId: stringPtr(testPreviousId),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.MoveVideoInPlaylist(tt.id, tt.payload)
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
		payload MoveVideoInPlaylistRequest
		wantErr bool
	}{
		{
			name: "Valid Move Video Request",
			id:   testPlaylistID,
			payload: MoveVideoInPlaylistRequest{
				CurrentId:  stringPtr(testNextId),
				NextId:     stringPtr(testCurrentId),
				PreviousId: stringPtr(testPreviousId),
			},
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Missing Current ID",
			id:      testPlaylistID,
			wantErr: true,
		},
		{
			name: "Missing Next ID",
			id:   testPlaylistID,
			payload: MoveVideoInPlaylistRequest{
				CurrentId: stringPtr(testCurrentId),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.MoveVideoInPlaylist(tt.id, tt.payload)
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

func TestPlaylistService_RemoveVideoFromPlaylist(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		itemId  string
		wantErr bool
	}{
		{
			name:    "Remove other",
			id:      testPlaylistID,
			itemId:  testPreviousId,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.RemoveVideoFromPlaylist(tt.id, tt.itemId)
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
		itemId  string
		wantErr bool
	}{
		{
			name:    "Valid Remove First Video Request",
			id:      testPlaylistID,
			itemId:  testPreviousId,
			wantErr: false,
		},
		{
			name:    "Valid Remove Second Video Request",
			id:      testPlaylistID,
			itemId:  testCurrentId,
			wantErr: false,
		},
		{
			name:    "Valid Remove Third Video Request",
			id:      testPlaylistID,
			itemId:  testNextId,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Missing Item ID",
			id:      testPlaylistID,
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
			resp, err := testClient.Playlist.RemoveVideoFromPlaylist(tt.id, tt.itemId)
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

func TestPlaylistService_DeletePlaylistThumbnail(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testPlaylistID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.DeletePlaylistThumbnail(tt.id)
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
			name:    "Valid Playlist ID",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
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
			resp, err := testClient.Playlist.DeletePlaylistThumbnail(tt.id)
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

func TestPlaylistService_DeletePlaylistById(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testPlaylistID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.DeletePlaylistById(tt.id)
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
			name:    "Valid Playlist ID",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
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
			resp, err := testClient.Playlist.DeletePlaylistById(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	for _, id := range deletePlaylistsLater {
		testClient.Playlist.DeletePlaylistById(id)
	}

}

func int32Ptr(i int32) *int32 {
	return &i
}
