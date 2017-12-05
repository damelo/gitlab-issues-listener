import logging
import json
import structlog
from flask import Flask
from flask import Response
from flask import request
from flask import jsonify

APP = Flask(__name__)

logging.basicConfig()

@APP.route('/')
def index():
    return 'Index Page'


@APP.route('/', methods=['GET', 'POST'])
def putIssue():
    if request.method == 'POST':
        logging.warning('Método Post: %s' % request.method )
        json_dict = request.get_json()
        logging.warning(json_dict)
    else:
        logging.warning('Método: %s ' % request.method )
