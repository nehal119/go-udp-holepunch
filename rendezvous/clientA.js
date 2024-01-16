#!/usr/bin/env node
var dgram = require('dgram');
var http = require('follow-redirects').http;
var fs = require('fs');
// based on http://www.bford.info/pub/net/p2pnat/index.html



var socket = dgram.createSocket('udp4');

// socket.on('message', function (message, remote) {
//     console.log(remote.address + ':' + remote.port +' - ' + message);
//     try{
//     	var publicEndpointB = JSON.parse(message);
//     	sendMessageToB(publicEndpointB.address, publicEndpointB.port);
//     }catch(err) {}
// });

function sendMessageToS () {
	// var serverPort = 33333;
	// var serverHost = '127.0.0.1';

	// var message = new Buffer('A');
	// socket.send(message, 0, message.length, serverPort, serverHost, function (err, nrOfBytesSent) {
	//     if (err) return console.log(err);
	//     console.log('UDP message sent to ' + serverHost +':'+ serverPort);
	//     // socket.close();
	// });
	var options = {
		'method': 'POST',
		'hostname': '127.0.0.1',
		'port': 33333,
		'path': '/register',
		'headers': {
			'Content-Type': 'application/json'
		},
		'maxRedirects': 20
	};

	var req = http.request(options, function (res) {
		var chunks = [];

		res.on("data", function (chunk) {
			chunks.push(chunk);
		});

		res.on("end", function (chunk) {
			var body = Buffer.concat(chunks);
			console.log("string", body.toString());
			console.log("body", JSON.parse(body.toString()));
			const {address, port} = JSON.parse(body.toString());
			sendMessageToB(address, port);
		});

		res.on("error", function (error) {
			console.error(error);
		});
	});

	var postData = JSON.stringify({
		"client": "A"
	});

	req.write(postData);

	req.end();
}

sendMessageToS();

var counter = 0;
function sendMessageToB (address, port) {
	// if(counter == 5) return;
	var message = new Buffer(counter++ + ': Hello B!');
	socket.send(message, 0, message.length, port, address, function (err, nrOfBytesSent) {
	    if (err) return console.log(err);
	    console.log('UDP message sent to B:', address +':'+ port);

	    setTimeout(function () {
	    	sendMessageToB(address, port);
	    }, 2000);
	});
}
