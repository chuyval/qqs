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