package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	foomodule "github.com/username/test/x/foo"
)

// App modules are defined as NewBasicManager arguments
var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	bank.AppModuleBasic{},
	staking.AppModuleBasic{},
	gov.NewAppModuleBasic([]govclient.ProposalHandler{
		paramsclient.ProposalHandler,
	}),
	foomodule.AppModuleBasic{},
)

type Foo struct{}

func (Foo) Name() string {
	return "foo"
}

func (Foo) GetKey(storeKey string) *storetypes.KVStoreKey { return nil }

func (Foo) TxConfig() client.TxConfig { return nil }

func (Foo) BeginBlocker(sdk.Context, abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return abci.ResponseBeginBlock{}
}

func (Foo) EndBlocker(sdk.Context, abci.RequestEndBlock) abci.ResponseEndBlock {
	return abci.ResponseEndBlock{}
}

func (Foo) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	_ = apiSvr.ClientCtx
}
