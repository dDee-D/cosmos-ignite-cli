package dex_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/interchange/testutil/keeper"
	"github.com/username/interchange/testutil/nullify"
	"github.com/username/interchange/x/dex"
	"github.com/username/interchange/x/dex/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
