from flask import Flask
from prometheus_client import Counter
import json
import logging

# create counter metric
c = Counter('page_views_counter','increment every time /metrics is loaded')

app = Flask(__name__)

@app.route("/")
def incCounter():
    c.inc()


if __name__ == "__main__":
    app.run()