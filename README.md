# Persian-Go

A high-performance Go-to-C shared library for rapid CSV-to-Avro conversion using zero-copy memory mapping (`CGO` and JNI) to interface with Java.

## Quick Start

1. **Build the shared library**:
```bash
   ./build-persian-lib.sh
```

2. **Run benchmarks**:
Run this in project root dir:
```bash
   zstd -d -k data/000000_0.csv.zst -o /tmp/000000_0.csv
   go test -v ./persian -bench=.
```

One can run `go-test.sh` instead:

Test results:
```
cpu: AMD Ryzen 5 4500U with Radeon Graphics         
BenchmarkConvert
2026/10/14 10:56:48 processed: 100000 lines.
2026/10/14 10:56:49 processed: 200000 lines.
2026/10/14 10:56:51 processed: 300000 lines.
2026/10/14 10:56:52 processed: 400000 lines.
2026/10/14 10:56:53 processed: 500000 lines.
2026/10/14 10:56:55 processed: 600000 lines.
2026/10/14 10:56:56 processed: 700000 lines.
2026/10/14 10:56:57 processed: 800000 lines.
2026/10/14 10:56:59 processed: 900000 lines.
2026/10/14 10:57:00 processed: 1000000 lines.
2026/10/14 10:57:00 time: 13.191800408s
2026/10/14 10:57:00 conv: 9.81
2026/10/14 10:57:00 expected crc: 95646885 got crc: 95646885
2026/10/14 10:57:00 Speed: 77878 record/second

```