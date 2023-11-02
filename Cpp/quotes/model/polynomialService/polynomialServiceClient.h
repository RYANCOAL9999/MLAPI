#ifndef PolynomialService_Client_H
#define PolynomialService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/polynomial_features.pb.h>
#include <ai-grpc-sdk/polynomial_features.grpc.pb.h>

using namespace std;
using namespace grpc;
using AIProto::PolynomialService;
using AIProto::PolynomialFeaturesReply;
using AIProto::PolynomialFeaturesFitTransformReply;

class PolynomialServiceClient {

    public:
        PolynomialServiceClient(string& url);

        PolynomialFeaturesReply polynomialFeaturesResponse(const string& request_msg);

        PolynomialFeaturesFitTransformReply polynomialFeaturesFitTransformResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<PolynomialService::Stub> stub_;

};



#endif // PolynomialService_Client_H