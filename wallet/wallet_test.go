package wallet_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/simapp"
	simappparams "cosmossdk.io/simapp/params"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/cosmos-go-wallet/client"
	"github.com/desmos-labs/cosmos-go-wallet/types"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
)

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

type WalletTestSuite struct {
	suite.Suite

	wallet *wallet.Wallet
	client *client.Client
}

// makeEncodingConfig returns the encoding config to be used for the tests
// Note: This is copied from the simapp package since we cannot import it directly
func (suite *WalletTestSuite) makeEncodingConfig() simappparams.EncodingConfig {
	encodingConfig := simappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	simapp.ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	simapp.ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

func (suite *WalletTestSuite) SetupSuite() {
	chainCfg := types.ChainConfig{
		Bech32Prefix:  "desmos",
		RPCAddr:       "https://rpc.morpheus.desmos.network:443",
		GRPCAddr:      "https://grpc.morpheus.desmos.network:443",
		GasPrice:      "0.01udaric",
		GasAdjustment: 1.5,
	}
	accountCfg := types.AccountConfig{
		Mnemonic: "forward service profit benefit punch catch fan chief jealous steel harvest column spell rude warm home melody hat broccoli pulse say garlic you firm",
		HDPath:   "m/44'/852'/0'/0/0",
	}

	// Set up the SDK config with the proper bech32 prefixes
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(chainCfg.Bech32Prefix, fmt.Sprintf("%spub", chainCfg.Bech32Prefix))

	encodingCfg := suite.makeEncodingConfig()

	c, err := client.NewClient(&chainCfg, encodingCfg.Codec)
	suite.Require().NoError(err)
	suite.client = c

	w, err := wallet.NewWallet(&accountCfg, c, encodingCfg.TxConfig)
	suite.Require().NoError(err)
	suite.wallet = w
}

func (suite *WalletTestSuite) TestBuildTx() {
	testCases := []struct {
		name      string
		msgs      []sdk.Msg
		shouldErr bool
		check     func(builder sdkclient.TxBuilder)
	}{
		{
			name:      "empty messages returns error",
			shouldErr: true,
		},
		{
			name: "valid messages returns no error",
			msgs: []sdk.Msg{
				banktypes.NewMsgSend(
					sdk.MustAccAddressFromBech32(suite.wallet.AccAddress()),
					sdk.MustAccAddressFromBech32("desmos1q62k9kvjy7v2wh0yt9jqaepnzezz3s49j9gnpk"),
					sdk.NewCoins(sdk.NewCoin("udaric", sdk.NewInt(10000))),
				),
			},
			check: func(builder sdkclient.TxBuilder) {
				tx := builder.GetTx()
				suite.Require().Lessf(tx.GetGas(), uint64(200_000), "MsgSend should take less than 200.000 gas")
				suite.Require().NotEmptyf(tx.GetFee(), "Fees should not be empty")
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			data := types.NewTransactionData(
				tc.msgs...,
			).WithGasAuto().WithFeeAuto().WithMemo("Custom memo").WithSequence(0)

			builder, err := suite.wallet.BuildTx(data)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			if tc.check != nil {
				tc.check(builder)
			}
		})
	}
}
