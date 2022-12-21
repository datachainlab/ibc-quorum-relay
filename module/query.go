package module

import (
	"context"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
)

// QueryClientConsensusState retrevies the latest consensus state for a client in state at a given height
func (c *Chain) QueryClientConsensusState(height int64, consensusHeight exported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	bz, found, err := c.ibcHost.GetConsensusState(
		c.CallOpts(context.TODO(), height),
		c.pathEnd.ClientID,
		pbToHostHeight(consensusHeight),
	)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("client consensus not found: %s", c.pathEnd.ClientID)
	}

	var consensusState exported.ConsensusState
	if err := c.Codec().UnmarshalInterface(bz, &consensusState); err != nil {
		return nil, err
	}

	consensusStateAny, err := clienttypes.PackConsensusState(consensusState)
	if err != nil {
		return nil, err
	}

	return clienttypes.NewQueryConsensusStateResponse(consensusStateAny, nil, ethHeightToPB(height)), nil
}

// QueryClientState returns the client state of dst chain
// height represents the height of dst chain
func (c *Chain) QueryClientState(height int64) (*clienttypes.QueryClientStateResponse, error) {
	bz, found, err := c.ibcHost.GetClientState(
		c.CallOpts(context.TODO(), height),
		c.pathEnd.ClientID,
	)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("client not found: %s", c.pathEnd.ClientID)
	}

	var clientState exported.ClientState
	if err := c.Codec().UnmarshalInterface(bz, &clientState); err != nil {
		return nil, err
	}

	clientStateAny, err := clienttypes.PackClientState(clientState)
	if err != nil {
		return nil, err
	}

	return clienttypes.NewQueryClientStateResponse(clientStateAny, nil, ethHeightToPB(height)), nil
}

// QueryConnection returns the remote end of a given connection
func (c *Chain) QueryConnection(height int64) (*conntypes.QueryConnectionResponse, error) {
	connEnd, found, err := c.ibcHost.GetConnection(
		c.CallOpts(context.TODO(), height),
		c.pathEnd.ConnectionID,
	)
	if err != nil {
		return nil, err
	} else if !found {
		connEnd = ibchost.ConnectionEndData{State: uint8(conntypes.UNINITIALIZED)}
	}
	return conntypes.NewQueryConnectionResponse(connectionEndToPB(connEnd), nil, ethHeightToPB(height)), nil
}

// QueryChannel returns the channel associated with a channelID
func (c *Chain) QueryChannel(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	chann, found, err := c.ibcHost.GetChannel(
		c.CallOpts(context.TODO(), height),
		c.pathEnd.PortID,
		c.pathEnd.ChannelID,
	)
	if err != nil {
		return nil, err
	} else if !found {
		chann = ibchost.ChannelData{State: uint8(chantypes.UNINITIALIZED)}
	}
	return chantypes.NewQueryChannelResponse(channelToPB(chann), nil, ethHeightToPB(height)), nil
}

// QueryPacketCommitment returns the packet commitment corresponding to a given sequence
func (c *Chain) QueryPacketCommitment(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	commitment, found, err := c.ibcHost.GetPacketCommitment(
		c.CallOpts(context.TODO(), height),
		c.pathEnd.PortID,
		c.pathEnd.ChannelID, seq)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf(
			"packet commitment not found: portId=%v, channelId=%v, seq=%v",
			c.pathEnd.PortID, c.pathEnd.ChannelID, seq,
		)
	}
	return chantypes.NewQueryPacketCommitmentResponse(commitment[:], nil, ethHeightToPB(height)), nil
}

// QueryPacketAcknowledgementCommitment returns the acknowledgement corresponding to a given sequence
func (c *Chain) QueryPacketAcknowledgementCommitment(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	commitment, found, err := c.ibcHost.GetPacketAcknowledgementCommitment(
		c.CallOpts(context.TODO(), height),
		c.pathEnd.PortID,
		c.pathEnd.ChannelID, seq)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf(
			"acknowledgement commitment not found: portId=%v, channelId=%v, seq=%v",
			c.pathEnd.PortID, c.pathEnd.ChannelID, seq,
		)
	}
	return chantypes.NewQueryPacketAcknowledgementResponse(commitment[:], nil, ethHeightToPB(height)), nil
}

