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

```text

cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkConvert
2026/10/14 19:51:32 processed: 100000 lines.
2026/10/14 19:51:33 processed: 200000 lines.
2026/10/14 19:51:34 processed: 300000 lines.
2026/10/14 19:51:35 processed: 400000 lines.
2026/10/14 19:51:36 processed: 500000 lines.
2026/10/14 19:51:37 processed: 600000 lines.
2026/10/14 19:51:38 processed: 700000 lines.
2026/10/14 19:51:39 processed: 800000 lines.
2026/10/14 19:51:39 processed: 900000 lines.
2026/10/14 19:51:40 processed: 1000000 lines.
2026/10/14 19:51:41 time: 9.178962777s
2026/10/14 19:51:41 conv: 6.78
2026/10/14 19:51:41 expected crc: 95646885 got crc: 95646885
2026/10/14 19:51:41 Speed: 112491 records/second

```

The linter was built from git repo with enforced go 1.27.2:

```bash
golangci-lint --version
golangci-lint has version 2.14.1-0.20261013143752-f4e2329e669f+dirty built with go1.27.2 from (f4e2329e669f818fe3cbd3181a3e116e5aa24db7, modified: true, mod sum: "") on 2026-10-13T14:37:52Z

```
