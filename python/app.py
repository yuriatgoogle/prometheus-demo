from flask import Flask
import logging
import random
import string
import time
import datetime
app = Flask(__name__)

@app.route('/')
def homePage():
    return ("home page")

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')