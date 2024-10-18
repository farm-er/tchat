package network

import (
	"fmt"
	"log"
	"net"
	"time"
)







func SendSignals( inter chan struct{},port int) error {

	log.Println("starting sending")
	addr, r := net.ResolveUDPAddr( "udp", castingAddr)

	if r != nil {
		return r
	}

	// making a connection
	s, r := net.DialUDP("udp", nil, addr)

	if r != nil {
		return r
	}

	stop := false

	for !stop {

		select {
		case <- inter:
			stop = true
		default:
		}

		s.Write([]byte(fmt.Sprintf("tchat: sending:%v", port)))

		time.Sleep(time.Second)

	}

	log.Println("finished sending")
	
	return nil

}


