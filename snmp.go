package main

import (
    "log"
    "time"

    "github.com/alouca/gosnmp"
)

type NetworkInterface struct {
    BytesReceived int64    `json:"bytesReceived"`
    BytesSent   int64      `json:"bytesSent"`
    SpeedDown   int64      `json:"speedDown"`
    SpeedUp     int64      `json:"speedUp"`
}


func SNMPRun() {
    go uplinkRunner()
}

func getFirstSNMPValue(oid string, s *gosnmp.GoSNMP) int64 {
    resp, err := s.Get(oid)
    if err != nil {
        log.Println("Error: ", err)
    }
    return resp.Variables[0].Value.(int64)
}

func calculateSpeed(current, past *int64, tpast *time.Time) int64 {
    now := time.Now()

    delta := *current - *past
    t := *tpast
    tdelta := now.UnixNano() - t.UnixNano()

    usage := (delta*8)/(1000*(tdelta/1000/1000/1000))

    *tpast = now
    return usage
}

func uplinkRunner() {
    // Initiate the SNMP client
    snmp, err := gosnmp.NewGoSNMP(config.Uplink.IP, config.Uplink.Community, gosnmp.Version2c, 5)
    if err != nil {
        log.Println("Error: ", err)
    }


    networkInterfacePastState := NetworkInterface{BytesSent: 0, BytesReceived: 0}
    networkInterfaceCurrentState := NetworkInterface{BytesSent: 0, BytesReceived: 0}
    var pastUpT time.Time
    var pastDownT time.Time

    for {
        //resp, err := s.Get(".1.3.6.1.2.1.31.1.1.1.6.4")  // Interface 4 linux
        //resp, err := s.Get(".1.3.6.1.2.1.31.1.1.1.6.9")    // Interface 9 cisco 2940
        networkInterfaceCurrentState.BytesReceived = getFirstSNMPValue(config.Uplink.InByte, snmp)
        networkInterfaceCurrentState.BytesSent = getFirstSNMPValue(config.Uplink.OutByte, snmp)

        networkInterfaceCurrentState.SpeedDown = calculateSpeed(&networkInterfaceCurrentState.BytesReceived, &networkInterfacePastState.BytesReceived, &pastDownT)
        networkInterfaceCurrentState.SpeedUp = calculateSpeed(&networkInterfaceCurrentState.BytesSent, &networkInterfacePastState.BytesSent, &pastUpT)

        log.Println(networkInterfaceCurrentState)

        prepareAndDistributeWSData("uplink", networkInterfaceCurrentState)

        networkInterfacePastState = networkInterfaceCurrentState

        time.Sleep(config.Misc.Interval * time.Second)
    }
}
