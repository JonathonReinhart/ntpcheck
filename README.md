ntpcheck
========
`ntpcheck` is a simple utility that queries an NTP server
and displays the result in a human-readable fashion.

It uses the [`github.com/beevik/ntp`](https://pkg.go.dev/github.com/beevik/ntp)
package.

### Usage
```
ntpcheck HOST...
```

Simply invoke `ntpcheck` with one or more NTP servers.

Sample output:
```
$ ./ntpcheck 0.beevik-ntp.pool.ntp.org
Response from 0.beevik-ntp.pool.ntp.org: (valid)
  Time:            2020-08-24 00:39:42.468735102 +0000 UTC
  ClockOffset:     41.777389ms behind
  RTT:             119.170566ms
  Precision:       59ns
  Stratum:         2
  ReferenceID:     0xED11CC5F
  ReferenceTime:   2020-08-24 00:32:49.036165874 +0000 UTC
  RootDelay:       9.307861ms
  RootDispersion:  22.30835ms
  RootDistance:    86.547563ms
  Leap:            0 (no warning)
  MinError:        0s
  KissCode:        ""
  Poll:            8s
```