// QueryPacketCommitments returns an array of packet commitments
func (c *Chain) QueryPacketCommitments(offset, limit uint64, height int64) (comRes *chantypes.QueryPacketCommitmentsResponse, err error) {
	packets, err := c.findAllPackets(
		context.TODO(),
		height,
		c.pathEnd.PortID,
		c.pathEnd.ChannelID,
	)
	if err != nil {
		return nil, err
	}
	var res chantypes.QueryPacketCommitmentsResponse
	for _, p := range packets {
		ps := chantypes.NewPacketState(
			p.SourcePort,
			p.SourceChannel,
			p.Sequence,
			chantypes.CommitPacket(c.Codec(), p),
		)
		res.Commitments = append(res.Commitments, &ps)
	}
	res.Height = ethHeightToPB(height)
	return &res, nil
}

// QueryUnrecievedPackets returns a list of unrelayed packet commitments
func (c *Chain) QueryUnrecievedPackets(height int64, seqs []uint64) ([]uint64, error) {
	var ret []uint64
	for _, seq := range seqs {
		found, err := c.ibcHost.HasPacketReceipt(
			c.CallOpts(context.TODO(), height),
			c.pathEnd.PortID,
			c.pathEnd.ChannelID,
			seq,
		)
		if err != nil {
			return nil, err
		} else if !found {
			ret = append(ret, seq)
		}
	}
	return ret, nil
}

// QueryPacketAcknowledgementCommitments returns an array of packet acks
func (c *Chain) QueryPacketAcknowledgementCommitments(offset, limit uint64, height int64) (comRes *chantypes.QueryPacketAcknowledgementsResponse, err error) {
	acks, err := c.findAllAcknowledgements(
		context.TODO(),
		height,
		c.pathEnd.PortID,
		c.pathEnd.ChannelID,
	)
	if err != nil {
		return nil, err
	}
	var res chantypes.QueryPacketAcknowledgementsResponse
	for _, a := range acks {
		ps := chantypes.NewPacketState(
			a.DestinationPortId,
			a.DestinationChannel,
			a.Sequence,
			a.Acknowledgement,
		)
		res.Acknowledgements = append(res.Acknowledgements, &ps)
	}
	res.Height = ethHeightToPB(height)
	return &res, nil
}

// QueryUnrecievedAcknowledgements returns a list of unrelayed packet acks
func (c *Chain) QueryUnrecievedAcknowledgements(height int64, seqs []uint64) ([]uint64, error) {
	var ret []uint64
	for _, seq := range seqs {
		_, found, err := c.ibcHost.GetPacketCommitment(
			c.CallOpts(context.TODO(), height),
			c.pathEnd.PortID,
			c.pathEnd.ChannelID,
			seq,
		)
		if err != nil {
			return nil, err
		} else if found {
			ret = append(ret, seq)
		}
	}
	return ret, nil
}

// QueryPacket returns the packet corresponding to a sequence
func (c *Chain) QueryPacket(height int64, sequence uint64) (*chantypes.Packet, error) {
	return c.findPacket(context.TODO(), height, c.pathEnd.PortID, c.pathEnd.ChannelID, sequence)
}

// QueryPacketAcknowledgement returns the acknowledgement corresponding to a sequence
func (c *Chain) QueryPacketAcknowledgement(height int64, sequence uint64) ([]byte, error) {
	return c.findAcknowledgement(
		context.TODO(),
		height,
		c.pathEnd.PortID,
		c.pathEnd.ChannelID,
		sequence,
	)
}

// QueryBalance returns the amount of coins in the relayer account
func (c *Chain) QueryBalance(address sdk.AccAddress) (sdk.Coins, error) {
	panic("not implemented")
}

// QueryDenomTraces returns all the denom traces from a given chain
func (c *Chain) QueryDenomTraces(offset, limit uint64, height int64) (*transfertypes.QueryDenomTracesResponse, error) {
	panic("not implemented")
}
