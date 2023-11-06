package eastmoney

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueryHistoricalFinaMainData(t *testing.T) {
	_, err := _em.QueryHistoricalFinaMainData(_ctx, "600188.SH")
	require.Nil(t, err)
}
