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

# Notes:
Build using

| Golang Version | Vendor Dir Exists | Build                          | Binary Size |
|:--------------:|:-----------------:|:------------------------------:|:-----------:|
| 1.10.3         | True              | `go build -ldflags=-s ./`      | 2457328     |
| 1.10.3         | True              | `go build -ldflags=-w ./`      | 1552112     |
| 1.10.3         | True              | `go build -ldflags "-s -w" ./` | 1552112     |
| 1.10.3         | False             | `go build -ldflags=-s ./`      | 2377872     |
| 1.10.3         | False             | `go build -ldflags=-w ./`      | 1550480     |
| 1.10.3         | False             | `go build -ldflags "-s -w" ./` | 1550480     |
| 1.11           | True              | `go build -ldflags=-s ./`      | 2220072     |
| 1.11           | True              | `go build -ldflags=-w ./`      | 1744936     |
| 1.11           | True              | `go build -ldflags "-s -w" ./` | 1744936     |
| 1.11           | False             | `go build -ldflags=-s ./`      | 2218440     |
| 1.11           | False             | `go build -ldflags=-w ./`      | 1743304     |
| 1.11           | False             | `go build -ldflags "-s -w" ./` | 1743304     |
