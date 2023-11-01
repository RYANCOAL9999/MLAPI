#include "svmServiceClient.h"

SVMServiceClient::SVMServiceClient(string &url): 
    channel_(grpc::CreateChannel(url, grpc::InsecureChannelCredentials())),
    stub_(AIProto::SVMService::NewStub(channel_)) 
{
}

AIProto::LinearReply SVMServiceClient::linearSVCResponse(const string &request_msg)
{
    return AIProto::LinearReply();
}

AIProto::LinearReply SVMServiceClient::linearSVRResponse(const string &request_msg)
{
    return AIProto::LinearReply();
}

AIProto::SVCReply SVMServiceClient::svcResponse(const string &request_msg)
{
    return AIProto::SVCReply();
}
