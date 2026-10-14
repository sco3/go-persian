export GOPATH=${GOPATH}:$(readlink -f .):$(readlink -f "${HOME}"/go)

cd src/main || exit
go build -compiler gccgo -gccgoflags "-march=native -O3" -o ../../converter-gcc
cd - || exit
