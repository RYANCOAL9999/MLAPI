const PROTO_PATH = __dirname + '../lib/AIProto/polynomial_features.proto';
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { defination } = require('./packageDefinitionpackageDefinition')
const { sendStubRequest, handleResponse, handleGRPCError, handleBadRequestError } = require('./helper')

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    defination
);

const polynomial_proto = grpc.loadPackageDefinition(packageDefinition).polynomial_features;

export const createPolynomialStub = (addr) => {
    return new polynomial_proto.NeighborsService(
        addr, 
        grpc.credentials.createInsecure()
    );
};

export const polynomialFeaturesResponse = (req, res, stub) => {

    const event_name = "polynomial features"

    const request_message = req.body;

    if (request_message) {
        sendStubRequest(stub, `PolynomialFeaturesEvent`, request_message)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

export const polynomialFeaturesFitTransformResponse = (req, res, stub) => {
    
    const event_name = "polynomial features fit transform"

    const request_message = req.body;

    if (request_message) {
        sendStubRequest(stub, `PolynomialFeaturesFitTransformEvent`, request_message)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name);
    }
};

