const PROTO_PATH = __dirname + '../lib/AIProto/general_expression.proto';
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { defination } = require('./packageDefinitionpackageDefinition')
const { sendStubRequest, handleResponse, handleGRPCError } = require('./helper')

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    defination
);

const general_proto = grpc.loadPackageDefinition(packageDefinition).general_expression;

export const createGeneralStub = (addr) => {
    return new general_proto.GeneralService(
        addr, 
        grpc.credentials.createInsecure()
    );
};

export const headerResponse = (req, res, stub) => {
    const request_message = {}; // Define the structure of the request message here
    sendStubRequest(stub, `HeaderEvent`, request_message)
    .then((response) => handleResponse(response, res))
    .catch((err) => handleGRPCError(err, res, `head`));
};

export const infoResponse = (req, res, stub) => {
    const request_message = {}; // Define the structure of the request message here
    sendStubRequest(stub, `InfoEvent`, request_message)
    .then((response) => handleResponse(response, res))
    .catch((err) => handleGRPCError(err, res, `info`));
};


export const describlerResponse = (req, res, stub) => {
    const request_message = {}; // Define the structure of the request message here
    sendStubRequest(stub, `DescriblerEvent`, request_message)
    .then((response) => handleResponse(response, res))
    .catch((err) => handleGRPCError(err, res, `describle`));
};
