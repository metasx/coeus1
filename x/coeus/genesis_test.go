package coeus_test

import (
	"testing"

	keepertest "coeus/testutil/keeper"
	"coeus/testutil/nullify"
	"coeus/x/coeus"
	"coeus/x/coeus/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoeusKeeper(t)
	coeus.InitGenesis(ctx, *k, genesisState)
	got := coeus.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
