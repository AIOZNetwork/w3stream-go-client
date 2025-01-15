package w3streamsdk

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testWebhookForUpdateAndDelete string
	webhookURL                    = "https://webhook.site/335e64d4-96f7-4bef-906a-b8cd3862a071"
	webhookName                   = "Test Webhook"
	deleteWebhooksLater           []string
)

func boolPtr(b bool) *bool {
	return &b
}

func TestWebhookService_Create(t *testing.T) {
	tests := []struct {
		name    string
		request CreateWebhookRequest
		wantErr bool
	}{
		{
			name: "Valid Create Request with All Fields",
			request: CreateWebhookRequest{
				EncodingFinished: boolPtr(true),
				EncodingStarted:  boolPtr(true),
				FileReceived:     boolPtr(true),
				Name:             stringPtr(webhookName),
				Url:              stringPtr(webhookURL),
			},
			wantErr: false,
		},
		{
			name: "Invalid Create Request Without Events",
			request: CreateWebhookRequest{
				Url:  stringPtr(webhookURL),
				Name: stringPtr(webhookName),
			},
			wantErr: true,
		},
		{
			name: "Invalid URL",
			request: CreateWebhookRequest{
				Url:  stringPtr("not-a-url"),
				Name: stringPtr(webhookName),
			},
			wantErr: true,
		},
		{
			name: "Missing URL",
			request: CreateWebhookRequest{
				Name: stringPtr(webhookName),
			},
			wantErr: true,
		},
		{
			name: "Missing Name",
			request: CreateWebhookRequest{
				Url:              stringPtr(webhookURL),
				EncodingFinished: boolPtr(true),
				EncodingStarted:  boolPtr(true),
				FileReceived:     boolPtr(true),
			},
			wantErr: false,
		},
		{
			name:    "Empty Request",
			request: CreateWebhookRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Webhook.Create(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data.Webhook.Id)
				deleteWebhooksLater = append(deleteWebhooksLater, *resp.Data.Webhook.Id)
				testWebhookForUpdateAndDelete = *resp.Data.Webhook.Id
			}
		})
	}
}

func TestWebhookService_Update(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		request UpdateWebhookRequest
		wantErr bool
	}{
		{
			name: "Update other",
			id:   testWebhookForUpdateAndDelete,
			request: UpdateWebhookRequest{
				Name: stringPtr("Updated Webhook"),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Webhook.Update(tt.id, tt.request)
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
		request UpdateWebhookRequest
		wantErr bool
	}{
		{
			name: "Valid Update All Fields",
			id:   testWebhookForUpdateAndDelete,
			request: UpdateWebhookRequest{
				EncodingFinished: boolPtr(true),
				EncodingStarted:  boolPtr(false),
				FileReceived:     boolPtr(true),
				Name:             stringPtr("Updated Webhook"),
				Url:              stringPtr(webhookURL),
			},
			wantErr: false,
		},
		{
			name: "Update Partial Fields, only Name",
			id:   testWebhookForUpdateAndDelete,
			request: UpdateWebhookRequest{
				Name: stringPtr("Updated Name Only"),
			},
			wantErr: true,
		},
		{
			name: "Invalid URL",
			id:   testWebhookForUpdateAndDelete,
			request: UpdateWebhookRequest{
				Url: stringPtr("not-a-url"),
			},
			wantErr: true,
		},
		{
			name:    "Invalid ID",
			id:      "invalid-id",
			request: UpdateWebhookRequest{},
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			request: UpdateWebhookRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Webhook.Update(tt.id, tt.request)
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

func TestWebhookService_List(t *testing.T) {
	tests := []struct {
		name    string
		request WebhookApiListRequest
		wantErr bool
		checkFn func(*testing.T, *GetWebhooksListResponse)
	}{
		{
			name:    "List All Webhooks",
			request: WebhookApiListRequest{},
			wantErr: false,
			checkFn: func(t *testing.T, resp *GetWebhooksListResponse) {
				assert.NotNil(t, resp.Data)
			},
		},
		{
			name: "List with Pagination",
			request: WebhookApiListRequest{}.
				Limit(10).
				Offset(0),
			wantErr: false,
			checkFn: func(t *testing.T, resp *GetWebhooksListResponse) {
				assert.NotNil(t, resp.Data)
				assert.LessOrEqual(t, len(*resp.Data.Webhooks), 10)
			},
		},
		{
			name: "List with Search",
			request: WebhookApiListRequest{}.
				Search("test"),
			wantErr: false,
		},
		{
			name: "List with Event Filters",
			request: WebhookApiListRequest{}.
				EncodingFinished(true).
				EncodingStarted(false).
				FileReceived(true),
			wantErr: false,
		},
		{
			name: "Invalid Offset",
			request: WebhookApiListRequest{}.
				Offset(-1),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Webhook.List(tt.request)
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

func TestWebhookService_Get(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testWebhookForUpdateAndDelete,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Webhook.Get(tt.id)
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
		checkFn func(*testing.T, *GetUserWebhookResponse)
	}{
		{
			name:    "Valid Get",
			id:      testWebhookForUpdateAndDelete,
			wantErr: false,
			checkFn: func(t *testing.T, resp *GetUserWebhookResponse) {
				assert.NotEmpty(t, resp.Data.Webhook.Id)
				assert.NotEmpty(t, resp.Data.Webhook.Url)
				assert.NotEmpty(t, resp.Data.Webhook.Name)
			},
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
			resp, err := testClient.Webhook.Get(tt.id)
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

func TestWebhookService_Delete(t *testing.T) {
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testWebhookForUpdateAndDelete,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Webhook.Delete(tt.id)
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
			id:      testWebhookForUpdateAndDelete,
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
			resp, err := testClient.Webhook.Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	for _, id := range deleteWebhooksLater {
		testClient.Webhook.Delete(id)
	}
}
