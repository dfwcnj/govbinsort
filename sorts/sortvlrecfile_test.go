package sorts

import (
	"fmt"
	"github.com/dfwcnj/govbinsort/merge"
	"github.com/dfwcnj/randomdata"
	"log"
	"os"
	"path"
	"slices"
	"testing"
)

func Test_sortvlrecfile(t *testing.T) {
	var l int = 32
	var r bool = true
	var e bool = false
	var nrs int64 = 1 << 20
	var nss int64
	var iomem int64 = 1<<24 + 1<<20

	//var klns kvallines
	var tklns kvallines
	var err error
	var nr int

	dn, err := merge.Initmergedir("/tmp", "somesort")

	log.Println("sortvlrecfile test")

	rsl := randomdata.Randomstrings(nrs, l, r, e)

	fn := path.Join(dn, "sortvlrecfiletest")
	fp, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i, _ := range rsl {
		fmt.Fprintln(fp, rsl[i])
		nr++
	}
	fp.Close()

	//log.Printf("sortvlrecfile test sortvlrecfile %s, %d\n", fn, iomem)
	_, fns, err := sortvlrecfile(fn, dn, "std", int(l)+1, 0, 0, iomem)

	//log.Println("sortvlrecfile test after  klns ", len(klns))
	//log.Println("sortvlrecfile test after fns ", fns)

	for _, f := range fns {
		//log.Println("sortvlrecfile chacking ", f)
		mfp, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		tklns, _, err = Vlreadn(mfp, 0, 0, 0, iomem*2)
		//log.Println("sortvlrecfile test tklns ", len(tklns))

		var lns = make([]string, 0)
		for _, t := range tklns {
			lns = append(lns, string(t.line))
		}
		if slices.IsSorted(lns) == false {
			log.Fatal(f, " is not sorted")
		}
		nss += int64(len(tklns))
	}
	if nrs != nss {
		log.Fatal("sortvlrecfile test wanted ", nrs, " got ", nss)
	}
	log.Println("sortvlrecfile passed")
}
