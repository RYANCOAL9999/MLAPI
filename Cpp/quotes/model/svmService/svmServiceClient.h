#ifndef SVMService_Client_H
#define SVMService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/svm_expression.pb.h>
#include <ai-grpc-sdk/svm_expression.grpc.pb.h>

using namespace std;

class SVMServiceClient {

    public:

        SVMServiceClient(string& url);

        google::protobuf::Any ConvertJsonViewToAny(const Aws::Utils::Json::JsonView &jsonView, string name);

        AIProto::LinearReply linearSVCResponse(const string &request_msg);

        AIProto::LinearReply linearSVRResponse(const string& request_msg);

        AIProto::SVCReply svcResponse(const string& request_msg);

    private:
        shared_ptr<grpc::Channel> channel_;
        unique_ptr<AIProto::SVMService::Stub> stub_;

};

#endif // SVMService_Client_H