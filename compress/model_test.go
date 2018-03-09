package compress

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewModel(t *testing.T) {
	require.NoError(t,NewModel("location.js","compress","data.go"))
}
