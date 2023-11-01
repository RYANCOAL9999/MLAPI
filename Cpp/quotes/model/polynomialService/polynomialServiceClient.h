#ifndef PolynomialService_Client_H
#define PolynomialService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/polynomial_features.pb.h>
#include <ai-grpc-sdk/polynomial_features.grpc.pb.h>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using namespace std;

class PolynomialServiceClient {

    public:
        PolynomialServiceClient(string& url);

        AIProto::PolynomialFeaturesReply polynomialFeaturesResponse(const string& request_msg);

        AIProto::PolynomialFeaturesFitTransformReply polynomialFeaturesFitTransformResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<AIProto::PolynomialService::Stub> stub_;

};



#endif // PolynomialService_Client_H