package sorts

import (
	"log"
	"slices"
	"testing"

	"github.com/dfwcnj/randomdata"
)

func Test_kvinsertionsort(t *testing.T) {

	//ls := []int{1 << 3, 1 << 4, 1 << 5, 1 << 6}
	ls := []int{1 << 5}
	//ns := []int64{1 << 3, 1 << 16, 1 << 20}
	ns := []int64{1 << 16}

	for _, ll := range ls {
		for _, nl := range ns {

			var l int = ll
			var r bool = false
			var e bool = false
			var keyoff = 0
			var reclen = ll
			var keylen = ll
			log.Print("testing kvinsertionsort of ", nl, " random strings length ", l)
			rsl := randomdata.Randomstrings(nl, l, r, e)
			if len(rsl) != int(nl) {
				t.Fatal("kvinsertionsort test rsl: wanted len ", nl, " got ", len(rsl))
			}
			lns := make([][]byte, 0, nl)
			for _, s := range rsl {
				lns = append(lns, []byte(s))
			}
			kvinsertionsort(lns, reclen, keyoff, keylen)
			ssl := make([]string, 0, nl)
			for _, bs := range lns {
				ssl = append(ssl, string(bs))
			}
			if !slices.IsSorted(ssl) {
				t.Fatal("kvinsertionsort test failed not sorted")
			} else {
				log.Print("kvinsertionsort test passed for ", nl)
			}

		}
	}
}
