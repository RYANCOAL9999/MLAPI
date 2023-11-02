#ifndef NeighborsService_Client_H
#define NeighborsService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/nearest_neighbors.pb.h>
#include <ai-grpc-sdk/nearest_neighbors.grpc.pb.h>

using namespace std;
using namespace grpc;
using AIProto::NeighborsService;
using AIProto::NearestNeighborsReply;
using AIProto::KDTreeReply;
using AIProto::NearestCentroidReply;

class NeighborsServiceClient {

    public:

        NeighborsServiceClient(string& url);

        NearestNeighborsReply nearestNeighborsResponse(const string& request_msg);

        KDTreeReply kdTreeResponse(const string& request_msg);

        NearestCentroidReply nearestCentroidResponse(const string& request_msg);

    private:
        shared_ptr<Channel> channel_;
        unique_ptr<NeighborsService::Stub> stub_;
};

#endif // NeighborsService_Client_H