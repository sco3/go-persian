export GOPATH=$GOPATH:$(readlink -f .):$(readlink -f $HOME/go)


cd src/main
go build -compiler gccgo -gccgoflags "-march=native -O3"  -o ../../converter-gcc
cd -
