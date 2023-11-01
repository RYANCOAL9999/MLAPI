#ifndef GeneralService_Client_H
#define GeneralService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/general_expression.pb.h>
#include <ai-grpc-sdk/general_expression.grpc.pb.h>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using AIProto::DataFrame;
using namespace std;

class GeneralServiceClient {

    public:
        GeneralServiceClient(string& url);

        DataFrame headerResponse(const string& request_msg);

        DataFrame infoResponse(const string& request_msg);

        DataFrame describlerResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<AIProto::GeneralService::Stub> stub_;

};



#endif // GeneralService_Client_H