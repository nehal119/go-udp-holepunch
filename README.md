# UDP Hole Punching Example in Go
This is a very simple repo to help demonstrate how UDP hole punching works in
Go. It includes the signal server and a client.

![Demo](/demo.gif "Demo")

## Using http rendezvous server
1. rendezvous server runs on http exposing `/register` route.
2. When two clients are connected, each get others udp ip address and port to start hole punching. 
```bash
  node rendezvous/server.js # <- Signalling Server
  ./hp A # <-- client 1
  ./hp B # <-- client 2
```

## Build
1. Build using golang 1.17
```bash
  go build -o hp .
  ./hp s  # <-- signaling server
  ./hp c1 # <-- client 1
  ./hp c2 # <-- client 2
```

#### Start background process
nohup ./hp s > output.txt 2>&1 < /dev/null &

## Usage
1. Set up signal server on a publicly accessible server.
```
./hp s :9595
```
2. Run client from a machine behind a firewall:
```
./hp c ip-address-of-signal:9595 :4545
```
3. Run client from a different machine behind a different firewall:
```
./hp c ip-address-of-signal:9595 :4545
```

After doing this the two machines behind firewalls should start sending `Hello!` back and forth to each other via UDP. You can shut off the signal server and they will still be able to communicate peer to peer.

## API
For both the server and the client you can specify the local address and port.
You may exclude the local address and specify just the port. This is what most use cases would do.
Leaving this off entirely will use `:9595` which is to say any local address on port :9595.

### Server
```
./hp s [local-address:port]
```

### Client
```
./hp c [signal-host:port] [local-address:port]
```
You may exclude the local address and specify just the port. This is what most use cases would do.
Leaving this off entirely will use `:9595`

## Notes
[https://en.wikipedia.org/wiki/UDP_hole_punching](https://en.wikipedia.org/wiki/UDP_hole_punching)

[Peer-to-Peer Communication Across Network Address Translators](https://bford.info/pub/net/p2pnat/)

