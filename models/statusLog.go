package models

import "time"

type StatusLog struct {
	ID         int       `json:"id"`
	EntityID   int       `json:"entity_id"`
	EntityType string    `json:"entity_type"`
	OldStatus  string    `json:"old_status"`
	NewStatus  string    `json:"new_status"`
	ChangedBy  string    `json:"changed_by"`
	Timestamp  time.Time `json:"timestamp"`
}

func NewStatusLog(entityID int, entityType, oldStatus, newStatus, changedBy string) *StatusLog {
	return &StatusLog{
		EntityID:   entityID,
		EntityType: entityType,
		OldStatus:  oldStatus,
		NewStatus:  newStatus,
		ChangedBy:  changedBy,
		Timestamp:  time.Now(),
	}
}
