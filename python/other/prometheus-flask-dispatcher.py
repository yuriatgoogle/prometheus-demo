from flask import Flask
from werkzeug.wsgi import DispatcherMiddleware
from prometheus_client import make_wsgi_app, Gauge, Counter
import random
from metrics import app as metricsApp



# Create my app
app = Flask(__name__)

# Add prometheus wsgi middleware to route /metrics requests
app_dispatch = DispatcherMiddleware(app, {
    '/metrics': metricsApp
})