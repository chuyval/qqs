# Why is a go application binary different sizes when building using dependencies in a vendor directory vs dependencies in gopath?

I have a main.go file with a glide.yaml file in a clean gopath (No other repositories)

Running `go get -u ./...` then `go build main.go` generates a binary of size 2377872 bytes.

Cleaning the gopath of any repos that were cloned from `go get`, and running `glide update` then `go build main.go` generates a binary of size 2457328 bytes.

Why are binaries of different sizes, if there was no code changes?
What does go build do differently?
Why does having a vendor directory (or not having one) affect this size?

# Software Versions
glide --version
glide version 0.13.1

go version
go version go1.10.3 darwin/amd64

Asked in [stackoverflow](https://stackoverflow.com/questions/52240961/why-is-a-go-application-binary-different-sizes-when-building-using-dependencies)