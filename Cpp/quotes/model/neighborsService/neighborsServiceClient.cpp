#include "neighborsServiceClient.h"

NeighborsServiceClient::NeighborsServiceClient(string &url): 
    channel_(grpc::CreateChannel(url, grpc::InsecureChannelCredentials())),
    stub_(AIProto::NeighborsService::NewStub(channel_)) 
{
}

AIProto::NearestNeighborsReply NeighborsServiceClient::nearestNeighborsResponse(const string &request_msg)
{
    return AIProto::NearestNeighborsReply();
}
AIProto::KDTreeReply NeighborsServiceClient::kdTreeResponse(const string &request_msg)
{
    return AIProto::KDTreeReply();
}

AIProto::NearestCentroidReply NeighborsServiceClient::nearestCentroidResponse(const string &request_msg)
{
    return AIProto::NearestCentroidReply();
}
