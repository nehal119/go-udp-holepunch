var express = require('express');
var app = express();
var bodyParser = require('body-parser');

app.use(bodyParser.json());

var publicEndpointA = null;
var publicEndpointB = null;

let clientA = undefined
let clientB = undefined

app.post('/register', function (req, res) {
    var { client, address } = req.body;
    // var remote = req.connection.remoteAddress + ':' + req.connection.remotePort;

    console.log(address + ' - ' + client);

    if(client == 'A') {
        publicEndpointA = {
            name: 'A',
            address: address,
        }
        clientA = res;
    }

    if(client == 'B') {
        publicEndpointB = {
            name: 'B',
            address: address,
        }
        clientB = res;
    }

    sendPublicDataToClients(res);
});

function sendPublicDataToClients (res) {
    if(publicEndpointA && publicEndpointB) {
        // res.json({
        //     clientA: publicEndpointA,
        //     clientB: publicEndpointB
        // });
        clientA.json(publicEndpointB);
        clientB.json(publicEndpointA);
        clientA = undefined;
        clientB = undefined;
        publicEndpointA = null;
        publicEndpointB = null;
        console.log('> public endpoint of B sent to A');
        console.log('> public endpoint of A sent to B');
    }
};

app.get('/ping', function (req, res) {
    res.status(200).json({ message: 'pong' });
});

app.listen(33333, '0.0.0.0', function () {
    console.log('HTTP Server listening on *:33333');
});