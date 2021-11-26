# dnsloader

Usage of ./dnsloader:
  -c int
        Connections (default 100)
  -d int
        Duration(s)
  -host string
        Resolve host
  -l int
        Limit(ms)
  -nameserver string
        Nameserver
  -protocol string
        tcp or udp (default "udp")

## load test nodelocaldns
./dnsloader -nameserver 169.254.20.10:53 -host mysql-product-console.pre012.intranet.cecloudcs.com -c 50 -d 180 -l 3000

## load test cc114
./dnsloader -nameserver 10.96.0.10:53 -host mysql-product-console.pre012.intranet.cecloudcs.com -c 50 -d 180 -l 3000
