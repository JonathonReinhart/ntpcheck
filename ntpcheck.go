package main

import (
    "fmt"
    "os"
    "time"

    "github.com/beevik/ntp"
)

func durationStr(t time.Duration) string {
    if t < 0 {
        return fmt.Sprintf("%v %s", -t, "behind")
    } else {
        return fmt.Sprintf("%v %s", t, "ahead")
    }
}

func describeLeap(v ntp.LeapIndicator) string {
    switch v {
    case ntp.LeapNoWarning:
        return "no warning"
    case ntp.LeapAddSecond:
        return "last minute has 61 seconds"
    case ntp.LeapDelSecond:
        return "last minute has 59 seconds"
    case ntp.LeapNotInSync:
        return "alarm; clock not synchronized"
    default:
        return "<invalid value>"
    }
}

type field struct {
    name string
    value interface{}
}

func printResponse(r *ntp.Response) {
    var fields = []field {
        {
            name: "Time",
            value: r.Time,
        },
        {
            name: "ClockOffset",
            value: durationStr(r.ClockOffset),
        },
        {
            name: "RTT",
            value: r.RTT,
        },
        {
            name: "Precision",
            value: r.Precision,
        },
        {
            name: "Stratum",
            value: r.Stratum,
        },
        {
            name: "ReferenceID",
            value: fmt.Sprintf("0x%08X", r.ReferenceID),
        },
        {
            name: "ReferenceTime",
            value: r.ReferenceTime,
        },
        {
            name: "RootDelay",
            value: r.RootDelay,
        },
        {
            name: "RootDispersion",
            value: r.RootDispersion,
        },
        {
            name: "RootDistance",
            value: r.RootDistance,
        },
        {
            name: "Leap",
            value: fmt.Sprintf("%v (%s)", r.Leap, describeLeap(r.Leap)),
        },
        {
            name: "MinError",
            value: r.MinError,
        },
        {
            name: "KissCode",
            value: fmt.Sprintf("%q", r.KissCode),
        },
        {
            name: "Poll",
            value: r.Poll,
        },
    }

    fmtstr := "  %-16s %v\n"

    for _, f := range fields {
        fmt.Printf(fmtstr, f.name + ":", f.value)
    }
}

func check(host string) {
    resp, err := ntp.Query(host)
    if err != nil {
        fmt.Printf("Error querying host %v: %v\n", host, err)
        return
    }

    err = resp.Validate()
    valid := "valid"
    if err != nil {
        valid = fmt.Sprintf("invalid: %s", err)
    }

    fmt.Printf("Response from %v: (%s)\n", host, valid)
    printResponse(resp)
}

func main() {
    hosts := os.Args[1:]
    if len(hosts) == 0 {
        fmt.Fprintln(os.Stderr, "Usage: ntpcheck HOST...")
        os.Exit(1)
    }
    for _, host := range os.Args[1:] {
        check(host)
    }
}
