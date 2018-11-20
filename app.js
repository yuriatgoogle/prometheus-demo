const express = require('express');
const app = express();
const client = require('prom-client');
const collectDefaultMetrics = client.collectDefaultMetrics;

// Probe every 5th second.
collectDefaultMetrics({ timeout: 5000 });

const projectID = 'yuri-next2019'; //TODO - make this a config setting

function getRandomInt(max) {
    return Math.floor(Math.random() * Math.floor(max));
  }

app.get('/', (req, res) => {
    console.log('request made');
    res.send('!');
})

app.listen(8080, () => console.log(`Example app listening on port 8080!`))
