set -euo pipefail

zstd -d -k data/000000_0.csv.zst -o /tmp/000000_0.csv

go test -v ./persian -bench=.
