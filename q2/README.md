# Why does proto marshal then unmarshal fail when it is run outside of project containing vendor directory?

I have a `main.go` file that uses the proto files in `pkg/models` to Marshal
and Unmarshal a proto struct like this:

```
// Convert to string
protoStr := proto.MarshalTextString(proto)

// Unmarshal string back to proto struct
var proto2 models.Stuff
err := proto.UnmarshalText(protoStr, &proto2)
```

The setup is here: https://github.com/chuyval/qqs/tree/master/q2

The project contains a vendor directory that only has the `github.com/golang/protobuf`
repo checked out. (run `glide install` to create vendor if it doesnt exist)

The `main.go` program works fine when running `go run main.go` from inside the project.

When I move the `main.go` file one level up to the parent directory, and run
the same command `go run main.go` at the parent level, it reports the following error:

```
line 2: unknown field name "value_list" in models.Stuff
```

When I delete the vendor directory in the project directory, and run
`go run main.go` at the parent level, I get no error.

__Why would having a vendor directory in the project repository make it error out?__


Sample code:
```
package main

import (
    "github.com/chuyval/qqs/q2/pkg/models"
    "github.com/golang/protobuf/proto"
    "log"
)

func main () {
    stuff := createProtoStuff()
    log.Printf("Stuff: %+v", stuff)

    // Convert to string
    stuffStr := proto.MarshalTextString(stuff)

    // Unmarshal string back to proto struct
    var stuff2 models.Stuff
    err := proto.UnmarshalText(stuffStr, &stuff2)
    if err != nil {
        log.Printf("It didnt work. Error: %s", err.Error())
    } else {
        log.Printf("It worked. Proto: %+v", stuff2)
    }
}

func createProtoStuff() *models.Stuff {
    someValueList := []*models.SomeValue{&models.SomeValue{Id: "Some Value ID"}}

    valueList := &models.SomeValueList{SomeValue: someValueList}

    stuffValueList := &models.Stuff_ValueList{
        ValueList: valueList,
    }

    stuff := &models.Stuff{
        Id: "Stuff List Id",
        Stuff: stuffValueList,
    }

    return stuff
}
```

# Software Versions
```
glide --version
glide version 0.13.1

go version
go version go1.10.3 darwin/amd64

protoc --version
libprotoc 3.6.0
```

## Build proto files
```
protoc -I. --go_out=$GOPATH/src ./*.proto
```

## Installing dependencies
```
glide install
```