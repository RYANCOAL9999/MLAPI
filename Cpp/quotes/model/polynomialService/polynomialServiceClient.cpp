#include "polynomialServiceClient.h"
#include <aws/core/utils/json/JsonSerializer.h>
#include <aws/core/utils/memory/stl/SimpleStringStream.h>
#include <../Helper/helper.h>

using grpc::Status;
using grpc::ClientContext;
using AIProto::PolynomialFeaturesRequest;
using AIProto::PolynomialFeaturesFitTransformkwargs;
using AIProto::PolynomialFeaturesFitTransformRequest;
using namespace Aws::Utils::Json;

PolynomialServiceClient::PolynomialServiceClient(string &url): 
    channel_(CreateChannel(url, InsecureChannelCredentials())),
    stub_(PolynomialService::NewStub(channel_)) 
{
}

PolynomialFeaturesReply PolynomialServiceClient::polynomialFeaturesResponse(const string &request_msg)
{
    PolynomialFeaturesReply response;

    ClientContext context;

    PolynomialFeaturesRequest request;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_key(v.GetString("key"));
    request.set_interaction_only(&Helper::ConvertJsonViewToAny(v, "interaction_only"));
    request.set_include_bias(&Helper::ConvertJsonViewToAny(v, "include_bias"));
    request.set_order(static_cast<AIProto::Order>(v.GetInteger("order")));
    request.set_allocated_kwargs(&Helper::ConvertJsonViewToAny(v, "kwargs"));

    Status status = stub_->PolynomialFeaturesEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return PolynomialFeaturesReply();
    }
}

PolynomialFeaturesFitTransformReply PolynomialServiceClient::polynomialFeaturesFitTransformResponse(const string &request_msg)
{
    PolynomialFeaturesFitTransformReply response;

    ClientContext context;

    PolynomialFeaturesFitTransformRequest request;

    PolynomialFeaturesFitTransformkwargs kwargs;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_key(v.GetString("key"));
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));
    request.set_degree(v.GetInteger("degree"));
    request.set_interaction_only(&Helper::ConvertJsonViewToAny(v, "interaction_only"));
    request.set_include_bias(&Helper::ConvertJsonViewToAny(v, "include_bias"));
    request.set_order(static_cast<AIProto::Order>(v.GetInteger("order")));
    request.set_allocated_coef_init(&Helper::ConvertJsonViewToAny(v, "coef_init"));
    request.set_allocated_intercept_init(&Helper::ConvertJsonViewToAny(v, "intercept_init"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        Aws::Utils::Json::JsonView body_v = json_kwags.View();
        kwargs.set_penalty(static_cast<AIProto::Penalty>(body_v.GetInteger("penalty")));
        kwargs.set_alpha(body_v.GetDouble("alpha"));
        kwargs.set_l1_ratio(body_v.GetDouble("l1_ratio"));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_max_iter(body_v.GetInt64("maxIter"));
        kwargs.set_allocated_tol(&Helper::ConvertJsonViewToAny(body_v, "tol"));
        kwargs.set_shuffle(body_v.GetBool("shuffle"));
        kwargs.set_verbose(body_v.GetInteger("verbose"));
        kwargs.set_eta0(body_v.GetDouble("eta0"));
        kwargs.set_allocated_n_jobs(&Helper::ConvertJsonViewToAny(body_v, "n_jobs"));
        kwargs.set_allocated_random_state(&Helper::ConvertJsonViewToAny(body_v, "randomState"));
        kwargs.set_early_stopping(body_v.GetBool("early_stopping"));
        kwargs.set_validation_fraction(body_v.GetBool("validation_fraction"));
        kwargs.set_n_iter_no_change(body_v.GetInteger("n_iter_no_change"));
        kwargs.set_allocated_class_weight(&Helper::ConvertJsonViewToAny(body_v, "classWeight"));
        kwargs.set_warm_start(body_v.GetBool("warm_start"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->PolynomialFeaturesFitTransformEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return PolynomialFeaturesFitTransformReply();
    }
}
