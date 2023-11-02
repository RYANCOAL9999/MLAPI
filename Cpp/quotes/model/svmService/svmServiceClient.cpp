#include "svmServiceClient.h"
#include <aws/core/utils/json/JsonSerializer.h>
#include <aws/core/utils/memory/stl/SimpleStringStream.h>
#include <../Helper/helper.h>

using grpc::Status;
using grpc::ClientContext;
using AIProto::SVCRequest;
using AIProto::Linearkwargs;
using AIProto::LinearRequest;
using namespace Aws::Utils::Json;

SVMServiceClient::SVMServiceClient(string &url): 
    channel_(CreateChannel(url, InsecureChannelCredentials())),
    stub_(SVMService::NewStub(channel_)) 
{
}

LinearReply SVMServiceClient::linearSVCResponse(const string &request_msg)
{

    LinearReply response;

    ClientContext context;

    LinearRequest request;
    
    Linearkwargs kwargs;

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
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));
    
    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_penalty(static_cast<AIProto::Penalty>(body_v.GetInteger("penalty")));
        kwargs.set_loss(static_cast<AIProto::Loss>(body_v.GetInteger("loss")));
        kwargs.set_dual(body_v.GetBool("dual"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_c(body_v.GetDouble("c"));
        kwargs.set_multi_class(static_cast<AIProto::Multi_Class>(body_v.GetInteger("multiClass")));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_intercept_scaling(body_v.GetDouble("interceptScaling"));
        kwargs.set_allocated_class_weight(&Helper::ConvertJsonViewToAny(body_v, "classWeight"));
        kwargs.set_verbose(body_v.GetInteger("verbose"));
        kwargs.set_allocated_random_state(&Helper::ConvertJsonViewToAny(body_v, "randomState"));
        kwargs.set_max_iter(body_v.GetInt64("maxIter"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->LinearSVCEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return LinearReply();
    }
}

LinearReply SVMServiceClient::linearSVRResponse(const string &request_msg)
{
    LinearReply response;

    ClientContext context;

    LinearRequest request;
    
    Linearkwargs kwargs;

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
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_penalty(static_cast<AIProto::Penalty>(body_v.GetInteger("penalty")));
        kwargs.set_loss(static_cast<AIProto::Loss>(body_v.GetInteger("loss")));
        kwargs.set_dual(body_v.GetBool("dual"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_c(body_v.GetDouble("c"));
        kwargs.set_multi_class(static_cast<AIProto::Multi_Class>(body_v.GetInteger("multiClass")));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_intercept_scaling(body_v.GetDouble("interceptScaling"));
        kwargs.set_allocated_class_weight(&Helper::ConvertJsonViewToAny(body_v, "classWeight"));
        kwargs.set_verbose(body_v.GetInteger("verbose"));
        kwargs.set_allocated_random_state(&Helper::ConvertJsonViewToAny(body_v, "randomState"));
        kwargs.set_max_iter(body_v.GetInt64("maxIter"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->LinearSVREvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return LinearReply();
    }
}

SVCReply SVMServiceClient::svcResponse(const string &request_msg)
{
    SVCReply response;

    ClientContext context;

    SVCRequest request;

    AIProto::SVCkwargs kwargs;

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
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_penalty(static_cast<AIProto::Penalty>(body_v.GetInteger("penalty")));
        kwargs.set_loss(static_cast<AIProto::Loss>(body_v.GetInteger("loss")));
        kwargs.set_dual(body_v.GetBool("dual"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_c(body_v.GetDouble("c"));
        kwargs.set_multi_class(static_cast<AIProto::Multi_Class>(body_v.GetInteger("multiClass")));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_intercept_scaling(body_v.GetDouble("interceptScaling"));
        kwargs.set_allocated_class_weight(&Helper::ConvertJsonViewToAny(body_v, "classWeight"));
        kwargs.set_verbose(body_v.GetInteger("verbose"));
        kwargs.set_allocated_random_state(&Helper::ConvertJsonViewToAny(body_v, "randomState"));
        kwargs.set_max_iter(body_v.GetInt64("maxIter"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->SVCEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return SVCReply();
    }
}
