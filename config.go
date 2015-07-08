package main

import (
    "os"
    "io/ioutil"
    "log"
    "time"

    "github.com/naoina/toml"
)

type Config struct {
    Server map[string]Server

    Misc    struct {
        Interval time.Duration
    }
}

type Server struct {
    IP      string
    Community string
    InByte  string
    OutByte string
}

func readConfig(filename string) Config {
    f, err := os.Open(filename)

    if err != nil {
        log.Fatalln(err)
    }

    defer f.Close()

    buf, err := ioutil.ReadAll(f)

    if err != nil {
        log.Fatalln(err)
    }

    var config Config

    if err := toml.Unmarshal(buf, &config); err != nil {
        log.Fatalln(err)
    }

    if err != nil {
        log.Fatalln(err)
    }

    return config
}

