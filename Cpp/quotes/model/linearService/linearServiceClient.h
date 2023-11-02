#ifndef LinearService_Client_H
#define LinearService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/linear_expression.pb.h>
#include <ai-grpc-sdk/linear_expression.grpc.pb.h>

using namespace std;
using namespace grpc;
using AIProto::LinearService;
using AIProto::LinearRegressionReply;
using AIProto::LinearRidgeReply;
using AIProto::LinearRidgeCVReply;
using AIProto::LassoExpressionReply;
using AIProto::LassoLarsLassoExpressionReply;
using AIProto::BayesianRidgeReply;
using AIProto::TweedieRegressorReply;
using AIProto::SGDClassifierReply;
using AIProto::ElasticNetReply;

class LinearServiceClient {

    public:
        LinearServiceClient(string& url);

        LinearRegressionReply linearRegressionResponse(const string& request_msg);

        LinearRidgeReply ridgeResponse(const string& request_msg);

        LinearRidgeCVReply ridgeCVResponse(const string& request_msg);

        LassoExpressionReply lassoResponse(const string& request_msg);

        LassoLarsLassoExpressionReply lassoLarsResponse(const string& request_msg);

        BayesianRidgeReply bayesianRidgeResponse(const string& request_msg);

        TweedieRegressorReply tweedieRegressorResponse(const string& request_msg);

        SGDClassifierReply sgdClassifierResponse(const string& request_msg);

        ElasticNetReply elasticNetResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<LinearService::Stub> stub_;

};

#endif // LinearService_Client_H