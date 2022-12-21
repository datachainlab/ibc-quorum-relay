package module

import (
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	committypes "github.com/cosmos/ibc-go/v4/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
)

func ethHeightToPB(height int64) clienttypes.Height {
	return clienttypes.NewHeight(0, uint64(height))

}

func connectionEndToPB(conn ibchost.ConnectionEndData) conntypes.ConnectionEnd {
	connpb := conntypes.ConnectionEnd{
		ClientId:    conn.ClientId,
		Versions:    []*conntypes.Version{},
		State:       conntypes.State(conn.State),
		DelayPeriod: conn.DelayPeriod,
		Counterparty: conntypes.Counterparty{
			ClientId:     conn.Counterparty.ClientId,
			ConnectionId: conn.Counterparty.ConnectionId,
			Prefix:       committypes.MerklePrefix(conn.Counterparty.Prefix),
		},
	}
	for _, v := range conn.Versions {
		ver := conntypes.Version(v)
		connpb.Versions = append(connpb.Versions, &ver)
	}
	return connpb
}

func channelToPB(chann ibchost.ChannelData) chantypes.Channel {
	return chantypes.Channel{
		State:          chantypes.State(chann.State),
		Ordering:       chantypes.Order(chann.Ordering),
		Counterparty:   chantypes.Counterparty(chann.Counterparty),
		ConnectionHops: chann.ConnectionHops,
		Version:        chann.Version,
	}
}

func pbToHostHeight(height exported.Height) ibchost.HeightData {
	return ibchost.HeightData{
		RevisionNumber: height.GetRevisionNumber(),
		RevisionHeight: height.GetRevisionHeight(),
	}
}

func pbToHandlerHeight(height exported.Height) ibchandler.HeightData {
	return ibchandler.HeightData{
		RevisionNumber: height.GetRevisionNumber(),
		RevisionHeight: height.GetRevisionHeight(),
	}
}
