// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SyncronizationSyncOutput struct {
	// Sync ID
	ID int64 `json:"id"`
	// Display name of the synchronization.
	Name string `json:"name"`
	// Name of the table in ServiceNow.
	Table *string `json:"table,omitempty"`
}