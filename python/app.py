from flask import Flask
from prometheus_client import Summary, Gauge, Counter, generate_latest, REGISTRY, Histogram
import random


app = Flask(__name__)
PYTHON_RANDOM_VALUE = Gauge('python_random_value', 'randomly generated in Python')


@app.route('/')
def homePage():
    PYTHON_RANDOM_VALUE.set(random.random())
    return ("home page")

@app.route('/metrics', methods=['GET'])
def stats():
    return generate_latest(REGISTRY), 200

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0', port=8080)