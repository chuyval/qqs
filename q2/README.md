# Why does proto marshal then unmarshal fail when dependency repository has a vendor directory?

I have a `main.go` file that has a dependency on a different repository
with a vendor directory. My program attempts to marshal and unmarshal a
proto struct located in the other repository like this:

```
// Convert to string
protoStr := proto.MarshalTextString(proto)

// Unmarshal string back to proto struct
var proto2 models.Stuff
err := proto.UnmarshalText(protoStr, &proto2)
```

but fails with error: `line 2: unknown field name "value_list" in models.Stuff`

When I delete the vendor directory in the dependent repository, build,
and run again, I get no error.

__Why would having a vendor directory in the dependent repository make it error out?__

Another thing to note, if I run the same `main.go` application inside of the
dependent repo, it works every time (with or without the vendor repository).

Sample code: (Also located in https://github.com/chuyval/qqs/tree/master/q2)
```
package main

import (
    "github.com/chuyval/qqs-helper/q2/pkg/models"
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
glide --version
glide version 0.13.1

go version
go version go1.10.3 darwin/amd64

protoc --version
libprotoc 3.6.0
