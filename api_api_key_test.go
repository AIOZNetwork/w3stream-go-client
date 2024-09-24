package w3streamsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var apiKeyListJSONResponse = `{
	"data": {
		"api_keys": [
			{
			"created_at": "createdAtValue",
			"expired_at": "expiredAtValue",
			"id": "idValue",
			"last_used_at": "lastUsedAtValue",
			"name": "nameValue",
			"public_key": "publicKeyValue",
			"secret": "secretValue",
			"truncated_secret": "truncatedSecretValue",
			"ttl": "ttlValue",
			"type": "typeValue",
			"updated_at": "updatedAtValue"
			}
		],
		"total": 1
	},
	"status": "success"
}`

var apiKeyJsonResponse = `{
	"data": {
		"api_key": {
			"created_at": "createdAtValue",
			"expired_at": "expiredAtValue",
			"id": "idValue",
			"last_used_at": "lastUsedAtValue",
			"name": "nameValue",
			"public_key": "publicKeyValue",
			"secret": "secretValue",
			"truncated_secret": "truncatedSecretValue",
			"ttl": "ttlValue",
			"type": "typeValue",
			"updated_at": "updatedAtValue"
		}
	},
	"status": "success"
}
`

var apiKey = &ApiKey{
	CreatedAt:       PtrString("createdAtValue"),
	ExpiredAt:       PtrString("expiredAtValue"),
	Id:              PtrString("idValue"),
	LastUsedAt:      PtrString("lastUsedAtValue"),
	Name:            PtrString("nameValue"),
	PublicKey:       PtrString("publicKeyValue"),
	Secret:          PtrString("secretValue"),
	TruncatedSecret: PtrString("truncatedSecretValue"),
	Ttl:             PtrString("ttlValue"),
	Type:            PtrString("typeValue"),
	UpdatedAt:       PtrString("updatedAtValue"),
}

var apiKeyListResponse = GetApiKeysResponse{
	Data: &GetApiKeysData{
		ApiKeys: &[]ApiKey{
			*apiKey,
		},
		Total: PtrInt32(1),
	},
	Status: PtrString("success"),
}

var apiKeyResponse = CreateApiKeyResponse{
	Data: &CreateApiKeyData{
		ApiKey: apiKey,
	},
	Status: PtrString("success"),
}

func TestApiKey_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/api_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, apiKeyJsonResponse)
	})

	resp, err := client.ApiKey.Create(CreateApiKeyRequest{
		ApiKeyName: PtrString("nameValue"),
		Ttl:        PtrString("ttlValue"),
		Type:       PtrString("typeValue"),
	})
	if err != nil {
		t.Errorf("ApiKey.Create error: %v", err)
	}

	expected := &apiKeyResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("ApiKey.Create\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestApiKey_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/api_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, apiKeyListJSONResponse)
	})

	resp, err := client.ApiKey.List(ApiKeyApiListRequest{
		offset: PtrInt32(0),
		limit:  PtrInt32(10),
	})
	if err != nil {
		t.Errorf("ApiKey.List error: %v", err)
	}

	expected := &apiKeyListResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("ApiKey.List\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestApiKey_Rename(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/api_keys/apiKeyId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.ApiKey.Update("apiKeyId", RenameAPIKeyRequest{
		ApiKeyName: PtrString("nameValue"),
	})
	if err != nil {
		t.Errorf("ApiKey.Rename error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("ApiKey.Rename\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestApiKey_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/api_keys/apiKeyId", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, successResponse)
	})

	resp, err := client.ApiKey.Delete("apiKeyId")
	if err != nil {
		t.Errorf("ApiKey.Delete error: %v", err)
	}

	expected := &successResponseStruct
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("ApiKey.Delete\n got=%#v\nwant=%#v", resp, expected)
	}
}
