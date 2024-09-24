package w3streamsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var webhookJSONResponse = `{
	"data": {
		"webhook": {
			"id": "webhookId",
			"url": "webhookUrl",
			"user_id": "userId",
			"encoding_completed": true,
			"created_at":"webhookCreatedAt",
			"updated_at":"webhookUpdatedAt"
		}
	},
	"status": "success"
}`

var webhookListJsonResponse = `{
	"data": {
		"total": 1,
		"webhooks": [
			{
			"id": "webhookId",
			"url": "webhookUrl",
			"user_id": "userId",
			"encoding_completed": true,
			"created_at":"webhookCreatedAt",
			"updated_at":"webhookUpdatedAt"
			}
		]
	},
	"status": "success"
}`

var webHook = &Webhook{
	Id:               PtrString("webhookId"),
	Url:              PtrString("webhookUrl"),
	UserId:           PtrString("userId"),
	EncodingFinished: PtrBool(true),
	EncodingStarted:  PtrBool(false),
	FileReceived:     PtrBool(false),
	LastTriggeredAt:  PtrString("webhookLastTriggeredAt"),
	Name:             PtrString("webhookName"),
	CreatedAt:        PtrString("webhookCreatedAt"),
	UpdatedAt:        PtrString("webhookUpdatedAt"),
}

var getWebhookListResponse = GetWebhooksListResponse{
	Data: &GetWebhooksListData{
		Total: PtrInt32(1),
		Webhooks: &[]Webhook{
			*webHook,
		},
	},
	Status: PtrString("success"),
}

var getWebhookResponse = GetUserWebhookResponse{
	Data: &GetUserWebhookData{
		Webhook: webHook,
	},
	Status: PtrString("success"),
}

var successResponse = `{
	"message": "success",
	"status": "success"
}`

var successResponseStruct = ResponseSuccess{
	Message: PtrString("success"),
	Status:  PtrString("success"),
}

var createWebhookStruct = CreateWebhookResponse{
	Data: &CreateWebhookData{
		Webhook: webHook,
	},
	Status: PtrString("success"),
}

func TestWebhook_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/webhooks/webhookId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, webhookJSONResponse)
	})

	resp, err := client.Webhook.Get("webhookId")
	if err != nil {
		t.Errorf("Webhook.Get error: %v", err)
	}

	expected := &getWebhookResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Webhook.Get\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestWebhook_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, webhookListJsonResponse)
	})

	resp, err := client.Webhook.List(WebhookApiListRequest{
		offset: PtrInt32(0),
		limit:  PtrInt32(10),
	})
	if err != nil {
		t.Errorf("Webhook.List error: %v", err)
	}

	expected := &getWebhookListResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Webhook.List\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestWebhook_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, webhookJSONResponse)
	})

	resp, err := client.Webhook.Create(CreateWebhookRequest{

		EncodingFinished: PtrBool(true),
		EncodingStarted:  PtrBool(true),
		FileReceived:     PtrBool(true),
		Name:             PtrString("webhookName"),
		Url:              PtrString("webhookUrl"),
	})
	if err != nil {
		t.Errorf("Webhook.Create error: %v", err)
	}

	expected := &createWebhookStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Webhook.Create\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestWebhook_Update(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/webhooks/webhookId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Webhook.Update("webhookId", UpdateWebhookRequest{
		EncodingFinished: PtrBool(true),
		EncodingStarted:  PtrBool(true),
		FileReceived:     PtrBool(true),
		Name:             PtrString("webhookName"),
		Url:              PtrString("webhookUrl"),
	})
	if err != nil {
		t.Errorf("Webhook.Update error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Webhook.Update\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestWebhook_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/webhooks/webhookId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.Webhook.Delete("webhookId")
	if err != nil {
		t.Errorf("Webhook.Delete error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Webhook.Delete\n got=%#v\nwant=%#v", resp, expected)
	}
}
