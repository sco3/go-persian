# Persian-Go

A high-performance Go-to-C shared library for rapid CSV-to-Avro conversion using zero-copy memory mapping (`CGO` and JNI) to interface with Java.

## Quick Start

1. **Build the shared library**:
```bash
   ./build-persian-lib.sh
```

2. **Run benchmarks**:
```bash
   go test -v ./persian -bench=.
```