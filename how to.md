### How to test if udp port is working

1. Send a message and receive on the udp server
    - `echo -n "Test message" | nc -u -w1 IP_ADDRESS 33333`

2. Using iperf3:
    - Start `iperf3 -s -p 33333` on the server
    - Test is using `iperf3 -u -c IP_ADDRESS -p 33333 -b 1M` on the client side

>Note: Make sure the firewall allows port 33333 for udp connections on the server.

## Reference
1. https://bford.info/pub/net/p2pnat/
1. 