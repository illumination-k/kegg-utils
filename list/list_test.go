package list_test

import (
	"kegg/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadList(t *testing.T) {
	example := `ko:K00001	E1.1.1.1, adh; alcohol dehydrogenase [EC:1.1.1.1]
	ko:K00002	AKR1A1, adh; alcohol dehydrogenase (NADP+) [EC:1.1.1.2]
	ko:K00003	hom; homoserine dehydrogenase [EC:1.1.1.3]
	ko:K00004	BDH, butB; (R,R)-butanediol dehydrogenase / meso-butanediol dehydrogenase / diacetyl reductase [EC:1.1.1.4 1.1.1.- 1.1.1.303]
	ko:K00005	gldA; glycerol dehydrogenase [EC:1.1.1.6]
	ko:K00006	GPD1; glycerol-3-phosphate dehydrogenase (NAD+) [EC:1.1.1.8]
	ko:K00007	dalD; D-arabinitol 4-dehydrogenase [EC:1.1.1.11]
	ko:K00008	SORD, gutB; L-iditol 2-dehydrogenase [EC:1.1.1.14]`

	t.Run("Test read list", func(t *testing.T) {
		idToName, err := list.ReadList(example)
		if err != nil {
			t.Error(idToName)
		}

		name, found := idToName["ko:K00001"]
		if !found {
			t.Error("ko:k00001 should be found")
		}

		assert.Equal(t, name, "E1.1.1.1, adh; alcohol dehydrogenase [EC:1.1.1.1]")
	})

}
