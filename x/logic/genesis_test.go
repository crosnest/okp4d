package logic_test

import (
	"testing"

	keepertest "github.com/okp4/okp4d/testutil/keeper"
	"github.com/okp4/okp4d/testutil/nullify"
	"github.com/okp4/okp4d/x/logic"
	"github.com/okp4/okp4d/x/logic/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.LogicKeeper(t)
	logic.InitGenesis(ctx, *k, genesisState)
	got := logic.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
