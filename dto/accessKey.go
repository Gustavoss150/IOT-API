package dto

import "api/entities"

type AccessKeyDTO struct {
	TypeKey       entities.TypeKey   `json:"type_key"`
	StatusKey     entities.StatusKey `json:"status_key"`
	Value         string             `json:"value"`
	AssignedTo    string             `json:"assigned_to"`
	ReservationID string             `json:"reservation_id"`
}
