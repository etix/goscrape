# Goscrape: Go library to scrape torrents on any UDP tracker

This library partially implement [BEP-0015](http://www.bittorrent.org/beps/bep_0015.html)
and returns the number of **seeders**, **leechers** and **completed** downloads for a given infohash list.

Usage:

`go get -u github.com/etix/goscrape`

Example:

```go
// Create a new instance of the library and specify the torrent tracker to use.
s, err := goscrape.New("udp://tracker.kali.org:6969/announce")
if err != nil {
    log.Fatal("Error:", err)
}

// A list of infohash to scrape, at most 74 infohash can be scraped at once.
// Be sure to provide infohash that are 40 hexadecimal characters long only.
infohash := [][]byte{
    []byte("176e2a9696092482d4acdef445b53ffcebb56960"),
    []byte("e80cb87fbd938f3b1e47db64c10c3ab04ad49987"),
    []byte("04098e49061bedb3f2d8f90204bf239019d198d9"),
}

// Connect to the tracker and scrape the list of infohash in only two UDP round trips.
res, err := s.Scrape(infohash...)
if err != nil {
    log.Fatal("Error:", err)
}

// Loop over the results and print them.
// Result are guaranteed to be in the same order they were requested.
for _, r := range res {
    fmt.Println("Infohash:\t", string(r.Infohash))
    fmt.Println("Seeders:\t", r.Seeders)
    fmt.Println("Leechers:\t", r.Leechers)
    fmt.Println("Completed:\t", r.Completed)
    fmt.Println("")
}
```

Result:

```
Infohash:        176e2a9696092482d4acdef445b53ffcebb56960
Seeders:         3053
Leechers:        248
Completed:       44171

Infohash:        e80cb87fbd938f3b1e47db64c10c3ab04ad49987
Seeders:         376
Leechers:        20
Completed:       4161

Infohash:        04098e49061bedb3f2d8f90204bf239019d198d9
Seeders:         100
Leechers:        6
Completed:       703
```