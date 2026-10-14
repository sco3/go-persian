package persian

import (
	"log"
	"os"
	//"os/exec"
	"testing"
)

func BenchmarkConvert(t *testing.B) {
	outfile := "/tmp/out-go.avro"
	_, err := os.Stat(outfile)
	if err == nil {
		err := os.Remove(outfile)
		if err != nil {
			panic(err)
		}
	}

	crc := RunConverter()

	log.Printf("crc: %x\n", crc)
	if 0x95646885 != crc {
		t.Fail()
	}

	_, err = os.Stat(outfile)
	if err != nil {
		t.Fail()
	}

	//	out, err := exec.Command("/usr/bin/crc32", outfile).Output()
	//	if err != nil {
	//		panic(err)
	//	}
	//	outstr := string(out)
	//	println("crc:", outstr[:8])
	//	if "95646885" != outstr[:8] {
	//		t.Fail()
	//	}

}
