const PROTO_PATH = __dirname + '../lib/AIProto/nearest_neighbors.proto';
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { defination } = require('./packageDefinitionpackageDefinition')
const { sendStubRequest, handleResponse, handleGRPCError, handleBadRequestError } = require('./helper')

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    defination
);

const neighbors_proto = grpc.loadPackageDefinition(packageDefinition).nearest_neighbors;

export const createNeighborsStub = (addr) => {
    return new neighbors_proto.NeighborsService(
        addr, 
        grpc.credentials.createInsecure()
    );
};

export const nearestNeighborsResponse = (req, res, stub) => {

    const event_name = "nearest neighbors"

    const request_message = req.body;

    if(request_message){
        sendStubRequest(stub, `NearestNeighborsEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name)
    }
};

export const kdTreeResponse = (req, res, stub) => {

    const event_name = "kd tree"

    const request_message = req.body;

    if(request_message){
        sendStubRequest(stub, `KDTreeEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name)
    }
};

export const nearestCentroidResponse = (req, res, stub) => {

    const event_name = "nearest centroid"

    const request_message = req.body;

    if(request_message){
        sendStubRequest(stub, `NearestCentroidEvent`)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else{
        handleBadRequestError(res, event_name)
    }
};
