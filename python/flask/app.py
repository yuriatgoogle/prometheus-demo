from flask import Flask
from prometheus_client import start_http_server, Summary, Gauge, generate_latest, REGISTRY
import logging
import random
import string
import time
import datetime

app = Flask(__name__)
RANDOM_VALUE = Gauge('random_value', 'randomly generated')


@app.route('/')
def homePage():
    RANDOM_VALUE.set(random.random())
    return ("home page")

@app.route('/stats', methods=['GET'])
def metrics():
    RANDOM_VALUE.set(random.random())
    return generate_latest(REGISTRY), 200

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0', port=8080)