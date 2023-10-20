package keeper_test

import (
	"testing"

	testkeeper "coeus/testutil/keeper"
	"coeus/x/coeus/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CoeusKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
