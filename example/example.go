package main

import (
	"fmt"
	"log"

	"github.com/etix/goscrape"
)

func main() {
	// Create a new instance of the library and specify the torrent tracker.
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
		fmt.Println("Completed:\t", r.Completed)
		fmt.Println("Leechers:\t", r.Leechers)
		fmt.Println("")
	}
}
