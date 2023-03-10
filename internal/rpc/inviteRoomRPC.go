package rpc

import (
	"context"
	"time"

	pr "github.com/IvanVojnic/bandEFnotif/proto"
	"github.com/google/uuid"
)

// InviteRoom is an interface with implemented methods from Invite service
type InviteRoom interface {
	StorageInvite(ctx context.Context, roomID, userCreatorID uuid.UUID, place string, date time.Time) error
}

// InviteRoomServer used to define invite server obj
type InviteRoomServer struct {
	pr.UnimplementedInviteRoomServer
	inviteServ InviteRoom
}

// timeLayout define layout for date parsing
const timeLayout = "2006-01-02 15:04:05"

// NewInviteRoomServer used to init invite serv obj
func NewInviteRoomServer(inviteServ InviteRoom) *InviteRoomServer {
	return &InviteRoomServer{inviteServ: inviteServ}
}
