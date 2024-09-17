package sorts

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func filelinecount(fn string) int64 {
	cmd := exec.Command("wc", "-l", fn)
	defer cmd.Wait()
	ofp, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("sortfiles test filelinecount pipe", err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal("sortfiles test filelinecount start", err)
	}
	r, err := io.ReadAll(ofp)
	if err != nil {
		log.Fatal("sortfiles test filelinecount read ", err)
	}
	rsl := strings.Split(string(r), " ")
	//log.Print(rsl, " ", len(rsl), " ", rsl[len(rsl)-1])
	i, err := strconv.ParseInt(rsl[len(rsl)-2], 10, 64)
	if err != nil {
		log.Fatal("sortfiles test filelinecount parse ", err)
	}
	return i
}

func filereccount(fn string, rlen int) int64 {
	fp, err := os.Open(fn)
	if err != nil {
		log.Fatal("filereccount ", err)
	}
	finf, err := fp.Stat()
	if err != nil {
		log.Fatal("filereccount ", err)
	}
	return finf.Size() / int64(rlen)
}
