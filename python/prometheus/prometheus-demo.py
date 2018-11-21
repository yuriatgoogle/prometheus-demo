from prometheus_client import start_http_server, Summary, Gauge, Counter
import random
import time

# Create metrics
c = Counter('incrementing_counter', 'incrementing value')
g = Gauge('random_value', 'randomly set gauge value')

if __name__ == '__main__':
    # Start up the server to expose the metrics.
    start_http_server(8000)
    # Generate some requests.
    while True: # on every request
        c.inc() # increment the counter
        g.set(random.random()) # randomize the gauge