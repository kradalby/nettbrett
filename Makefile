SHELL := /bin/bash

build:
	go build

dev:
	go get

run: 
	go run *.go

test: 
	go test

clean:
	rm -rf ~/go/pkg

snmp-tunnel:
	@echo "Opening SNMP tunnel to butterfree at port 16100"
	tmux new-session -s snmp-tunnel -d
	tmux new-window -a -n server -t snmp-tunnel -P "ssh butterfree -L 10001:localhost:10001 \"killall socat; socat -T10 TCP4-LISTEN:10001,fork UDP4:localhost:161\""
	tmux new-window -a -n client -t snmp-tunnel -P "socat -T15 udp4-recvfrom:16100,reuseaddr,fork tcp:localhost:10001"

close-snmp-tunnel:
	@echo "Closing SNMP tunnel to butterfree"
	tmux kill-session -t snmp-tunnel

all: run
