package sina

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeywordSearch(t *testing.T) {
	results, err := _s.KeywordSearch(_ctx, "五粮液")
	require.Nil(t, err)
	t.Log(results)
	fmt.Printf("\n%+v", results)
}
