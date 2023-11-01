#include "linearServiceClient.h"

LinearServiceClient::LinearServiceClient(string &url): 
    channel_(grpc::CreateChannel(url, grpc::InsecureChannelCredentials())),
    stub_(AIProto::LinearService::NewStub(channel_)) 
{
}

AIProto::LinearRegressionReply LinearServiceClient::linearRegressionResponse(const string &request_msg)
{
    return AIProto::LinearRegressionReply();
}

AIProto::LinearRidgeReply LinearServiceClient::ridgeResponse(const string &request_msg)
{
    return AIProto::LinearRidgeReply();
}

AIProto::LinearRidgeCVReply LinearServiceClient::ridgeCVResponse(const string &request_msg)
{
    return AIProto::LinearRidgeCVReply();
}

AIProto::LassoExpressionReply LinearServiceClient::lassoResponse(const string &request_msg)
{
    return AIProto::LassoExpressionReply();
}

AIProto::LassoLarsLassoExpressionReply LinearServiceClient::lassoLarsResponse(const string &request_msg)
{
    return AIProto::LassoLarsLassoExpressionReply();
}

AIProto::BayesianRidgeReply LinearServiceClient::bayesianRidgeResponse(const string &request_msg)
{
    return AIProto::BayesianRidgeReply();
}

AIProto::TweedieRegressorReply LinearServiceClient::tweedieRegressorResponse(const string &request_msg)
{
    return AIProto::TweedieRegressorReply();
}

AIProto::SGDClassifierReply LinearServiceClient::sgdClassifierResponse(const string &request_msg)
{
    return AIProto::SGDClassifierReply();
}

AIProto::ElasticNetReply LinearServiceClient::elasticNetResponse(const string &request_msg)
{
    return AIProto::ElasticNetReply();
}
