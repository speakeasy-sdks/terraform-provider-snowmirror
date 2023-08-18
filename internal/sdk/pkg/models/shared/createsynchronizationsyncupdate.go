// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateSynchronizationSyncUpdateColumns struct {
	Name *string `json:"name,omitempty"`
}

type CreateSynchronizationSyncUpdateColumnsToExclude struct {
	Name *string `json:"name,omitempty"`
}

type CreateSynchronizationSyncUpdateDeleteStrategy string

const (
	CreateSynchronizationSyncUpdateDeleteStrategyAudit    CreateSynchronizationSyncUpdateDeleteStrategy = "AUDIT"
	CreateSynchronizationSyncUpdateDeleteStrategyTruncate CreateSynchronizationSyncUpdateDeleteStrategy = "TRUNCATE"
	CreateSynchronizationSyncUpdateDeleteStrategyDiff     CreateSynchronizationSyncUpdateDeleteStrategy = "DIFF"
	CreateSynchronizationSyncUpdateDeleteStrategyNone     CreateSynchronizationSyncUpdateDeleteStrategy = "NONE"
)

func (e CreateSynchronizationSyncUpdateDeleteStrategy) ToPointer() *CreateSynchronizationSyncUpdateDeleteStrategy {
	return &e
}

func (e *CreateSynchronizationSyncUpdateDeleteStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AUDIT":
		fallthrough
	case "TRUNCATE":
		fallthrough
	case "DIFF":
		fallthrough
	case "NONE":
		*e = CreateSynchronizationSyncUpdateDeleteStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncUpdateDeleteStrategy: %v", v)
	}
}

type CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType string

const (
	CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionTypeCleanAndSynchronize CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType = "CLEAN_AND_SYNCHRONIZE"
	CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionTypeDifferential        CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType = "DIFFERENTIAL."
)

func (e CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType) ToPointer() *CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType {
	return &e
}

func (e *CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CLEAN_AND_SYNCHRONIZE":
		fallthrough
	case "DIFFERENTIAL.":
		*e = CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType: %v", v)
	}
}

type CreateSynchronizationSyncUpdateFullLoadSchedulerType string

const (
	CreateSynchronizationSyncUpdateFullLoadSchedulerTypeManually     CreateSynchronizationSyncUpdateFullLoadSchedulerType = "MANUALLY"
	CreateSynchronizationSyncUpdateFullLoadSchedulerTypeDaily        CreateSynchronizationSyncUpdateFullLoadSchedulerType = "DAILY"
	CreateSynchronizationSyncUpdateFullLoadSchedulerTypeWeekly       CreateSynchronizationSyncUpdateFullLoadSchedulerType = "WEEKLY"
	CreateSynchronizationSyncUpdateFullLoadSchedulerTypePeriodically CreateSynchronizationSyncUpdateFullLoadSchedulerType = "PERIODICALLY"
	CreateSynchronizationSyncUpdateFullLoadSchedulerTypeCron         CreateSynchronizationSyncUpdateFullLoadSchedulerType = "CRON"
)

func (e CreateSynchronizationSyncUpdateFullLoadSchedulerType) ToPointer() *CreateSynchronizationSyncUpdateFullLoadSchedulerType {
	return &e
}

func (e *CreateSynchronizationSyncUpdateFullLoadSchedulerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANUALLY":
		fallthrough
	case "DAILY":
		fallthrough
	case "WEEKLY":
		fallthrough
	case "PERIODICALLY":
		fallthrough
	case "CRON":
		*e = CreateSynchronizationSyncUpdateFullLoadSchedulerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncUpdateFullLoadSchedulerType: %v", v)
	}
}

type CreateSynchronizationSyncUpdateFullLoadScheduler struct {
	BeginDate     *string                                                        `json:"beginDate,omitempty"`
	ExecutionType *CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType `json:"executionType,omitempty"`
	Type          *CreateSynchronizationSyncUpdateFullLoadSchedulerType          `json:"type,omitempty"`
}

