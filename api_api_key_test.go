package w3streamsdk

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testApiKeyForUpdateAndDelete string
	testApiKeyName               string = "Test API Key"
)

func TestApiKeyService_Create(t *testing.T) {
	tests := []struct {
		name    string
		request CreateApiKeyRequest
		wantErr bool
	}{
		{
			name: "Valid Create",
			request: CreateApiKeyRequest{
				ApiKeyName: stringPtr(testApiKeyName),
				Type:       stringPtr("full_access"),
				Ttl:        stringPtr("100000000"),
			},
			wantErr: false,
		},
		{
			name: "Empty Name",
			request: CreateApiKeyRequest{
				Type: stringPtr("read"),
				Ttl:  stringPtr("24h"),
			},
			wantErr: true,
		},
		{
			name: "Invalid Type",
			request: CreateApiKeyRequest{
				ApiKeyName: stringPtr(testApiKeyName),
				Type:       stringPtr("invalid"),
				Ttl:        stringPtr("24h"),
			},
			wantErr: true,
		},
		{
			name: "Invalid TTL",
			request: CreateApiKeyRequest{
				ApiKeyName: stringPtr(testApiKeyName),
				Type:       stringPtr("read"),
				Ttl:        stringPtr("invalid"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKey.Create(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				testApiKeyForUpdateAndDelete = *resp.Data.ApiKey.Id
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data.ApiKey.Id)
				assert.Equal(t, *tt.request.ApiKeyName, *resp.Data.ApiKey.Name)
			}
		})
	}
}

func TestApiKeyService_Update(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
		request RenameAPIKeyRequest
	}{
		{
			name: "Update other",
			id:   testApiKeyForUpdateAndDelete,
			request: RenameAPIKeyRequest{
				ApiKeyName: stringPtr("Updated API Key"),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.ApiKey.Update(tt.id, tt.request)
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
		request RenameAPIKeyRequest
		wantErr bool
	}{
		{
			name: "Valid Update",
			id:   testApiKeyForUpdateAndDelete,
			request: RenameAPIKeyRequest{
				ApiKeyName: stringPtr("Updated API Key"),
			},
			wantErr: false,
		},
		{
			name: "Invalid ID",
			id:   "invalid-id",
			request: RenameAPIKeyRequest{
				ApiKeyName: stringPtr("Updated API Key"),
			},
			wantErr: true,
		},
		{
			name: "Empty Name",
			id:   testApiKeyForUpdateAndDelete,
			request: RenameAPIKeyRequest{
				ApiKeyName: stringPtr(""),
			},
			wantErr: true,
		},
		{
			name: "Not Exist ID",
			id:   notExistId,
			request: RenameAPIKeyRequest{
				ApiKeyName: stringPtr("Updated API Key"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKey.Update(tt.id, tt.request)
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

func TestApiKeyService_Delete(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testApiKeyForUpdateAndDelete,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.ApiKey.Delete(tt.id)
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
			id:      testApiKeyForUpdateAndDelete,
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
			resp, err := testClient.ApiKey.Delete(tt.id)
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

func TestApiKeyService_List(t *testing.T) {
	tests := []struct {
		name    string
		request ApiKeyApiListRequest
		wantErr bool
	}{
		{
			name:    "Valid List No Filters",
			request: ApiKeyApiListRequest{},
			wantErr: false,
		},
		{
			name: "Valid List With Filters",
			request: ApiKeyApiListRequest{}.
				Search("test").
				SortBy("created_at").
				OrderBy("desc").
				Offset(0).
				Limit(10),
			wantErr: false,
		},
		{
			name: "Invalid Offset",
			request: ApiKeyApiListRequest{}.
				Offset(-1),
			wantErr: true,
		},
		{
			name: "Invalid Limit",
			request: ApiKeyApiListRequest{}.
				Limit(1001),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKey.List(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
				assert.NotNil(t, resp.Data.ApiKeys)
			}
		})
	}
}
