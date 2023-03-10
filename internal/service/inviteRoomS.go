package service

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type InviteRoom interface {
	StorageInvite(ctx context.Context, roomID, userCreatorID uuid.UUID, place string, date time.Time) error
}

// InviteRoomService used to define invite server obj
type InviteRoomService struct {
	inviteRepo InviteRoom
}

func NewInviteRoomService(inviteRepo InviteRoom) *InviteRoomService {
	return &InviteRoomService{inviteRepo: inviteRepo}
}

func (s *InviteRoomService) StorageInvite(ctx context.Context, roomID, userCreatorID uuid.UUID, place string, date time.Time) error {
	return s.inviteRepo.StorageInvite(ctx, roomID, userCreatorID, place, date)
}
