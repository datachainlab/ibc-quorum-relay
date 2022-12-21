package module

import (
	"context"
	"log"

	"github.com/avast/retry-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	proto "github.com/gogo/protobuf/proto"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
)

// SendMsgs sends msgs to the chain
func (c *Chain) SendMsgs(msgs []sdk.Msg) ([]byte, error) {
	ctx := context.TODO()
	for _, msg := range msgs {
		var (
			tx  *types.Transaction
			err error
		)
		opts := c.TransactOpts(ctx)
		switch msg := msg.(type) {
		case *clienttypes.MsgCreateClient:
			tx, err = c.txCreateClient(opts, msg)
		case *clienttypes.MsgUpdateClient:
			tx, err = c.txUpdateClient(opts, msg)
		case *conntypes.MsgConnectionOpenInit:
			tx, err = c.txConnectionOpenInit(opts, msg)
		case *conntypes.MsgConnectionOpenTry:
			tx, err = c.txConnectionOpenTry(opts, msg)
		case *conntypes.MsgConnectionOpenAck:
			tx, err = c.txConnectionOpenAck(opts, msg)
		case *conntypes.MsgConnectionOpenConfirm:
			tx, err = c.txConnectionOpenConfirm(opts, msg)
		case *chantypes.MsgChannelOpenInit:
			tx, err = c.txChannelOpenInit(opts, msg)
		case *chantypes.MsgChannelOpenTry:
			tx, err = c.txChannelOpenTry(opts, msg)
		case *chantypes.MsgChannelOpenAck:
			tx, err = c.txChannelOpenAck(opts, msg)
		case *chantypes.MsgChannelOpenConfirm:
			tx, err = c.txChannelOpenConfirm(opts, msg)
		case *chantypes.MsgRecvPacket:
			tx, err = c.txRecvPacket(opts, msg)
		case *chantypes.MsgAcknowledgement:
			tx, err = c.txAcknowledgement(opts, msg)
		default:
			panic("illegal msg type")
		}
		if err != nil {
			return nil, err
		}
		if _, err := c.waitForReceiptAndGet(ctx, tx); err != nil {
			return nil, err
		}
		if c.msgEventListener != nil {
			if err := c.msgEventListener.OnSentMsg([]sdk.Msg{msg}); err != nil {
				log.Println("failed to OnSendMsg call", "msg", msg, "err", err)
			}
		}
	}
	return nil, nil
}

// Send sends msgs to the chain and logging a result of it
// It returns a boolean value whether the result is success
func (c *Chain) Send(msgs []sdk.Msg) bool {
	_, err := c.SendMsgs(msgs)
	return err == nil
}

func (c *Chain) waitForReceiptAndGet(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	var receipt *types.Receipt
	err := retry.Do(
		func() error {
			r, err := c.ethClient.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				if err != ethereum.NotFound {
					err = retry.Unrecoverable(err)
				}
				return err
			}
			receipt = r
			return nil
		},
	)
	return receipt, err
}

func (c *Chain) txCreateClient(opts *bind.TransactOpts, msg *clienttypes.MsgCreateClient) (*types.Transaction, error) {
	var clientState exported.ClientState
	if err := c.codec.UnpackAny(msg.ClientState, &clientState); err != nil {
		return nil, err
	}
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	consensusStateBytes, err := proto.Marshal(msg.ConsensusState)
	if err != nil {
		return nil, err
	}
	return c.ibcHandler.CreateClient(opts, ibchandler.IBCMsgsMsgCreateClient{
		ClientType:          clientState.ClientType(),
		Height:              pbToHandlerHeight(clientState.GetLatestHeight()),
		ClientStateBytes:    clientStateBytes,
		ConsensusStateBytes: consensusStateBytes,
	})
}

