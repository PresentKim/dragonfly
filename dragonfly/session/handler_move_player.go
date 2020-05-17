package session

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// MovePlayerHandler handles the MovePlayer packet.
type MovePlayerHandler struct{}

// Handle ...
func (*MovePlayerHandler) Handle(p packet.Packet, s *Session) error {
	pk := p.(*packet.MovePlayer)

	if pk.EntityRuntimeID != selfEntityRuntimeID {
		return ErrSelfRuntimeID
	}
	pk.Position = pk.Position.Sub(mgl32.Vec3{0, 1.62}) // Subtract the base offset of players from the pos.

	s.c.Move(pk.Position.Sub(s.c.Position()))
	s.c.Rotate(pk.Yaw-s.c.Yaw(), pk.Pitch-s.c.Pitch())

	s.chunkLoader.Move(pk.Position)
	s.writePacket(&packet.NetworkChunkPublisherUpdate{
		Position: protocol.BlockPos{int32(pk.Position[0]), int32(pk.Position[1]), int32(pk.Position[2])},
		Radius:   uint32(s.chunkRadius * 16),
	})
	return nil
}