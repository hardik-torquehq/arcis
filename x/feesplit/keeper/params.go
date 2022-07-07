package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Ambiplatforms-TORQUE/arcis/v6/x/feesplit/types"
)

// GetParams returns the total set of fees parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the fees parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
