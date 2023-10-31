#include "generalServiceClient.h"
#include "json/json.h"

using namespace std;

GeneralServiceClient::GeneralServiceClient(string& url): 
    channel_(grpc::CreateChannel(url, grpc::InsecureChannelCredentials())),
    stub_(GeneralService::NewStub(channel_)) 
{
}

DataFrame GeneralServiceClient::headerResponse(const string& request_msg) 
{

    AIProto::HeaderRequest request;

    DataFrame response;
    ClientContext context;

    Status status = stub_->HeaderEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return DataFrame();
    }
}

DataFrame GeneralServiceClient::infoResponse(const string& request_msg) 
{

    AIProto::InfoRequest request;

    DataFrame response;
    ClientContext context;

    Status status = stub_->InfoEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return DataFrame();
    }
}

DataFrame GeneralServiceClient::describlerResponse(const string& request_msg) 
{

    AIProto::DescribeRequest request;

    DataFrame response;
    ClientContext context;

    Status status = stub_->DescriblerEvent(&context, request, &response);

    if (status.ok()) {
        return DataFrame(response);
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return DataFrame();
    }
}

