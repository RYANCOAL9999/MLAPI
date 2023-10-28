const PROTO_PATH = __dirname + '../lib/AIProto/linear_expression.proto';
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { defination } = require('./packageDefinitionpackageDefinition')
const { sendStubRequest, handleResponse, handleGRPCError, handleBadRequestError } = require('./helper')

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    defination
);

const linear_proto = grpc.loadPackageDefinition(packageDefinition).linear_expression;

export const createLinearStub = (addr) => {
    return new linear_proto.LinearService(
        addr, 
        grpc.credentials.createInsecure()
    );
};

export const linearRegressionResponse = (req, res, stub) => {

    const event_name = "linear regression"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `LinearRegressionEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const ridgeResponse = (req, res, stub) => {

    const event_name = "ridge"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `LinearRidgeEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const ridgeCVResponse = (req, res, stub) => {

    const event_name = "ridge cv"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `LinearRidgeCVEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const lassoResponse = (req, res, stub) => {

    const event_name = "lasso"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `LassoExpressionEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const lassoLarsResponse = (req, res, stub) => {

    const event_name = "lasso lars"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `LassoLarsLassoExpressionEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const bayesianRidgeResponse = (req, res, stub) => {

    const event_name = "bayesian ridge"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `BayesianRidgeEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const tweedieRegressorResponse = (req, res, stub) => {

    const event_name = "tweedie regressor"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `TweedieRegressorEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const sgdClassifierResponse = (req, res, stub) => {

    const event_name = "sgd classifier"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `SGDClassifierEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const elasticNetResponse = (req, res, stub) => {

    const event_name = "elastic net"

    const request_message = req.body; 

    if(request_message){
        sendStubRequest(stub, `ElasticNetEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

