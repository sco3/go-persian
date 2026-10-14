package persian

import (
	"bufio"
	"hash/crc32"
	"log"
	"os"
	"time"
)

// RunConverter test converter
func RunConverter() uint32 {
	out := AllocMemBuffer(10000)
	out.Buf[0] = 0
	out.Buf[1] = 0
	out.Buf[2] = 0
	out.Buf[3] = 0
	out.Buf[4] = 2

	start := time.Now()

	file, err := os.Open("/tmp/000000_0.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	outfile, err := os.Create("/tmp/out-go.avro")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	scanner := bufio.NewScanner(file)

	schema := []byte("sssssslsssssssllllssllllllllllddddddddddlsslsssssssssssdllssssssllsssdddssssssssssldssssssslssd")

	var nanos time.Duration
	cnt := 0
	hash := crc32.NewIEEE()

	for scanner.Scan() {
		var line = scanner.Bytes()
		out.Offset = 5
		start := time.Now()
		flds := new(FldInfo)
		Convert(
			schema,
			line,
			out,
			flds,
		)
		nanos += time.Since(start)
		buf := out.Buf[:out.Offset]
		hash.Write(buf)
		outfile.Write(buf)

		cnt++

		if cnt%100000 == 0 {
			log.Printf("processed: %d lines.\n", cnt)
		}
	}
	took := time.Since(start)
	log.Printf("time: %s\n", took)
	log.Printf("conv: %.2f\n", float64(nanos)/1000000000.0)
	crc := hash.Sum32()
	log.Printf("expected crc: 95646885 got crc: %x\n", crc)
	log.Printf("Speed: %v records/second\n", cnt/int(took.Seconds()))
	return crc
}
