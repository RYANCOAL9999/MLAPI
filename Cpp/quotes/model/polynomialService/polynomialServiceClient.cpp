#include "polynomialServiceClient.h"

PolynomialServiceClient::PolynomialServiceClient(string &url): 
    channel_(grpc::CreateChannel(url, grpc::InsecureChannelCredentials())),
    stub_(AIProto::PolynomialService::NewStub(channel_)) 
{
}

AIProto::PolynomialFeaturesReply PolynomialServiceClient::polynomialFeaturesResponse(const string &request_msg)
{
    return AIProto::PolynomialFeaturesReply();
}

AIProto::PolynomialFeaturesFitTransformReply PolynomialServiceClient::polynomialFeaturesFitTransformResponse(const string &request_msg)
{
    return AIProto::PolynomialFeaturesFitTransformReply();
}
