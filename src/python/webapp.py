import logging
import io, json
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
        logging.warning('Metodo Post: %s' % request.method )
        json_dict = request.get_json()
        
        with io.open('data.json', 'w', encoding='utf-8') as f:
            f.write(json.dumps(json_dict, ensure_ascii=False))
        logging.warning("Recebi o Post")
    else:
        logging.warning('Metodo: %s ' % request.method )
    return 'Yeah'
