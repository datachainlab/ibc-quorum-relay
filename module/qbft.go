package module

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v4/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

func (cs *ClientState) ClientType() string {
	panic("not implemented")
}

func (cs *ClientState) GetLatestHeight() exported.Height {
	return clienttypes.NewHeight(0, uint64(cs.LatestHeight))
}

func (cs *ClientState) Validate() error {
	panic("not implemented")
}

func (cs *ClientState) Initialize(sdk.Context, codec.BinaryCodec, sdk.KVStore, exported.ConsensusState) error {
	panic("not implemented")
}

func (cs *ClientState) Status(ctx sdk.Context, clientStore sdk.KVStore, cdc codec.BinaryCodec) exported.Status {
	panic("not implemented")
}

func (cs *ClientState) ExportMetadata(sdk.KVStore) []exported.GenesisMetadata {
	panic("not implemented")
}

func (cs *ClientState) CheckHeaderAndUpdateState(sdk.Context, codec.BinaryCodec, sdk.KVStore, exported.Header) (exported.ClientState, exported.ConsensusState, error) {
	panic("not implemented")
}

func (cs *ClientState) CheckMisbehaviourAndUpdateState(sdk.Context, codec.BinaryCodec, sdk.KVStore, exported.Misbehaviour) (exported.ClientState, error) {
	panic("not implemented")
}

func (cs *ClientState) CheckSubstituteAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore, substituteClientStore sdk.KVStore, substituteClient exported.ClientState) (exported.ClientState, error) {
	panic("not implemented")
}

func (cs *ClientState) VerifyUpgradeAndUpdateState(
	ctx sdk.Context,
	cdc codec.BinaryCodec,
	store sdk.KVStore,
	newClient exported.ClientState,
	newConsState exported.ConsensusState,
	proofUpgradeClient,
	proofUpgradeConsState []byte,
) (exported.ClientState, exported.ConsensusState, error) {
	panic("not implemented")
}

func (cs *ClientState) ZeroCustomFields() exported.ClientState {
	panic("not implemented")
}

func (cs *ClientState) VerifyClientState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	counterpartyClientIdentifier string,
	proof []byte,
	clientState exported.ClientState,
) error {
	panic("not implemented")
}

func (cs *ClientState) VerifyClientConsensusState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	counterpartyClientIdentifier string,
	consensusHeight exported.Height,
	prefix exported.Prefix,
	proof []byte,
	consensusState exported.ConsensusState,
) error {
	panic("not implemented")
}

func (cs *ClientState) VerifyConnectionState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
	connectionID string,
	connectionEnd exported.ConnectionI,
) error {
	panic("not implemented")
}

func (cs *ClientState) VerifyChannelState(
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	channel exported.ChannelI,
) error {
	panic("not implemented")
}

func (cs *ClientState) VerifyPacketCommitment(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
	commitmentBytes []byte,
) error {
	panic("not implemented")
}

func (cs *ClientState) VerifyPacketAcknowledgement(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
	acknowledgement []byte,
) error {
	panic("not implemented")
}

func (cs *ClientState) VerifyPacketReceiptAbsence(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
) error {
	panic("not implemented")
}

func (cs *ClientState) VerifyNextSequenceRecv(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	prefix exported.Prefix,
	proof []byte,
	portID,
	channelID string,
	nextSequenceRecv uint64,
) error {
	panic("not implemented")
}