func (c *Chain) txUpdateClient(opts *bind.TransactOpts, msg *clienttypes.MsgUpdateClient) (*types.Transaction, error) {
	headerBytes, err := proto.Marshal(msg.Header)
	if err != nil {
		return nil, err
	}
	return c.ibcHandler.UpdateClient(opts, ibchandler.IBCMsgsMsgUpdateClient{
		ClientId:      msg.ClientId,
		ClientMessage: headerBytes,
	})
}

func (c *Chain) txConnectionOpenInit(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenInit) (*types.Transaction, error) {
	return c.ibcHandler.ConnectionOpenInit(opts, ibchandler.IBCMsgsMsgConnectionOpenInit{
		ClientId: msg.ClientId,
		Counterparty: ibchandler.CounterpartyData{
			ClientId:     msg.Counterparty.ClientId,
			ConnectionId: msg.Counterparty.ConnectionId,
			Prefix:       ibchandler.MerklePrefixData(msg.Counterparty.Prefix),
		},
		DelayPeriod: msg.DelayPeriod,
	})
}

func (c *Chain) txConnectionOpenTry(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenTry) (*types.Transaction, error) {
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	var versions []ibchandler.VersionData
	for _, v := range msg.CounterpartyVersions {
		versions = append(versions, ibchandler.VersionData(*v))
	}
	return c.ibcHandler.ConnectionOpenTry(opts, ibchandler.IBCMsgsMsgConnectionOpenTry{
		PreviousConnectionId: msg.PreviousConnectionId,
		Counterparty: ibchandler.CounterpartyData{
			ClientId:     msg.Counterparty.ClientId,
			ConnectionId: msg.Counterparty.ConnectionId,
			Prefix:       ibchandler.MerklePrefixData(msg.Counterparty.Prefix),
		},
		DelayPeriod:          msg.DelayPeriod,
		ClientId:             msg.ClientId,
		ClientStateBytes:     clientStateBytes,
		CounterpartyVersions: versions,
		ProofInit:            msg.ProofInit,
		ProofClient:          msg.ProofClient,
		ProofConsensus:       msg.ProofConsensus,
		ProofHeight:          pbToHandlerHeight(msg.ProofHeight),
		ConsensusHeight:      pbToHandlerHeight(msg.ConsensusHeight),
	})
}

func (c *Chain) txConnectionOpenAck(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenAck) (*types.Transaction, error) {
	clientStateBytes, err := proto.Marshal(msg.ClientState)
	if err != nil {
		return nil, err
	}
	return c.ibcHandler.ConnectionOpenAck(opts, ibchandler.IBCMsgsMsgConnectionOpenAck{
		ConnectionId:     msg.ConnectionId,
		ClientStateBytes: clientStateBytes,
		Version: ibchandler.VersionData{
			Identifier: msg.Version.Identifier,
			Features:   msg.Version.Features,
		},
		CounterpartyConnectionID: msg.CounterpartyConnectionId,
		ProofTry:                 msg.ProofTry,
		ProofClient:              msg.ProofClient,
		ProofConsensus:           msg.ProofConsensus,
		ProofHeight:              pbToHandlerHeight(msg.ProofHeight),
		ConsensusHeight:          pbToHandlerHeight(msg.ConsensusHeight),
	})
}

