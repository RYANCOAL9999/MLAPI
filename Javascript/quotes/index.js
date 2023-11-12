const serverless = require("serverless-http");
const express = require("express");

const {
    createGeneralStub, 
    headerResponse, 
    infoResponse, 
    describlerResponse 
} = require("./model/generalService")

const {
    createSVMStub, 
    linearSVCResponse, 
    linearSVRResponse, 
    svcResponse 
} = require("./model/svmService")

const {
    createPolynomialStub, 
    polynomialFeaturesResponse, 
    polynomialFeaturesFitTransformResponse 
} = require("./model/polynomialService")

const {
    createNeighborsStub, 
    nearestNeighborsResponse, 
    kdTreeResponse, 
    nearestCentroidResponse 
} = require("./model/neighborsService")

const {
    createLinearStub, 
    linearRegressionResponse,
    ridgeResponse,
    ridgeCVResponse,
    lassoResponse,
    lassoLarsResponse,
    bayesianRidgeResponse,
    tweedieRegressorResponse,
    sgdClassifierResponse,
    elasticNetResponse 
} = require("./model/linearService");

const generalStub = createGeneralStub("localhost:50051");

const linearStub = createLinearStub("localhost:50051");

const polynomialStub = createPolynomialStub("localhost:50052");

const neighborsStub = createNeighborsStub("localhost:50053");

const svmStub = createSVMStub("localhost:50054");

const app = express();

app.get('/api/head', (req, res) => headerResponse(req, res, generalStub));

app.get('/api/info', (req, res) => infoResponse(req, res, generalStub));

app.get('/api/describle', (req, res) => describlerResponse(req, res, generalStub));

app.get('/api/linearSVC', (req, res) => linearSVCResponse(req, res, svmStub));

app.get('/api/linearSVR', (req, res) => linearSVRResponse(req, res, svmStub));

app.get('/api/linearSVR', (req, res) => svcResponse(req, res, svmStub));

app.get('/api/polynomialFeatures', (req, res) => polynomialFeaturesResponse(req, res, polynomialStub));

app.get('/api/polynomialFeaturesFitTransform', (req, res) => polynomialFeaturesFitTransformResponse(req, res, polynomialStub));

app.get('/api/nearestNeighbors', (req, res) => nearestNeighborsResponse(req, res, neighborsStub));

app.get('/api/kdTree', (req, res) => kdTreeResponse(req, res, neighborsStub));

app.get('/api/nearestCentroid', (req, res) => nearestCentroidResponse(req, res, neighborsStub));

app.get('/api/linearRegression', (req, res) => linearRegressionResponse(req, res, linearStub));

app.get('/api/ridge', (req, res) => ridgeResponse(req, res, linearStub));

app.get('/api/ridgeCV', (req, res) => ridgeCVResponse(req, res, linearStub));

app.get('/api/lasso', (req, res) => lassoResponse(req, res, linearStub));

app.get('/api/lassoLars', (req, res) => lassoLarsResponse(req, res, linearStub));

app.get('/api/bayesianRidge', (req, res) => bayesianRidgeResponse(req, res, linearStub));

app.get('/api/tweedieRegressor', (req, res) => tweedieRegressorResponse(req, res, linearStub));

app.get('/api/sgdClassifier', (req, res) => sgdClassifierResponse(req, res, linearStub));

app.get('/api/elasticNet', (req, res) => elasticNetResponse(req, res, linearStub));
  
app.use((req, res, next) => {return res.status(404).json({error: "Not Found"});});

model.exports.handler = serverless(app);