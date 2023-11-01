#include "neighborsServiceClient.h"
#include <aws/core/utils/json/JsonSerializer.h>
#include <aws/core/utils/memory/stl/SimpleStringStream.h>
#include <../Helper/helper.h>

using grpc::Status;
using grpc::ClientContext;
using AIProto::NearestNeighborsReply;
using AIProto::KDTreeReply;
using AIProto::NearestCentroidReply;
using namespace Aws::Utils::Json;

NeighborsServiceClient::NeighborsServiceClient(string &url): 
    channel_(grpc::CreateChannel(url, grpc::InsecureChannelCredentials())),
    stub_(AIProto::NeighborsService::NewStub(channel_)) 
{
}

NearestNeighborsReply NeighborsServiceClient::nearestNeighborsResponse(const string &request_msg)
{
    NearestNeighborsReply response;

    ClientContext context;

    AIProto::NearestNeighborsRequest request;
    
    AIProto::NearestNeighborskwargs kwargs;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_n_neighbors(body_v.GetInteger("n_neighbors"));
        kwargs.set_radius(body_v.GetDouble("radius"));
        kwargs.set_algorithm(static_cast<AIProto::Algorithm>(body_v.GetInteger("algorithm")));
        kwargs.set_leaf_size(body_v.GetInteger("leaf_size"));
        kwargs.set_allocated_metric(&Helper::ConvertJsonViewToAny(v, "metric"));
        kwargs.set_p(body_v.GetDouble("p"));
        kwargs.set_allocated_metric_params(&Helper::ConvertJsonViewToAny(v, "metric_params"));
        kwargs.set_allocated_n_jobs(&Helper::ConvertJsonViewToAny(body_v, "n_jobs"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->NearestNeighborsEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return NearestNeighborsReply();
    }
}
KDTreeReply NeighborsServiceClient::kdTreeResponse(const string &request_msg)
{
    KDTreeReply response;

    ClientContext context;

    AIProto::KDTreeRequest request;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));
    request.set_k(v.GetInteger("k"));
    request.set_leaf_size(v.GetInteger("leaf_size"));
    request.set_metric(Helper::ConvertJsonViewToAny(v, "metric"));
    request.set_returns_distance(v.GetBool("return_distance")); 
    request.set_dualtree(v.GetBool("dualtree"));
    request.set_breadth_first(v.GetBool("breadth_first"));
    request.set_sort_results(v.GetBool("sort_results"));
    request.set_allocated_kwargs(&Helper::ConvertJsonViewToAny(v, "kwags"));

    Status status = stub_->KDTreeEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return KDTreeReply();
    }
}

NearestCentroidReply NeighborsServiceClient::nearestCentroidResponse(const string &request_msg)
{
    NearestCentroidReply response;

    ClientContext context;

    AIProto::NearestCentroidRequest request;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));
    request.set_allocated_metric(&Helper::ConvertJsonViewToAny(v, "metric"));
    request.set_shrink_threshold(v.GetDouble("shrink_threshold"));

    Status status = stub_->NearestCentroidEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return NearestCentroidReply();
    }
}
