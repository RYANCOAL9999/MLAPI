import os
import logging
from typing import Any

from Python.quotes.model.svmService import SVMService
from Python.quotes.model.linearService import LinearService
from Python.quotes.model.generalService import GeneralService
from Python.quotes.model.neighborsService import NeighborsService
from Python.quotes.model.polynomialService import PolynomialService
from Python.quotes.lib.helper import checkAPI, generatePortMulti
from flask import Flask, jsonify, make_response, request, Response
from flask.status import HTTP_204_NO_CONTENT, HTTP_200_OK, HTTP_404_NOT_FOUND

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

@app.route("/api/head", methods=["GET"])
def head()->Response:  
    return make_response(
        jsonify(
            generalService.headerResponse(

            )
        ),
        HTTP_200_OK
    )

@app.route("/api/info", methods=["GET"])
def info()->Response: 
    return make_response(
        jsonify(
            generalService.infoResponse(

            )
        ),
        HTTP_200_OK
    )

@app.route("/api/describle", methods=["GET"])
def describle()->Response: 
    return make_response(
        jsonify(
            generalService.describlerResponse(

            )
        ),
        HTTP_200_OK
    )

@app.route("/api/linearSVC", methods=["POST"])
def linearSVC()->Response:
    response = checkAPI(request, svmService.linearSVCResponse)
    if response is None:
        response = make_response(
            jsonify(
                svmService.linearSVCResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/linearSVR", methods=["POST"])
def linearSVR()->Response:
    response = checkAPI(request, svmService.linearSVRResponse)
    if response is None:
        response = make_response(
            jsonify(
                svmService.linearSVRResponse(
                    request.json
                )
            )
        )
    return response

@app.route("/api/svc", methods=["POST"])
def svc()->Response:
    response = checkAPI(request, svmService.svcResponse)
    if response is None:
        response = make_response(
            jsonify(
                svmService.svcResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response 

@app.route("/api/polynomialFeatures", methods=["POST"])
def polynomialFeatures()->Response:
    response = checkAPI(request, polynomialService.polynomialFeaturesResponse)
    if response is None:
        response = make_response(
            jsonify(
                polynomialService.polynomialFeaturesResponse(
                    request.json
                )
            ),
            HTTP_200_OK    
        )
    return response 

@app.route("/api/polynomialFeaturesFitTransform", methods=["POST"])
def polynomialFeaturesFitTransform()->Response: 
    response = checkAPI(request, polynomialService.polynomialFeaturesFitTransformResponse)
    if response is None:
        response = make_response(
            jsonify(
                polynomialService.polynomialFeaturesFitTransformResponse(
                    request.json
                )
            ),
            HTTP_200_OK 
        )
    return response


@app.route("/api/nearestNeighbors", methods=["POST"])
def nearestNeighbors()->Response: 
    response = checkAPI(request,neighborsService.nearestNeighborsResponse)
    if response is None:
        response = make_response(
            jsonify(
                neighborsService.nearestNeighborsResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/kdTree", methods=["POST"])
def kdTree()->Response: 
    response = checkAPI(request, neighborsService.kdTreeResponse)
    if response is None:
        response = make_response(
            jsonify(
                neighborsService.kdTreeResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response    

@app.route("/api/nearestCentroid", methods=["POST"])
def nearestCentroid()->Response: 
    response = checkAPI(request, neighborsService.nearestCentroidResponse)
    if response is None:
        response = make_response(
            jsonify(
                neighborsService.nearestCentroidResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/linearRegression", methods=["POST"])
def linearRegression()->Response: 
    response = checkAPI(request, linearService.linearRegressionResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.linearRegressionResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/ridge", methods=["POST"])
def ridge()->Response: 
    response = checkAPI(request, linearService.ridgeResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.ridgeResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/ridgeCV", methods=["POST"])
def ridgeCV()->Response: 
    response = checkAPI(request, linearService.ridgeCVResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.ridgeCVResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/lasso", methods=["POST"])
def lasso()->Response: 
    response = checkAPI(request, linearService.lassoResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.lassoResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response    

@app.route("/api/lassoLars", methods=["POST"])
def lassoLars()->Response: 
    response = checkAPI(request, linearService.lassoLarsResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.lassoLarsResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/bayesianRidge", methods=["POST"])
def bayesianRidge()->Response: 
    response = checkAPI(request, linearService.BayesianRidgeResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.BayesianRidgeResponse(
                    request.json
                )
            ),
            HTTP_200_OK 
        )
    return response    

@app.route("/api/tweedieRegressor", methods=["POST"])
def tweedieRegressor()->Response: 
    response = checkAPI(request, linearService.tweedieRegressorResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.tweedieRegressorResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response

@app.route("/api/sgdClassifier", methods=["POST"])
def sgdClassifier()->Response:
    response = checkAPI(request, linearService.sgdClassifierResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.sgdClassifierResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response 

@app.route("/api/elasticNet", methods=["POST"])
def elasticNet()->Response:
    response = checkAPI(request, linearService.elasticNetResponse)
    if response is None:
        response = make_response(
            jsonify(
                linearService.elasticNetResponse(
                    request.json
                )
            ),
            HTTP_200_OK
        )
    return response 

@app.errorhandler(404)
def resource_not_found()->Any: 
    return make_response(
        jsonify(
            error = 'Not found!'
        ), 
        HTTP_404_NOT_FOUND
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