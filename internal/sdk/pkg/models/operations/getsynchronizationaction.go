// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// GetSynchronizationAction200ApplicationJSON - OK
type GetSynchronizationAction200ApplicationJSON struct {
	Actions []string `json:"actions,omitempty"`
}

type GetSynchronizationActionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetSynchronizationAction200ApplicationJSONObject *GetSynchronizationAction200ApplicationJSON
}
