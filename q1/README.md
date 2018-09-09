# Why is go app binary different size when building using libraries in vendor directory vs libraries in gopath?

I have a main.go file with a glide.yaml file in a clean gopath (No other
projects).

Running `go get -u ./...` then `go build main.go` generates a binary of
size 2377872 bytes.

Cleaning the gopath of any repos that were cloned from `go get`, and
running `glide update` then `go build main.go` generates a binary
of size 2457328 bytes.

Why are binaries of different sizes, if there was no code changes?
Why does having a vendor directory (or not having one) affect this size?



# Software Versions
glide --version
glide version 0.13.1

go version
go version go1.10.3 darwin/amd64