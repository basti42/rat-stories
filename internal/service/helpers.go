package service

import "github.com/google/uuid"

func valueOrNilUUID(value *uuid.UUID) uuid.UUID {
	if value != &uuid.Nil {
		return *value
	}
	return uuid.Nil
}
