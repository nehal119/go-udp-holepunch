#!/usr/bin/env node
var dgram = require('dgram');

// based on http://www.bford.info/pub/net/p2pnat/index.html



var socket = dgram.createSocket('udp4');

socket.bind(4545, '0.0.0.0');

socket.on('listening', function () {
    console.log('UDP Server listening on' + socket.address().address + ":" + socket.address().port);
});

socket.on('message', function (message, remote) {
    console.log("Message received from", remote.address + ":" + remote.port, "the message is", message.toString());
});

// function sendMessageToS () {
// 	var request = require('request');
// 	var options = {
// 		'method': 'POST',
// 		'url': 'http://IP:33333/register',
// 		'headers': {
// 			'Content-Type': 'application/json'
// 		},
// 		body: JSON.stringify({
// 			"client": "A",
// 			"address": "CLIENTB_IP:4545"
// 		})

// 	};
// 	request(options, function (error, response) {
// 		if (error) throw new Error(error);
// 		console.log(response.body);
// 		var publicEndpointB = JSON.parse(response.body);
// 		const address = publicEndpointB.address.split(':');
//     sendMessageToB(address[0], address[1]);
// 	});
// }

// sendMessageToS();
// sendMessageToB(address, port);
setInterval(function () {
	sendMessageToB('34.93.82.205', 4546);
}, 2000);

var counter = 0;
function sendMessageToB (address, port) {
	// if(counter == 5) return;
	var message = new Buffer(counter++ + ': Hello B!');
	socket.send(message, 0, message.length, port, address, function (err, nrOfBytesSent) {
	    if (err) return console.log(err);
	    console.log('UDP message sent to B:', address +':'+ port);

	    // setTimeout(function () {
	    // 	sendMessageToB(address, port);
	    // }, 2000);
	});
}
