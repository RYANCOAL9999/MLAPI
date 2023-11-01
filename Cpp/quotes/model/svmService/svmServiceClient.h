#ifndef SVMService_Client_H
#define SVMService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/svm_expression.pb.h>
#include <ai-grpc-sdk/svm_expression.grpc.pb.h>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using namespace std;

class SVMServiceClient {

    public:

        SVMServiceClient(string& url);

        AIProto::LinearReply linearSVCResponse(const string& request_msg);

        AIProto::LinearReply linearSVRResponse(const string& request_msg);

        AIProto::SVCReply svcResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<AIProto::SVMService::Stub> stub_;

};

#endif // SVMService_Client_H