type CreateSynchronizationSyncUpdateSchedulerDays string

const (
	CreateSynchronizationSyncUpdateSchedulerDaysMonday   CreateSynchronizationSyncUpdateSchedulerDays = "MONDAY"
	CreateSynchronizationSyncUpdateSchedulerDaysTuesday  CreateSynchronizationSyncUpdateSchedulerDays = "TUESDAY"
	CreateSynchronizationSyncUpdateSchedulerDaysWednesda CreateSynchronizationSyncUpdateSchedulerDays = "WEDNESDA"
	CreateSynchronizationSyncUpdateSchedulerDaysThursday CreateSynchronizationSyncUpdateSchedulerDays = "THURSDAY"
	CreateSynchronizationSyncUpdateSchedulerDaysFriday   CreateSynchronizationSyncUpdateSchedulerDays = "FRIDAY"
	CreateSynchronizationSyncUpdateSchedulerDaysSaturday CreateSynchronizationSyncUpdateSchedulerDays = "SATURDAY"
	CreateSynchronizationSyncUpdateSchedulerDaysSunday   CreateSynchronizationSyncUpdateSchedulerDays = "SUNDAY"
)

func (e CreateSynchronizationSyncUpdateSchedulerDays) ToPointer() *CreateSynchronizationSyncUpdateSchedulerDays {
	return &e
}

func (e *CreateSynchronizationSyncUpdateSchedulerDays) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONDAY":
		fallthrough
	case "TUESDAY":
		fallthrough
	case "WEDNESDA":
		fallthrough
	case "THURSDAY":
		fallthrough
	case "FRIDAY":
		fallthrough
	case "SATURDAY":
		fallthrough
	case "SUNDAY":
		*e = CreateSynchronizationSyncUpdateSchedulerDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncUpdateSchedulerDays: %v", v)
	}
}

// CreateSynchronizationSyncUpdateSchedulerType - Specifies when the incremental load synchronization will run
type CreateSynchronizationSyncUpdateSchedulerType string

const (
	CreateSynchronizationSyncUpdateSchedulerTypeManually     CreateSynchronizationSyncUpdateSchedulerType = "MANUALLY"
	CreateSynchronizationSyncUpdateSchedulerTypeDaily        CreateSynchronizationSyncUpdateSchedulerType = "DAILY"
	CreateSynchronizationSyncUpdateSchedulerTypeWeekly       CreateSynchronizationSyncUpdateSchedulerType = "WEEKLY"
	CreateSynchronizationSyncUpdateSchedulerTypePeriodically CreateSynchronizationSyncUpdateSchedulerType = "PERIODICALLY"
	CreateSynchronizationSyncUpdateSchedulerTypeCron         CreateSynchronizationSyncUpdateSchedulerType = "CRON"
)

func (e CreateSynchronizationSyncUpdateSchedulerType) ToPointer() *CreateSynchronizationSyncUpdateSchedulerType {
	return &e
}

func (e *CreateSynchronizationSyncUpdateSchedulerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANUALLY":
		fallthrough
	case "DAILY":
		fallthrough
	case "WEEKLY":
		fallthrough
	case "PERIODICALLY":
		fallthrough
	case "CRON":
		*e = CreateSynchronizationSyncUpdateSchedulerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncUpdateSchedulerType: %v", v)
	}
}

type CreateSynchronizationSyncUpdateScheduler struct {
	BeginDate            *string                                        `json:"beginDate,omitempty"`
	Days                 []CreateSynchronizationSyncUpdateSchedulerDays `json:"days,omitempty"`
	IncLoadExecutionType *string                                        `json:"incLoadExecutionType,omitempty"`
	Time                 *string                                        `json:"time,omitempty"`
	// Specifies when the incremental load synchronization will run
	Type    *CreateSynchronizationSyncUpdateSchedulerType `json:"type,omitempty"`
	Visible *bool                                         `json:"visible,omitempty"`
}