func (c *Chain) txConnectionOpenConfirm(opts *bind.TransactOpts, msg *conntypes.MsgConnectionOpenConfirm) (*types.Transaction, error) {
	return c.ibcHandler.ConnectionOpenConfirm(opts, ibchandler.IBCMsgsMsgConnectionOpenConfirm{
		ConnectionId: msg.ConnectionId,
		ProofAck:     msg.ProofAck,
		ProofHeight:  pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) txChannelOpenInit(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenInit) (*types.Transaction, error) {
	return c.ibcHandler.ChannelOpenInit(opts, ibchandler.IBCMsgsMsgChannelOpenInit{
		PortId: msg.PortId,
		Channel: ibchandler.ChannelData{
			State:          uint8(msg.Channel.State),
			Ordering:       uint8(msg.Channel.Ordering),
			Counterparty:   ibchandler.ChannelCounterpartyData(msg.Channel.Counterparty),
			ConnectionHops: msg.Channel.ConnectionHops,
			Version:        msg.Channel.Version,
		},
	})
}

func (c *Chain) txChannelOpenTry(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenTry) (*types.Transaction, error) {
	return c.ibcHandler.ChannelOpenTry(opts, ibchandler.IBCMsgsMsgChannelOpenTry{
		PortId:            msg.PortId,
		PreviousChannelId: msg.PreviousChannelId,
		Channel: ibchandler.ChannelData{
			State:          uint8(msg.Channel.State),
			Ordering:       uint8(msg.Channel.Ordering),
			Counterparty:   ibchandler.ChannelCounterpartyData(msg.Channel.Counterparty),
			ConnectionHops: msg.Channel.ConnectionHops,
			Version:        msg.Channel.Version,
		},
		CounterpartyVersion: msg.CounterpartyVersion,
		ProofInit:           msg.ProofInit,
		ProofHeight:         pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) txChannelOpenAck(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenAck) (*types.Transaction, error) {
	return c.ibcHandler.ChannelOpenAck(opts, ibchandler.IBCMsgsMsgChannelOpenAck{
		PortId:                msg.PortId,
		ChannelId:             msg.ChannelId,
		CounterpartyVersion:   msg.CounterpartyVersion,
		CounterpartyChannelId: msg.CounterpartyChannelId,
		ProofTry:              msg.ProofTry,
		ProofHeight:           pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) txChannelOpenConfirm(opts *bind.TransactOpts, msg *chantypes.MsgChannelOpenConfirm) (*types.Transaction, error) {
	return c.ibcHandler.ChannelOpenConfirm(opts, ibchandler.IBCMsgsMsgChannelOpenConfirm{
		PortId:      msg.PortId,
		ChannelId:   msg.ChannelId,
		ProofAck:    msg.ProofAck,
		ProofHeight: pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) txRecvPacket(opts *bind.TransactOpts, msg *chantypes.MsgRecvPacket) (*types.Transaction, error) {
	return c.ibcHandler.RecvPacket(opts, ibchandler.IBCMsgsMsgPacketRecv{
		Packet: ibchandler.PacketData{
			Sequence:           msg.Packet.Sequence,
			SourcePort:         msg.Packet.SourcePort,
			SourceChannel:      msg.Packet.SourceChannel,
			DestinationPort:    msg.Packet.DestinationPort,
			DestinationChannel: msg.Packet.DestinationChannel,
			Data:               msg.Packet.Data,
			TimeoutHeight:      ibchandler.HeightData(msg.Packet.TimeoutHeight),
			TimeoutTimestamp:   msg.Packet.TimeoutTimestamp,
		},
		Proof:       msg.ProofCommitment,
		ProofHeight: pbToHandlerHeight(msg.ProofHeight),
	})
}

func (c *Chain) txAcknowledgement(opts *bind.TransactOpts, msg *chantypes.MsgAcknowledgement) (*types.Transaction, error) {
	return c.ibcHandler.AcknowledgePacket(opts, ibchandler.IBCMsgsMsgPacketAcknowledgement{
		Packet: ibchandler.PacketData{
			Sequence:           msg.Packet.Sequence,
			SourcePort:         msg.Packet.SourcePort,
			SourceChannel:      msg.Packet.SourceChannel,
			DestinationPort:    msg.Packet.DestinationPort,
			DestinationChannel: msg.Packet.DestinationChannel,
			Data:               msg.Packet.Data,
			TimeoutHeight:      ibchandler.HeightData(msg.Packet.TimeoutHeight),
			TimeoutTimestamp:   msg.Packet.TimeoutTimestamp,
		},
		Acknowledgement: msg.Acknowledgement,
		Proof:           msg.ProofAcked,
		ProofHeight:     pbToHandlerHeight(msg.ProofHeight),
	})
}
