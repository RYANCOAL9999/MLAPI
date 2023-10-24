import os
import logging
from typing import Any

from Python.quotes.lib.helper import generatePortMulti
from Python.quotes.model.svmService import SVMService
from Python.quotes.model.linearService import LinearService
from Python.quotes.model.generalService import GeneralService
from Python.quotes.model.neighborsService import NeighborsService
from Python.quotes.model.polynomialService import PolynomialService
from flask import Flask, jsonify, make_response, request, Response, status

# https://flask-api.github.io/flask-api/api-guide/status-codes/

gRPCKey = str(os.environ['GRPCKEY'])

port = int(os.environ['PORT'])

address = "localhost:"

generalService = GeneralService("localhost:" + generatePortMulti(gRPCKey, port, 50051))

linearService = LinearService("localhost:" + generatePortMulti(gRPCKey, port, 50051))

neighborsService = NeighborsService("localhost:" + generatePortMulti(gRPCKey, port, 50052) )

polynomialService = PolynomialService("localhost:" + generatePortMulti(gRPCKey, port, 50053) )

svmService = SVMService("localhost:" + generatePortMulti(gRPCKey, port, 50054) )

app = Flask(__name__)

@app.route("/", methods=["GET"])
def handle_root_request()->Response: 
    return make_response(
        jsonify(
            message = 'Do not wanna attack my server!'
        ),
        status.HTTP_204_NO_CONTENT
    )

@app.route("/api/head", methods=["GET"])
def head()->Response:  
    return make_response(
        jsonify(
            generalService.headerResponse(

            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/info", methods=["GET"])
def info()->Response: 
    return make_response(
        jsonify(
            generalService.infoResponse(

            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/describle", methods=["GET"])
def describle()->Response: 
    return make_response(
        jsonify(
            generalService.describleResponse(

            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/linearSVC", methods=["POST"])
def linearSVC()->Response:
    return make_response(
        jsonify(
            svmService.linearSVCResponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/linearSVR", methods=["POST"])
def linearSVR()->Response: 
    return make_response(
        jsonify(
            svmService.linearSVRResponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/svc", methods=["POST"])
def svc()->Response: 
    return make_response(
        jsonify(
            svmService.linearSVCResponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/polynomialFeatures", methods=["POST"])
def polynomialFeatures()->Response: 
    return make_response(
        jsonify(
            polynomialService.polynomialFeaturesResponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/polynomialFeaturesFitTransform", methods=["POST"])
def polynomialFeaturesFitTransform()->Response: 
    return make_response(
        jsonify(
            polynomialService.polynomialFeaturesFitTransformResponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )


@app.route("/api/nearestNeighbors", methods=["POST"])
def nearestNeighbors()->Response: 
    return make_response(
        jsonify(
            neighborsService.nearestCentroidReponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/kdTree", methods=["POST"])
def kdTree()->Response: 
    return make_response(
        jsonify(
            neighborsService.kdTreeResponse(request.json)
        ),
        status.HTTP_200_OK
    )

@app.route("/api/nearestCentroid", methods=["POST"])
def nearestCentroid()->Response: 
    return make_response(
        jsonify(
            neighborsService.nearestCentroidReponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/linearRegression", methods=["POST"])
def linearRegression()->Response: 
    return make_response(
        jsonify(
            linearService.linearRegressionResponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/ridge", methods=["POST"])
def ridge()->Response: 
    return make_response(
        jsonify(
            linearService.ridgeRespone(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/ridgeCV", methods=["POST"])
def ridgeCV()->Response: 
    return make_response(
        jsonify(
            linearService.ridgeCVReponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/lasso", methods=["POST"])
def lasso()->Response: 
    return make_response(
        jsonify(
            linearService.lassoReponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/lassoLars", methods=["POST"])
def lassoLars()->Response: 
    return make_response(
        jsonify(
            linearService.lassoLarsReponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/bayesianRidge", methods=["POST"])
def bayesianRidge()->Response: 
    return make_response(
        jsonify(
            linearService.BayesianRidgeResponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/tweedieRegressor", methods=["POST"])
def tweedieRegressor()->Response: 
    return make_response(
        jsonify(
            linearService.tweedieRegressorResponse(request.json)
        ),
        status.HTTP_200_OK
    )

@app.route("/api/sgdClassifier", methods=["POST"])
def sgdClassifier()->Response: 
    return make_response(
        jsonify(
            linearService.sgdClassifierReponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.route("/api/elasticNet", methods=["POST"])
def elasticNet()->Response: 
    return make_response(
        jsonify(
            linearService.elasticNetReponse(
                request.json
            )
        ),
        status.HTTP_200_OK
    )

@app.errorhandler(404)
def resource_not_found()->Any: 
    return make_response(
        jsonify(
            error = 'Not found!'
        ), 
        status.HTTP_404_NOT_FOUND
    )

if __name__ == "__main__":
    logging.basicConfig(
        filename = None,
        filemode = "",
        format = "",
        datefmt = None,
        style = "",
        level = logging.INFO,
        stream = None,
        handlers = None,
        force = None,
        encoding = None,
        errors = None
    )