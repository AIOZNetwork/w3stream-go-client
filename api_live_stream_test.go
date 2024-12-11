package w3streamsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	liveStreamKeyName = "name"
	liveStreamKeyID   string
	streamId          string
	streamingTitle    = "title"
	qualities         = []string{
		"1080p",
		"720p",
		"360p",
	}
)

func TestLiveStreamService_CreateLiveStreamKey(t *testing.T) {
	tests := []struct {
		name    string
		input   CreateLiveStreamKeyRequest
		wantErr bool
	}{
		{
			name: "Valid Create Live Stream Key",
			input: CreateLiveStreamKeyRequest{
				Name: &liveStreamKeyName,
				Save: boolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "Invalid Create Live Stream Key with Name is nil",
			input: CreateLiveStreamKeyRequest{
				Name: nil,
			},
			wantErr: true,
		},
		{
			name: "Valid Create Live Stream Key with Save is nil",
			input: CreateLiveStreamKeyRequest{
				Name: &liveStreamKeyName,
				Save: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.CreateLiveStreamKey(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				if resp != nil {
					liveStreamKeyID = *resp.Data.Id
				}
			}
		})
	}
}

func TestLiveStreamService_GetLiveStreamKey(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      liveStreamKeyID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.LiveStream.GetLiveStreamKey(tt.id)
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
			name:    "Valid Get Live Stream Key",
			id:      liveStreamKeyID,
			wantErr: false,
		},
		{
			name:    "Invalid Get Live Stream Key",
			id:      "invalid-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.GetLiveStreamKey(tt.id)
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

func TestLiveStreamService_GetLiveStreamKeys(t *testing.T) {
	tests := []struct {
		name    string
		request LiveStreamApiGetLiveStreamKeysRequest
		wantErr bool
	}{
		{
			name:    "Valid Get Live Stream Keys with Limit, Offset, OrderBy, SortBy",
			request: LiveStreamApiGetLiveStreamKeysRequest{}.Limit(10).Offset(0).SortBy("created_at").OrderBy("desc"),
			wantErr: false,
		},
		{
			name:    "Valid Get Live Stream Keys with Search",
			request: LiveStreamApiGetLiveStreamKeysRequest{}.Search(liveStreamKeyName),
			wantErr: false,
		},
		{
			name:    "Valid Get Live Stream Keys with no fields",
			request: LiveStreamApiGetLiveStreamKeysRequest{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.GetLiveStreamKeys(tt.request)
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

func TestLiveStreamService_UpdateLiveStreamKey(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		input   UpdateLiveStreamKeyRequest
		wantErr bool
	}{
		{
			name: "Update other",
			id:   liveStreamKeyID,
			input: UpdateLiveStreamKeyRequest{
				Name: &liveStreamKeyName,
				Save: boolPtr(true),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.LiveStream.UpdateLiveStreamKey(tt.id, tt.input)
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
		input   UpdateLiveStreamKeyRequest
		wantErr bool
	}{
		{
			name: "Valid Update Live Stream Key",
			id:   liveStreamKeyID,
			input: UpdateLiveStreamKeyRequest{
				Name: &liveStreamKeyName,
				Save: boolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "Update Live Stream Key with Name is nil",
			id:   liveStreamKeyID,
			input: UpdateLiveStreamKeyRequest{
				Name: nil,
				Save: boolPtr(true),
			},
			wantErr: false,
		},
		{
			name: "Update Live Stream Key with Save is nil",
			id:   liveStreamKeyID,
			input: UpdateLiveStreamKeyRequest{
				Name: &liveStreamKeyName,
				Save: nil,
			},
			wantErr: false,
		},
		{
			name: "Invalid Update Live Stream Key with ID is empty",
			id:   "",
			input: UpdateLiveStreamKeyRequest{
				Name: &liveStreamKeyName,
				Save: boolPtr(true),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.UpdateLiveStreamKey(tt.id, tt.input)
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

func TestLiveStreamService_GetLiveStreamVideo(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		request GetLiveStreamVideosRequest
		wantErr bool
	}{
		{
			name: "Valid Get Live Stream Videos",
			id:   liveStreamKeyID,
			request: GetLiveStreamVideosRequest{
				Limit:   int32Ptr(10),
				Offset:  int32Ptr(0),
				SortBy:  stringPtr("created_at"),
				OrderBy: stringPtr("desc"),
			},
			wantErr: false,
		},
		{
			name:    "Invalid Get Live Stream Videos",
			id:      "invalid-id",
			request: GetLiveStreamVideosRequest{},
			wantErr: true,
		},
		{
			name: "Invalid Get Live Stream Videos with Limit is -1",
			id:   liveStreamKeyID,
			request: GetLiveStreamVideosRequest{
				Limit: int32Ptr(-1),
			},
			wantErr: true,
		},
		{
			name: "Invalid Get Live Stream Videos with Offset is -1",
			id:   liveStreamKeyID,
			request: GetLiveStreamVideosRequest{
				Offset: int32Ptr(-1),
			},
			wantErr: true,
		},
		{
			name: "Invalid Get Live Stream Videos with SortBy is empty",
			id:   liveStreamKeyID,
			request: GetLiveStreamVideosRequest{
				SortBy: stringPtr(""),
			},
			wantErr: true,
		},
		{
			name: "Invalid Get Live Stream Videos with OrderBy is empty",
			id:   liveStreamKeyID,
			request: GetLiveStreamVideosRequest{
				OrderBy: stringPtr(""),
			},
			wantErr: true,
		},
		{
			name: "Invalid Get Live Stream Videos with Search is empty",
			id:   liveStreamKeyID,
			request: GetLiveStreamVideosRequest{
				Search: stringPtr(""),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		resp, err := testClient.LiveStream.GetLiveStreamVideos(tt.id, tt.request)
		t.Run(tt.name, func(t *testing.T) {
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
func TestLiveStreamService_CreateStreaming(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		input   CreateStreamingRequest
		wantErr bool
	}{
		{
			name: "Valid Create Streaming",
			id:   liveStreamKeyID,
			input: CreateStreamingRequest{
				Title:     &streamingTitle,
				Qualities: &qualities,
				Save:      boolPtr(false),
			},
			wantErr: false,
		},
		{
			name:    "Invalid Create Streaming with Invalid ID",
			id:      "invalid-id",
			input:   CreateStreamingRequest{},
			wantErr: true,
		},
		{
			name: "Invalid Create Streaming with Title is nil",
			id:   liveStreamKeyID,
			input: CreateStreamingRequest{
				Title:     nil,
				Qualities: &qualities,
				Save:      boolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "Invalid Create Streaming with Qualities is nil",
			id:   liveStreamKeyID,
			input: CreateStreamingRequest{
				Title:     &streamingTitle,
				Qualities: nil,
				Save:      boolPtr(true),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.CreateStreaming(tt.id, tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				if resp != nil {
					streamId = *resp.Data.Id
				}
			}
		})
	}
}

func TestLiveStreamService_GetStreaming(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		streamId string
		wantErr  bool
	}{
		{
			name:     "Valid Get Streaming",
			id:       liveStreamKeyID,
			streamId: streamId,
			wantErr:  false,
		},
		{
			name:     "Invalid Get Streaming",
			id:       "invalid-id",
			streamId: "invalid-stream-id",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.GetStreaming(tt.id, tt.streamId)
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

func TestLiveStreamService_GetStreamings(t *testing.T) {
	tests := []struct {
		id      string
		name    string
		wantErr bool
	}{
		{
			name:    "Valid Get Streamings",
			id:      liveStreamKeyID,
			wantErr: false,
		},
		{
			name:    "Invalid Get Streamings",
			id:      "invalid-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.GetStreamings(tt.id)
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

func TestLiveStreamService_DeleteLiveStreamVideo(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      streamId,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.LiveStream.DeleteLiveStreamVideo(tt.id)
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
			name:    "Valid Delete Live Stream Video",
			id:      streamId,
			wantErr: false,
		},
		{
			name:    "Invalid Delete Live Stream Video",
			id:      "invalid-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.DeleteLiveStreamVideo(tt.id)
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

func TestLiveStreamService_DeleteLiveStreamKey(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      liveStreamKeyID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.LiveStream.DeleteLiveStreamKey(tt.id)
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
			name:    "Valid Delete Live Stream Key with ID",
			id:      liveStreamKeyID,
			wantErr: false,
		},
		{
			name:    "Invalid Delete Live Stream Key with Invalid ID",
			id:      "invalid-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.LiveStream.DeleteLiveStreamKey(tt.id)
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
