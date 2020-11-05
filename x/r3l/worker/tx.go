package worker

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	"github.com/cosmos/cosmos-sdk/client"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/relevant-community/r3l/x/r3l/keeper"
	"github.com/relevant-community/r3l/x/r3l/types"
)

func SetScores(ctx sdk.Context, k keeper.Keeper, scores []types.Score) (*sdk.TxResponse, error) {
	kb, err := keyring.New(sdk.KeyringServiceName(), "test", "~/r3ld", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get address from Keybase: %w", err)
	}

	acc, err := kb.Key("user1")
	if err != nil {
		return nil, err
	}

	accAddress := acc.GetAddress()
	fmt.Println("From address", accAddress)

	// TODO the acc seq might be out of date, we should query rest endpoint for current seq
	account := k.AccountKeeper.GetAccount(ctx, accAddress)

	interfaceRegistry := codecTypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := authTx.NewTxConfig(marshaler, authTx.DefaultSignModes)
	clientCtx := client.Context{}.
		WithInterfaceRegistry(interfaceRegistry).
		WithNodeURI("http://0.0.0.0:26657").
		WithTxConfig(txCfg).
		WithFromAddress(accAddress).
		WithFromName("user1")

	txf := tx.Factory{}.
		WithTxConfig(txCfg).
		WithAccountRetriever(authTypes.AccountRetriever{}).
		WithChainID(ctx.ChainID()).
		WithAccountNumber(account.GetAccountNumber()).
		WithSequence(account.GetSequence()).
		WithKeybase(kb).
		WithGas(200000).
		WithGasPrices("0.0token")

	fmt.Println("seq", account.GetSequence(), account.GetAccountNumber())

	msg := types.NewMsgScores(accAddress, ctx.BlockHeight(), scores)

	// TODO - we should use AccountRetriever to make sure we are using an up-todate sequence
	// We could use this util method but AccountRetriever is causing issues because
	// interfaceRegistry doesnt have AccountI
	// err = tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)

	txBuilder, err := tx.BuildUnsignedTx(txf, msg)
	err = tx.Sign(txf, clientCtx.GetFromName(), txBuilder)
	if err != nil {
		fmt.Println("err", err)
		return nil, fmt.Errorf("Failed to Generate and broadcast TX: %w", err)
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	// broadcast to a Tendermint node
	res, err := clientCtx.BroadcastTxSync(txBytes)
	if err != nil {
		fmt.Println("err", err)
		return nil, fmt.Errorf("Failed to Generate and broadcast TX: %w", err)
	}
	fmt.Println(res)
	return res, nil
}
