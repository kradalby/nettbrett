package main

import (
    "log"
    "time"

    "github.com/alouca/gosnmp"
)

type NetworkInterface struct {
    bytesSent   int64
    bytesReceived int64
}


func SNMPRun() {

    // Initiate the SNMP client
    snmp, err := gosnmp.NewGoSNMP(cfg.Snmp.Server, cfg.Snmp.Community, gosnmp.Version2c, 5)
    if err != nil {
        log.Println("Error: ", err)
    }


    var pastUp int64
    var pastDown int64
    var pastUpT time.Time
    var pastDownT time.Time

    var speed int64 = cfg.Misc.IntSpeed

    for {
        //resp, err := s.Get(".1.3.6.1.2.1.31.1.1.1.6.4")  // Interface 4 linux
        //resp, err := s.Get(".1.3.6.1.2.1.31.1.1.1.6.9")    // Interface 9 cisco 2940
        down := getFirstSNMPValue(cfg.Snmp.InByte, s)
        up := getFirstSNMPValue(cfg.Snmp.OutByte, s)

        downKBits := calculateSpeed(down, pastDown, &pastDownT)
        upKBits := calculateSpeed(up, pastUp, &pastUpT)


        pastUp = up
        pastDown = down

        time.Sleep(cfg.Misc.Interval * time.Second)
    }


}

func getFirstSNMPValue(oid string, s *gosnmp.GoSNMP) int64 {
    resp, err := s.Get(oid)
    if err != nil {
        log.Println("Error: ", err)
    }
    return resp.Variables[0].Value.(int64)
}

func calculateSpeed(current, past int64, tpast *time.Time) int64 {
    now := time.Now()

    delta := current - past
    t := *tpast
    tdelta := now.UnixNano() - t.UnixNano()

    usage := (delta*8)/(1000*(tdelta/1000/1000/1000))

    *tpast = now
    return usage
}
