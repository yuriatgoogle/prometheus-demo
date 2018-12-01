const express = require('express');
const app = express();
const prometheus = require('prom-client');
const collectDefaultMetrics = prometheus.collectDefaultMetrics;
const projectID = 'yuri-next2019'; //TODO - make this a config setting

//counter metric
const counter = new prometheus.Counter({
    name: 'example_counter',
    help: 'counter_help'
});

// Probe every 5th second.
collectDefaultMetrics({ timeout: 5000 });

app.get('/', (req, res) => {
    console.log('request made');
    counter.inc();
    res.send("home page");
})

app.get('/metrics', (req, res) => {
    res.set('Content-Type', prometheus.register.contentType)
    res.end(prometheus.register.metrics())
  })

app.listen(8080, () => console.log(`Example app listening on port 8080!`))