type CreateSynchronizationSyncUpdateSchedulerPriority string

const (
	CreateSynchronizationSyncUpdateSchedulerPriorityHighest CreateSynchronizationSyncUpdateSchedulerPriority = "HIGHEST"
	CreateSynchronizationSyncUpdateSchedulerPriorityHigh    CreateSynchronizationSyncUpdateSchedulerPriority = "HIGH"
	CreateSynchronizationSyncUpdateSchedulerPriorityNormal  CreateSynchronizationSyncUpdateSchedulerPriority = "NORMAL"
	CreateSynchronizationSyncUpdateSchedulerPriorityLow     CreateSynchronizationSyncUpdateSchedulerPriority = "LOW"
	CreateSynchronizationSyncUpdateSchedulerPriorityLowest  CreateSynchronizationSyncUpdateSchedulerPriority = "LOWEST"
)

func (e CreateSynchronizationSyncUpdateSchedulerPriority) ToPointer() *CreateSynchronizationSyncUpdateSchedulerPriority {
	return &e
}

func (e *CreateSynchronizationSyncUpdateSchedulerPriority) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HIGHEST":
		fallthrough
	case "HIGH":
		fallthrough
	case "NORMAL":
		fallthrough
	case "LOW":
		fallthrough
	case "LOWEST":
		*e = CreateSynchronizationSyncUpdateSchedulerPriority(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncUpdateSchedulerPriority: %v", v)
	}
}

type CreateSynchronizationSyncUpdate struct {
	// true - synchronization is active and can be scheduled to synchronize data from ServiceNow
	// false - synchronization is deactivated and cannot be scheduled to synchronize data from ServiceNowNow
	Active *bool `json:"active,omitempty"`
	// SnowMirror checks if columns exist in ServiceNow. If this flag is set to true,
	AllowInheritedColumns *bool `json:"allowInheritedColumns,omitempty"`
	// Configures how to check for schema changes in ServiceNow.
	//
	// Enabled (true) - whenever a synchronization is executed, SnowMirror checks for schema changes in ServiceNow. Automatically adds, updates (data type, max. length of a column) and removes columns. If a new column is added SnowMirror clears the mirror table and downloads all records from scratch.
	// Enabled (no truncation) (ENABLED_WITHOUT_TRUNCATION) - the same as Enabled option. It handles new columns differently, though. If a new column is added SnowMirror does not clear the mirror table. Instead, it creates the column and populates it with a default value (which is defined in ServiceNow).
	//
	AutoSchemaUpdate  *string                                           `json:"autoSchemaUpdate,omitempty"`
	Columns           []CreateSynchronizationSyncUpdateColumns          `json:"columns,omitempty"`
	ColumnsToExclude  []CreateSynchronizationSyncUpdateColumnsToExclude `json:"columnsToExclude,omitempty"`
	DeleteStrategy    *CreateSynchronizationSyncUpdateDeleteStrategy    `json:"deleteStrategy,omitempty"`
	EncodedQuery      *string                                           `json:"encodedQuery,omitempty"`
	FullLoadScheduler *CreateSynchronizationSyncUpdateFullLoadScheduler `json:"fullLoadScheduler,omitempty"`
	// Id of the synchronization.
	ID *int64 `json:"id,omitempty"`
	// Name of the table in mirror database where the data will be migrated.
	MirrorTable *string `json:"mirrorTable,omitempty"`
	// Display name of the synchronization.
	Name *string `json:"name,omitempty"`
	// Defines how to synchronize reference field types.
	ReferenceFieldType *string                                           `json:"referenceFieldType,omitempty"`
	Scheduler          *CreateSynchronizationSyncUpdateScheduler         `json:"scheduler,omitempty"`
	SchedulerPriority  *CreateSynchronizationSyncUpdateSchedulerPriority `json:"schedulerPriority,omitempty"`
	// Name of the table in ServiceNow.
	Table *string `json:"table,omitempty"`
	// Name of the view in ServiceNow.
	View *string `json:"view,omitempty"`
}
