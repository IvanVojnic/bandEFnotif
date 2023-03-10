package rpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
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

// StorageInvite used to storage invite by serv
func (s *InviteRoomServer) StorageInvite(ctx context.Context, req *pr.StorageInviteRequest) (*pr.StorageInviteResponse, error) { // nolint:dupl, gocritic
	userCreatorID, err := uuid.Parse(req.GetUserCreatorID())
	if err != nil {
		return &pr.StorageInviteResponse{}, fmt.Errorf("error while parsing user creator ID, %s", err)
	}
	roomID, err := uuid.Parse(req.GetRoomID())
	if err != nil {
		return &pr.StorageInviteResponse{}, fmt.Errorf("error while parsing room ID, %s", err)
	}
	date, err := time.Parse(timeLayout, req.GetDate())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"date": date,
		}).Errorf("error parsing date (send invite), %s", err)
		return &pr.StorageInviteResponse{}, fmt.Errorf("error while parsing date, %s", err)
	}
	err = s.inviteServ.StorageInvite(ctx, roomID, userCreatorID, req.GetPlace(), date)
	if err != nil {
		logrus.Errorf("error storage invite, %s", err)
		return &pr.StorageInviteResponse{}, fmt.Errorf("error while storage invite, %s", err)
	}
	return &pr.StorageInviteResponse{}, nil
}
