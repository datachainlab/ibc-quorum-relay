package module

import (
	"context"
	fmt "fmt"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
)

// findPacket returns a packet that matches sourcePortID, sourceChannelID, and sequence
func (c *Chain) findPacket(
	ctx context.Context,
	height int64,
	sourcePortID string,
	sourceChannelID string,
	sequence uint64,
) (*chantypes.Packet, error) {
	packets, err := c.findAllPackets(ctx, height, sourcePortID, sourceChannelID)
	if err != nil {
		return nil, err
	}
	for _, p := range packets {
		if p.Sequence == sequence {
			return p, nil
		}
	}
	return nil, fmt.Errorf(
		"packet not found: sourcePortID=%v, sourceChannelID=%v, sequence=%v",
		sourcePortID, sourceChannelID, sequence)
}

// findAllPackets returns all packets that match sourcePortID and sourceChannelID from events
func (c *Chain) findAllPackets(
	ctx context.Context,
	height int64,
	sourcePortID string,
	sourceChannel string,
) ([]*chantypes.Packet, error) {
	end := uint64(height)
	opts := &bind.FilterOpts{
		End:     &end,
		Context: ctx,
	}
	it, err := c.ibcHandler.FilterSendPacket(opts)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var packets []*chantypes.Packet
	for it.Next() {
		p := it.Event.Packet
		if p.SourcePort == sourcePortID && p.SourceChannel == sourceChannel {
			packet := &chantypes.Packet{
				Sequence:           p.Sequence,
				SourcePort:         p.SourcePort,
				SourceChannel:      p.SourceChannel,
				DestinationPort:    p.DestinationPort,
				DestinationChannel: p.DestinationChannel,
				Data:               p.Data,
				TimeoutHeight:      clienttypes.Height(p.TimeoutHeight),
				TimeoutTimestamp:   p.TimeoutTimestamp,
			}
			packets = append(packets, packet)
		}
	}
	if err := it.Error(); err != nil {
		return nil, err
	}
	return packets, nil
}

// findAcknowledgement returns an ack data that matches dstPortID, dstChannelID, and sequence
func (c *Chain) findAcknowledgement(
	ctx context.Context,
	height int64,
	dstPortID string,
	dstChannelID string,
	sequence uint64,
) ([]byte, error) {
	acks, err := c.findAllAcknowledgements(ctx, height, dstPortID, dstChannelID)
	if err != nil {
		return nil, err
	}
	for _, ack := range acks {
		if ack.Sequence == sequence {
			return ack.Acknowledgement, nil
		}
	}
	return nil, fmt.Errorf(
		"ack not found: dstPortID=%v, dstChannelID=%v, sequence=%v",
		dstPortID, dstChannelID, sequence)
}

// findAllAcknowledgements returns all WriteAcknowledgement events that match dstPortID and dstChannelID
func (c *Chain) findAllAcknowledgements(
	ctx context.Context,
	height int64,
	dstPortID string,
	dstChannelID string,
) ([]*ibchandler.IbchandlerWriteAcknowledgement, error) {
	end := uint64(height)
	opts := &bind.FilterOpts{
		End:     &end,
		Context: ctx,
	}
	it, err := c.ibcHandler.FilterWriteAcknowledgement(opts)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var events []*ibchandler.IbchandlerWriteAcknowledgement
	for it.Next() {
		ev := it.Event
		if ev.DestinationPortId == dstPortID && ev.DestinationChannel == dstChannelID {
			events = append(events, ev)
		}
	}
	if err := it.Error(); err != nil {
		return nil, err
	}
	return events, nil
}
