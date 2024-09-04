package sorts

import (
	"bufio"
	"fmt"
	"github.com/dfwcnj/govbinsort/types"
	"github.com/dfwcnj/randomdata"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_savemergefile(t *testing.T) {
	var l int = 32
	var r bool = true
	var e bool = false
	var nrs int64 = 1 << 20
	var dlim string
	dlim = "\n"

	log.Print("savemergefile test")
	dn, err := initmergedir("/tmp", "somesort")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dn)

	for i := range 10 {
		var klns kvallines
		var kln kvalline

		rsl := randomdata.Randomstrings(nrs, l, r, e)
		for _, s := range rsl {
			bln := []byte(s)
			kln.line = bln
			kln.key = kln.line
			klns = append(klns, kln)
		}
		if len(klns) != int(nrs) {
			//log.Print(klns)
			log.Fatal("klns: before sort wanted len ", l, " got ", len(klns))
		}

		slns := Klslicessort(klns, 0)
		var fn = filepath.Join(dn, fmt.Sprint("file", i))
		savemergefile(slns, fn, dlim)

		fp, err := os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		var rlns []string
		for scanner.Scan() {
			l := scanner.Text()
			if len(l) == 0 {
				continue
			}
			var sep = make([]byte, 1)
			sa := strings.Split(l, string(sep))
			if len(sa) != 2 {
				log.Fatal("split ", l, " wanted ", 2, " got ", len(sa))
			}
			rlns = append(rlns, sa[1])
		}
		if len(rlns) != int(nrs) {
			log.Fatal("rlns wanted ", nrs, " got ", len(rlns))
		}
	}
	log.Print("savemergefile test passed")
}
