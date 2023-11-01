#ifndef NeighborsService_Client_H
#define NeighborsService_Client_H

#include "string"
#include <grpc++/grpc++.h>
#include <ai-grpc-sdk/nearest_neighbors.pb.h>
#include <ai-grpc-sdk/nearest_neighbors.grpc.pb.h>

using namespace std;

class NeighborsServiceClient {

    public:

        NeighborsServiceClient(string& url);

        AIProto::NearestNeighborsReply nearestNeighborsResponse(const string& request_msg);

        AIProto::KDTreeReply kdTreeResponse(const string& request_msg);

        AIProto::NearestCentroidReply nearestCentroidResponse(const string& request_msg);

    private:
        shared_ptr<grpc::Channel> channel_;
        unique_ptr<AIProto::NeighborsService::Stub> stub_;
};

#endif // NeighborsService_Client_H