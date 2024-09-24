package w3streamsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var getWatermarkListJSONResponse = `
{
    "status": "success",
    "data": {
        "watermarks": [
            {
                "id": "waterMarkId",
                "user_id": "userId",
                "watermark_name": "waterMarkName",
                "width": 100,
                "height": 100,
                "created_at": "waterMarkCreatedAt",
                "updated_at": "watermarkUpdatedAt"
            }
	],
	"total":1
	}
}
`

var getWatermarkJSONResponse = `
	{
	"status": "success",
    "data": {
            "watermark_id": "waterMarkId",
            "created_at": "waterMarkCreatedAt"
    }
	}
`

var watermark = Watermark{
	Id:            PtrString("waterMarkId"),
	UserId:        PtrString("userId"),
	WatermarkName: PtrString("waterMarkName"),
	Width:         PtrInt32(100),
	Height:        PtrInt32(100),
	CreatedAt:     PtrString("waterMarkCreatedAt"),
	UpdatedAt:     PtrString("watermarkUpdatedAt"),
}

var getWatermarkListResponse = GetAllWatermarkResponse{
	Data: &GetAllWatermarkData{
		Watermarks: &[]Watermark{
			watermark,
		},
		Total: PtrInt32(1),
	},
	Status: PtrString("success"),
}

var getWatermarkResponse = CreateWatermarkResponse{
	Data: &CreateWatermarkData{
		WatermarkId: PtrString("waterMarkId"),
		CreatedAt:   PtrString("waterMarkCreatedAt"),
	},
	Status: PtrString("success"),
}

func TestWatermark_Upload(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/watermarks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, getWatermarkJSONResponse)
	})

	reader := strings.NewReader("data")
	resp, err := client.Watermark.Upload("png.png", reader)
	if err != nil {
		t.Errorf("Watermark.Create error: %v", err)
	}

	expected := &getWatermarkResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Watermark.Create\n got=%#v\nwant=%#v", resp, expected)
	}
}

func TestWatermark_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/watermarks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, getWatermarkListJSONResponse)
	})

	resp, err := client.Watermark.List(WatermarkApiListRequest{
		offset: PtrInt32(0),
		limit:  PtrInt32(10),
	})
	if err != nil {
		t.Errorf("Watermark.List error: %v", err)
	}

	expected := &getWatermarkListResponse
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Watermark.List\n got=%#v\nwant=%#v", resp, expected)
	}
}
