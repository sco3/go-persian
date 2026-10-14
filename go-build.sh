
export GOPATH=$GOPATH:$(readlink -f .):$(readlink -f $HOME/go)

#-c -v -gcflags "-N -l"

cd src/main
go build -o ../../converter
cd -
