package w3streamsdk

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	testPublicKey          string
	testSecretKey          string
	testAnonymousSecretKey string
	testAnonymousPublicKey string

	testClient          *Client
	testAnonymousClient *Client

	testVideoID        string
	title              = "Test Video"
	description        = "Test Description"
	testVideoCaptionID = "c91e1c8b-e93c-423c-98dd-690fdfa19659"
)

func init() {
	if err := loadEnvVariables(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	initializeClients()
}

func loadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	var missingVars []string

	testPublicKey = os.Getenv("TEST_PUBLIC_KEY")
	if testPublicKey == "" {
		missingVars = append(missingVars, "TEST_PUBLIC_KEY")
	}

	testSecretKey = os.Getenv("TEST_SECRET_KEY")
	if testSecretKey == "" {
		missingVars = append(missingVars, "TEST_SECRET_KEY")
	}

	testAnonymousSecretKey = os.Getenv("TEST_ANONYMOUS_SECRET_KEY")
	if testAnonymousSecretKey == "" {
		missingVars = append(missingVars, "TEST_ANONYMOUS_SECRET_KEY")
	}

	testAnonymousPublicKey = os.Getenv("TEST_ANONYMOUS_PUBLIC_KEY")

	if len(missingVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return nil
}

func initializeClients() {
	testClient = ClientBuilder(AuthCredentials{
		PublicKey: testPublicKey,
		SecretKey: testSecretKey,
	}).Build()

	testAnonymousClient = ClientBuilder(AuthCredentials{
		PublicKey: testAnonymousPublicKey,
		SecretKey: testAnonymousSecretKey,
	}).Build()
}

func openTestVideoFile(t *testing.T) *os.File {
	file, err := os.Open("test-assets/558k.mp4")
	if err != nil {
		t.Fatal(err)
	}
	return file
}

func getFileHash(t *testing.T, file *os.File) string {
	hash := md5.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func TestVideoService_Create(t *testing.T) {
	validMetadata := []Metadata{
		{Key: stringPtr("key1"), Value: stringPtr("value1")},
		{Key: stringPtr("key2"), Value: stringPtr("value2")},
	}

	validTags := []string{"tag1", "tag2"}
	tests := []struct {
		name    string
		request CreateVideoRequest
		wantErr bool
	}{
		{
			name: "Valid Complete Request",
			request: CreateVideoRequest{
				Title:       stringPtr("Test Video"),
				Description: stringPtr("Test Description"),
				IsPublic:    boolPtr(true),
				Metadata:    &validMetadata,
				Qualities:   &[]string{"1080p", "720p", "360p"},
				Tags:        &validTags,
			},
			wantErr: false,
		},
		{
			name: "Valid Minimal Request",
			request: CreateVideoRequest{
				Title:     stringPtr("Test Video"),
				Qualities: &[]string{"720p"},
			},
			wantErr: false,
		},
		{
			name: "Invalid Title - Empty",
			request: CreateVideoRequest{
				Qualities: &[]string{"720p"},
			},
			wantErr: true,
		},
		{
			name: "Invalid Title - Too Long",
			request: CreateVideoRequest{
				Title:     stringPtr(strings.Repeat("a", 256)),
				Qualities: &[]string{"720p"},
			},
			wantErr: true,
		},
		{
			name: "Invalid Description - Too Long",
			request: CreateVideoRequest{
				Title:       stringPtr("Test Video"),
				Description: stringPtr(strings.Repeat("a", 1001)),
				Qualities:   &[]string{"720p"},
			},
			wantErr: true,
		},
		{
			name: "Invalid Metadata - Key Too Long",
			request: CreateVideoRequest{
				Title:     stringPtr("Test Video"),
				Qualities: &[]string{"720p"},
				Metadata:  &[]Metadata{{Key: stringPtr(strings.Repeat("a", 256)), Value: stringPtr("value")}},
			},
			wantErr: true,
		},
		{
			name: "Invalid Qualities - Empty",
			request: CreateVideoRequest{
				Title:     stringPtr("Test Video"),
				Qualities: &[]string{},
			},
			wantErr: false,
		},
		{
			name: "Invalid Qualities - Invalid Value",
			request: CreateVideoRequest{
				Title:     stringPtr("Test Video"),
				Qualities: &[]string{"invalid_quality"},
			},
			wantErr: true,
		},
		{
			name: "Valid Qualities - All Supported",
			request: CreateVideoRequest{
				Title:     stringPtr("Test Video"),
				Qualities: &[]string{"2160p", "1440p", "1080p", "720p", "360p", "240p", "144p"},
			},
			wantErr: false,
		},
		{
			name: "Invalid Tags - Tag Too Long",
			request: CreateVideoRequest{
				Title:     stringPtr("Test Video"),
				Qualities: &[]string{"720p"},
				Tags:      &[]string{strings.Repeat("a", 256)},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.Create(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data.Id)
				testVideoID = *resp.Data.Id
			}
		})
	}

}

func TestVideoService_List(t *testing.T) {
	tests := []struct {
		name    string
		request GetVideoListRequest
		wantErr bool
	}{
		{
			name:    "Valid Get Video List With No Filter",
			request: GetVideoListRequest{},
			wantErr: false,
		},
		{
			name: "Valid Get Video List With Filter",
			request: GetVideoListRequest{
				Limit:   int32Ptr(10),
				Offset:  int32Ptr(0),
				OrderBy: stringPtr("created_at"),
				SortBy:  stringPtr("desc"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.GetVideoList(tt.request)
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

func TestVideoService_Update(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		input   UpdateVideoInfoRequest
		wantErr bool
	}{
		{
			name:    "Update other",
			id:      testVideoID,
			input:   UpdateVideoInfoRequest{Title: &title},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Video.Update(tt.id, tt.input)
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
		input   UpdateVideoInfoRequest
		wantErr bool
	}{
		{
			name: "Valid Update All Fields",
			id:   testVideoID,
			input: UpdateVideoInfoRequest{
				Title:       &title,
				Description: &description,
			},
			wantErr: false,
		},
		{
			name: "Valid Update Title Only",
			id:   testVideoID,
			input: UpdateVideoInfoRequest{
				Title: &title,
			},
			wantErr: false,
		},
		{
			name: "Valid Update Description Only",
			id:   testVideoID,
			input: UpdateVideoInfoRequest{
				Description: &description,
			},
			wantErr: false,
		},
		{
			name: "Invalid Title Length",
			id:   testVideoID,
			input: UpdateVideoInfoRequest{
				Title: stringPtr(strings.Repeat("a", 256)),
			},
			wantErr: true,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-uuid",
			input:   UpdateVideoInfoRequest{Title: &title},
			wantErr: true,
		},
		{
			name:    "Non-existent Video ID",
			id:      "12345678-1234-1234-1234-123456789012",
			input:   UpdateVideoInfoRequest{Title: &title},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.Update(tt.id, tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestVideoService_GetDetail(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testVideoID,
			wantErr: true,
		},
	}
	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Video.GetDetail(tt.id)
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
		checkFn func(*testing.T, *GetVideoDetailResponse)
	}{
		{
			name:    "Valid Get Detail",
			id:      testVideoID,
			wantErr: false,
			checkFn: func(t *testing.T, resp *GetVideoDetailResponse) {
				assert.NotEmpty(t, resp.Data.Id)
				assert.NotEmpty(t, resp.Data.Title)
				assert.NotEmpty(t, resp.Status)
			},
		},
		{
			name:    "Invalid ID Format",
			id:      "invalid-uuid",
			wantErr: true,
			checkFn: nil,
		},
		{
			name:    "Non-existent ID",
			id:      "12345678-1234-1234-1234-123456789012",
			wantErr: true,
			checkFn: nil,
		},
		{
			name:    "Empty ID",
			id:      "",
			wantErr: true,
			checkFn: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.GetDetail(tt.id)
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

func TestVideoService_UploadThumbnail(t *testing.T) {
	validThumbnail, err := openTestImageFile(t)
	if err != nil {
		t.Fatal("Failed to open test image file:", err)
	}
	if validThumbnail == nil {
		t.Fatal("Valid thumbnail file is nil")
	}
	defer validThumbnail.Close()

	anotherThumbnail, err := openTestImageFile(t)
	if err != nil {
		t.Fatal("Failed to open test image file:", err)
	}

	if anotherThumbnail == nil {
		t.Fatal("anotherThumbnail is nil")
	}
	defer anotherThumbnail.Close()

	invalidFile := openInvalidFile(t)
	if invalidFile == nil {
		t.Fatal("Invalid file is nil")
	}
	defer invalidFile.Close()

	t.Run("Anonymous Client Tests", func(t *testing.T) {
		tests := []struct {
			name     string
			id       string
			fileName string
			file     *os.File
			wantErr  bool
		}{
			{
				name:     "Unauthorized thumbnail upload",
				id:       testVideoID,
				fileName: "thumbnail.jpg",
				file:     anotherThumbnail,
				wantErr:  true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				resp, err := testAnonymousClient.Video.UploadThumbnail(tt.id, tt.fileName, tt.file)
				assert.Equal(t, tt.wantErr, err != nil)
				if tt.wantErr {
					assert.Nil(t, resp)
				}
			})
		}
	})

	tests := []struct {
		name     string
		id       string
		fileName string
		file     *os.File
		wantErr  bool
	}{
		{
			name:     "Valid Thumbnail Upload",
			id:       testVideoID,
			fileName: "thumbnail.jpg",
			file:     validThumbnail,
			wantErr:  false,
		},
		{
			name:     "Invalid File Type",
			id:       testVideoID,
			fileName: "thumbnail.gif",
			file:     invalidFile,
			wantErr:  true,
		},
		{
			name:     "Invalid Video ID",
			id:       "invalid-id",
			fileName: "thumbnail.jpg",
			file:     validThumbnail,
			wantErr:  true,
		},
		{
			name:     "Empty File Name",
			id:       testVideoID,
			fileName: "",
			file:     validThumbnail,
			wantErr:  true,
		},
		{
			name:     "Nil File",
			id:       testVideoID,
			fileName: "thumbnail.jpg",
			file:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.UploadThumbnail(tt.id, tt.fileName, tt.file)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestVideoService_GetCost(t *testing.T) {
	tests := []struct {
		name      string
		qualities string
		duration  float32
		wantErr   bool
	}{
		{
			name:      "Valid Single Quality",
			qualities: "720p",
			duration:  120.5,
			wantErr:   false,
		},
		{
			name:      "Valid Multiple Qualities",
			qualities: "720p,1080p",
			duration:  120.5,
			wantErr:   false,
		},
		{
			name:      "Invalid Quality",
			qualities: "invalid",
			duration:  120.5,
			wantErr:   true,
		},
		{
			name:      "Empty Quality",
			qualities: "",
			duration:  120.5,
			wantErr:   true,
		},
		{
			name:      "Negative Duration",
			qualities: "720p",
			duration:  -1,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.GetCost(tt.qualities, tt.duration)
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

func TestVideoService_UploadPart(t *testing.T) {

	video := openTestVideoFile(t)
	videoHash := getFileHash(t, video)

	index := "1"

	tests := []struct {
		name    string
		id      string
		hash    *string
		index   *string
		file    *os.File
		wantErr bool
	}{
		{
			name:    "Valid File Upload",
			id:      testVideoID,
			hash:    &videoHash,
			index:   &index,
			file:    video,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			hash:    &videoHash,
			index:   &index,
			file:    video,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInfo, _ := tt.file.Stat()
			resp, err := testClient.Video.UploadPart(tt.id, tt.hash, tt.index, tt.file.Name(), tt.file, fileInfo.Size())
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestVideoService_UploadWhenVideoComplete(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Get Upload When Video Complete",
			id:      testVideoID,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty Video ID",
			id:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.UploadVideoComplete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Check other video complete",
			id:      testVideoID,
			wantErr: true,
		},
	}
	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Video.UploadVideoComplete(tt.id)
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

func TestVideoService_GetVideoPlayerInfo(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		request VideoApiGetVideoPlayerInfoRequest
		wantErr bool
	}{
		{
			name:    "Valid Get Video Player Info",
			id:      testVideoID,
			request: VideoApiGetVideoPlayerInfoRequest{},
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			request: VideoApiGetVideoPlayerInfoRequest{},
			wantErr: true,
		},
		{
			name:    "Empty Video ID",
			id:      "",
			request: VideoApiGetVideoPlayerInfoRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.GetVideoPlayerInfo(tt.id, tt.request)
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

func TestVideoService_CreateVideoCaptions(t *testing.T) {
	tmpFile := createTempVTTFile(t)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tests := []struct {
		name    string
		id      string
		lang    string
		file    *os.File
		wantErr bool
	}{
		{
			name:    "Valid Create Video Captions",
			id:      testVideoID,
			lang:    testLang,
			file:    tmpFile,
			wantErr: false,
		},
		{
			name:    "Empty Language",
			id:      testVideoID,
			lang:    "",
			file:    tmpFile,
			wantErr: true,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			lang:    testLang,
			file:    tmpFile,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.CreateCaption(tt.id, tt.lang, tt.file.Name(), tt.file)
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

func TestVideoService_GetVideoCaptions(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testVideoCaptionID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Video.GetCaptions(tt.id, VideoApiGetCaptionsRequest{})
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
		request VideoApiGetCaptionsRequest
		wantErr bool
	}{
		{
			name:    "Valid Get Video Captions",
			id:      testVideoCaptionID,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty Video ID",
			id:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.GetCaptions(tt.id, tt.request)
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

func TestVideoService_SetDefaultCaption(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		lang    string
		wantErr bool
	}{
		{
			name:    "Set other",
			id:      testVideoID,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Video.SetDefaultCaption(tt.id, tt.lang)
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
		lang    string
		wantErr bool
	}{
		{
			name:    "Valid Set Default Caption",
			id:      testVideoID,
			lang:    testLang,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty Video ID",
			id:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.SetDefaultCaption(tt.id, testLang)
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

func TestVideoService_DeleteVideoCaptions(t *testing.T) {
	anonymousTest := []struct {
		name    string
		id      string
		lang    string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testVideoID,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Video.DeleteCaption(tt.id, tt.lang)
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
		lang    string
		wantErr bool
	}{
		{
			name:    "Valid Delete Video Captions",
			id:      testVideoID,
			lang:    testLang,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Empty Video ID",
			id:      "",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Invalid Language",
			id:      testVideoID,
			lang:    "invalid",
			wantErr: true,
		},
		{
			name:    "Empty Language",
			id:      testVideoID,
			lang:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.DeleteCaption(tt.id, tt.lang)
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

func TestVideoService_Delete(t *testing.T) {
	validMetadata := []Metadata{
		{Key: stringPtr("key1"), Value: stringPtr("value1")},
		{Key: stringPtr("key2"), Value: stringPtr("value2")},
	}

	validTags := []string{"tag1", "tag2"}
	createRequest := CreateVideoRequest{
		Title:       stringPtr("Test Video for Deletion"),
		Description: stringPtr("Test Description"),
		IsPublic:    boolPtr(true),
		Metadata:    &validMetadata,
		Qualities:   &[]string{"1080p", "720p", "360p"},
		Tags:        &validTags,
	}

	resp, err := testClient.Video.Create(createRequest)
	if err != nil {
		t.Fatalf("Failed to create video: %v", err)
	}
	testVideoID := *resp.Data.Id

	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testVideoID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Video.Delete(tt.id)
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
			id:      testVideoID,
			wantErr: false,
		},
		{
			name:    "Invalid Video ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty Video ID",
			id:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Video.Delete(tt.id)
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

func stringPtr(s string) *string {
	return &s
}
