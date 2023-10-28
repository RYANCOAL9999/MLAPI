const PROTO_PATH = __dirname + '../lib/AIProto/svm_expression.proto';
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { defination } = require('./packageDefinitionpackageDefinition')
const { sendStubRequest, handleResponse, handleGRPCError } = require('./helper')

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    defination
);

const svm_proto = grpc.loadPackageDefinition(packageDefinition).svm_expression;

export const createSVMStub = (addr) => {
    return new svm_proto.SVMService(
        addr, 
        grpc.credentials.createInsecure()
    );
};

export const linearSVCResponse = (req, res, stub) => {

    const event_name = `linear svc`

    const request_message = req.body; 

    if (request_message) {
        sendStubRequest(stub, "LinearSVCEvent", request_message)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    } 
    else {
        handleBadRequestError(res, event_name);
    }

};

export const linearSVRResponse = (req, res, stub) => {

    const event_name = `linear svr`

    const request_message = req.body; 

    if (request_message) {
        sendStubRequest(stub, `LinearSVREvent`, request_message)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else {
        handleBadRequestError(res, event_name);
    }
};


export const svcResponse = (req, res, stub) => {

    const event_name = `svc`

    const request_message = req.body;

    if (request_message) {
        sendStubRequest(stub, `SVCEvent`, request_message)
        .then((response) => handleResponse(response, res))
        .catch((err) => handleGRPCError(err, res, event_name));
    }
    else {
        handleBadRequestError(res, event_name);
    }
};
