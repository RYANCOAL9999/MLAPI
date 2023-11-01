#ifndef LinearService_Client_H
#define LinearService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/linear_expression.pb.h>
#include <ai-grpc-sdk/linear_expression.grpc.pb.h>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using namespace std;

class LinearServiceClient {

    public:
        LinearServiceClient(string& url);

        AIProto::LinearRegressionReply linearRegressionResponse(const string& request_msg);

        AIProto::LinearRidgeReply ridgeResponse(const string& request_msg);

        AIProto::LinearRidgeCVReply ridgeCVResponse(const string& request_msg);

        AIProto::LassoExpressionReply lassoResponse(const string& request_msg);

        AIProto::LassoLarsLassoExpressionReply lassoLarsResponse(const string& request_msg);

        AIProto::BayesianRidgeReply bayesianRidgeResponse(const string& request_msg);

        AIProto::TweedieRegressorReply tweedieRegressorResponse(const string& request_msg);

        AIProto::SGDClassifierReply sgdClassifierResponse(const string& request_msg);

        AIProto::ElasticNetReply elasticNetResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<AIProto::LinearService::Stub> stub_;

};

#endif // LinearService_Client_H