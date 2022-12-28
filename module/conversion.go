package module

import (
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
)

func ethHeightToPB(height int64) clienttypes.Height {
	return clienttypes.NewHeight(0, uint64(height))

}
