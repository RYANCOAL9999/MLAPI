#ifndef SVMService_Client_H
#define SVMService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/svm_expression.pb.h>
#include <ai-grpc-sdk/svm_expression.grpc.pb.h>

using namespace std;
using namespace grpc;
using AIProto::SVCReply;
using AIProto::SVMService;
using AIProto::LinearReply;


class SVMServiceClient {

    public:

        SVMServiceClient(string& url);

        LinearReply linearSVCResponse(const string &request_msg);

        LinearReply linearSVRResponse(const string& request_msg);

        SVCReply svcResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<SVMService::Stub> stub_;

};

#endif // SVMService_Client_H