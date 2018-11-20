package main

import (
	"log"

	"github.com/copernet/go-electrum/irc"
)

func main() {
	log.Println(irc.FindElectrumServers())
}
