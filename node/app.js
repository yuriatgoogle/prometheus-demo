const express = require('express');
const app = express();
const client = require('prom-client');
const collectDefaultMetrics = client.collectDefaultMetrics;
const projectID = 'yuri-next2019'; //TODO - make this a config setting

//counter metric
const counter = new client.Counter({
    name: 'example_counter',
    help: 'counter_help'
});

// Probe every 5th second.
collectDefaultMetrics({ timeout: 5000 });



function getRandomInt(max) {
    return Math.floor(Math.random() * Math.floor(max));
  }

app.get('/', (req, res) => {
    console.log('request made');
    counter.inc();
    res.send('current counter value is: ' + counter.hash_map.value);
})

app.listen(8080, () => console.log(`Example app listening on port 8080!`))
