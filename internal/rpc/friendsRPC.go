package rpc

import (
	"context"
	"time"

	pr "github.com/IvanVojnic/bandEFnotif/proto"
	"github.com/google/uuid"
)

// Friends is an interface with implemented methods from Invite service
type Friends interface {
	StorageFriendsRequest(ctx context.Context, roomID, userCreatorID uuid.UUID, place string, date time.Time) error
}

// FriendsServer used to define invite server obj
type FriendsServer struct {
	pr.UnimplementedFriendsServer
	friendsServ Friends
}

// NewFriendsServer used to init invite serv obj
func NewFriendsServer(friendsServ Friends) *FriendsServer {
	return &FriendsServer{friendsServ: friendsServ}
}
