package types

import (
	// standard library imports
	"testing"

	// external library imports
	"github.com/stretchr/testify/require"
)

func TestValidateParams(t *testing.T) {
	require.NoError(t, DefaultParams().Validate())
	require.NoError(t, NewParams(true, false).Validate())
}
