package sample

import (
	"github.com/reapchain/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/reapchain/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}
