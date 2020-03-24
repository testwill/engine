package service

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mesg-foundation/engine/x/service/internal/types"
)

// NewHandler creates an sdk.Handler for all the service type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreate:
			return handleMsgCreate(ctx, k, &msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

// handleMsgCreate creates a new process.
func handleMsgCreate(ctx sdk.Context, k Keeper, msg *MsgCreate) (*sdk.Result, error) {
	s, err := k.Create(ctx, msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreated),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeHash, s.Hash.String()),
		),
	)

	return &sdk.Result{
		Data:   s.Hash,
		Events: ctx.EventManager().Events(),
	}, nil
}
