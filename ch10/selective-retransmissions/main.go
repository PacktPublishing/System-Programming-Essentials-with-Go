package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	maxDatagramSize = 1024
	packetLossRate  = 0.2
)

type Packet struct {
	SeqNum  uint32
	Payload []byte
}

func main() {
	rand.Seed(time.Now().UnixNano())

	addr, err := net.ResolveUDPAddr("udp", ":5000")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, maxDatagramSize)
		for {
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				continue
			}
			if n < 4 { // Ensure the packet has at least 4 bytes (for the sequence number)
				fmt.Println("Received packet too short")
				continue
			}
			receivedSeq, _ := unpackUint32(buf[:4])
			fmt.Printf("Received: %d from %s\n", receivedSeq, addr.String())

			sendAck(conn, addr, receivedSeq)
		}
	}()

	clientAddr := &net.UDPAddr{} // Placeholder for client address
	var nextSeqNum uint32 = 0

	for {
		packet := &Packet{
			SeqNum:  nextSeqNum,
			Payload: []byte("Test Payload"),
		}
		sendPacket(conn, clientAddr, packet)
		time.Sleep(1 * time.Second) // Adjust as needed
		nextSeqNum++
	}
}

func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, packet *Packet) {
	if addr == nil || addr.IP == nil {
		return // No client to send to yet
	}
	buf := make([]byte, 4+len(packet.Payload))
	binary.BigEndian.PutUint32(buf[:4], packet.SeqNum)
	copy(buf[4:], packet.Payload)

	// Simulate packet loss
	if rand.Float32() > packetLossRate {
		_, err := conn.WriteToUDP(buf, addr)
		if err != nil {
			fmt.Println("Error sending packet:", err)
		} else {
			fmt.Printf("Sent: %d to %s\n", packet.SeqNum, addr.String())
		}
	} else {
		fmt.Printf("Simulated packet loss, seq: %d\n", packet.SeqNum)
	}
}

func sendAck(conn *net.UDPConn, addr *net.UDPAddr, seqNum uint32) {
	ackPacket := make([]byte, 4)
	binary.BigEndian.PutUint32(ackPacket, seqNum)
	_, err := conn.WriteToUDP(ackPacket, addr)
	if err != nil {
		fmt.Println("Error sending ACK:", err)
	}
}

func unpackUint32(buf []byte) (uint32, error) {
	if len(buf) < 4 {
		return 0, fmt.Errorf("buffer too short")
	}
	return binary.BigEndian.Uint32(buf), nil
}
