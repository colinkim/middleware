package types

import (
	"testing"

	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"middleware/testutil/sample"
)

func TestMsgAnchoring_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAnchoring
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAnchoring{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAnchoring{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
