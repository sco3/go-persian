export GOPATH=${GOPATH}:$(readlink -f .):$(readlink -f "${HOME}"/go)

#-c -v -gcflags "-N -l"

cd src/main || exit
go build -o ../../converter
cd - || exit
