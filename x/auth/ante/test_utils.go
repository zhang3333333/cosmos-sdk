package ante

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	dbm "github.com/tendermint/tendermint/libs/db"
)

func setupMultiStore() (sdk.MultiStore, *sdk.KVStoreKey, *sdk.KVStoreKey) {
	db := dbm.NewMemDB()
	capKey := sdk.NewKVStoreKey("capkey")
	capKey2 := sdk.NewKVStoreKey("capkey2")
	ms := store.NewCommitMultiStore(db)

	ms.MountStoreWithDB(capKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(capKey2, sdk.StoreTypeIAVL, db)
	ms.LoadLatestVersion()

	return ms, capKey, capKey2
}

func newTestMsg(addrs ...sdk.AccAddress) *sdk.TestMsg {
	return sdk.NewTestMsg(addrs...)
}

func newStdFee() auth.StdFee {
	return auth.NewStdFee(5000,
		sdk.NewInt64Coin("atom", 150),
	)
}

// coins to more than cover the fee
func newCoins() sdk.Coins {
	return sdk.Coins{
		sdk.NewInt64Coin("atom", 10000000),
	}
}

// generate a priv key and return it with its address
func privAndAddr() (crypto.PrivKey, sdk.AccAddress) {
	priv := ed25519.GenPrivKey()
	addr := sdk.AccAddress(priv.PubKey().Address())
	return priv, addr
}
