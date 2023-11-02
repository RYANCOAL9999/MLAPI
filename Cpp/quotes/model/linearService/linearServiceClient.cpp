#include "linearServiceClient.h"
#include <aws/core/utils/json/JsonSerializer.h>
#include <aws/core/utils/memory/stl/SimpleStringStream.h>
#include <../Helper/helper.h>

using grpc::Status;
using grpc::ClientContext;
using AIProto::LinearRegressionRequest;
using AIProto::LinearRidgeRequest;
using AIProto::LinearRidgeCVRequest;
using AIProto::LassoExpressionRequest;
using AIProto::LassoLarsLassoExpressionRequest;
using AIProto::BayesianRidgeRequest;
using AIProto::TweedieRegressorRequest;
using AIProto::SGDClassifierRequest;
using AIProto::ElasticNetRequest;
using AIProto::LinearRegressionkwargs;
using AIProto::LinearRidgekwargs;
using AIProto::LinearRidgeCVkwargs;
using AIProto::LassoExpressionkwargs;
using AIProto::LassoLarsLassoExpressionkwargs;
using AIProto::BayesianRidgekwargs;
using AIProto::TweedieRegressorkwargs;
using AIProto::SGDClassifierkwargs;
using AIProto::ElasticNetkwargs;
using namespace Aws::Utils::Json;

LinearServiceClient::LinearServiceClient(string &url): 
    channel_(CreateChannel(url, InsecureChannelCredentials())),
    stub_(LinearService::NewStub(channel_)) 
{
}

LinearRegressionReply LinearServiceClient::linearRegressionResponse(const string &request_msg)
{
    LinearRegressionReply response;

    ClientContext context;

    LinearRegressionRequest request;
    
    LinearRegressionkwargs kwargs;

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
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_copy_x(body_v.GetBool("copy_x"));
        kwargs.set_n_jobs(body_v.GetInteger("n_jobs"));
        kwargs.set_postive(body_v.GetBool("postive"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->LinearRegressionEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return LinearRegressionReply();
    }

}

LinearRidgeReply LinearServiceClient::ridgeResponse(const string &request_msg)
{
    LinearRidgeReply response;

    ClientContext context;

    LinearRidgeRequest request;
    
    LinearRidgekwargs kwargs;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_allocated_alpha(&Helper::ConvertJsonViewToAny(v, "alpha"));
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_copy_x(body_v.GetBool("copy_x"));
        kwargs.set_allocated_max_iter(&Helper::ConvertJsonViewToAny(body_v, "max_iter"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_solver(static_cast<AIProto::Solver>(body_v.GetInteger("solver")));
        kwargs.set_postive(body_v.GetBool("postive"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->LinearRidgeEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return LinearRidgeReply();
    }
}

LinearRidgeCVReply LinearServiceClient::ridgeCVResponse(const string &request_msg)
{
    LinearRidgeCVReply response;

    ClientContext context;

    LinearRidgeCVRequest request;
    
    LinearRidgeCVkwargs kwargs;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_allocated_alpha(&Helper::ConvertJsonViewToAny(v, "alpha"));
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_allocated_scoring(&Helper::ConvertJsonViewToAny(body_v, "scoring"));
        kwargs.set_allocated_cv(&Helper::ConvertJsonViewToAny(body_v, "cv"));
        kwargs.set_allocated_gcv_mode(&Helper::ConvertJsonViewToAny(body_v, "gcv_mode"));
        kwargs.set_store_cv_values(body_v.GetDouble("store_cv_values"));
        kwargs.set_alpha_per_target(body_v.GetDouble("alpha_per_target"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->LinearRidgeCVEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return LinearRidgeCVReply();
    }
}

LassoExpressionReply LinearServiceClient::lassoResponse(const string &request_msg)
{
    LassoExpressionReply response;

    ClientContext context;

    LassoExpressionRequest request;
    
    LassoExpressionkwargs kwargs;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_alpha(v.GetDouble("alpha"));
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_allocated_precompute(&Helper::ConvertJsonViewToAny(body_v, "precompute"));
        kwargs.set_copy_x(body_v.GetBool("copy_x"));
        kwargs.set_max_iter(body_v.GetInteger("max_iter"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_warm_start(body_v.GetBool("warm_start"));
        kwargs.set_positive(body_v.GetBool("positive"));
        kwargs.set_allocated_random_state(&Helper::ConvertJsonViewToAny(body_v, "random_state"));
        kwargs.set_selection(static_cast<AIProto::Selection>(body_v.GetInteger("selection")));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->LassoExpressionEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return LassoExpressionReply();
    }
}

LassoLarsLassoExpressionReply LinearServiceClient::lassoLarsResponse(const string &request_msg)
{
    LassoLarsLassoExpressionReply response;

    ClientContext context;

    LassoLarsLassoExpressionRequest request;
    
    LassoLarsLassoExpressionkwargs kwargs;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_alpha(v.GetDouble("alpha"));
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_allocated_verbose(&Helper::ConvertJsonViewToAny(body_v, "verbose"));
        kwargs.set_allocated_normalize(&Helper::ConvertJsonViewToAny(body_v, "normalize"));
        kwargs.set_precompute(body_v.GetString("precompute"));
        kwargs.set_max_iter(body_v.GetInteger("max_iter"));
        kwargs.set_eps(body_v.GetBool("eps"));
        kwargs.set_copy_x(body_v.GetBool("copy_x"));
        kwargs.set_fit_path(body_v.GetBool("fit_path"));
        kwargs.set_positive(body_v.GetBool("positive"));
        kwargs.set_jitter(body_v.GetDouble("jitter"));
        kwargs.set_allocated_random_state(&Helper::ConvertJsonViewToAny(body_v, "random_state"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->LassoLarsLassoExpressionEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return LassoLarsLassoExpressionReply();
    }
}

BayesianRidgeReply LinearServiceClient::bayesianRidgeResponse(const string &request_msg)
{
    BayesianRidgeReply response;

    ClientContext context;

    BayesianRidgeRequest request;
    
    BayesianRidgekwargs kwargs;

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
        kwargs.set_max_iter(body_v.GetInteger("max_iter"));
        kwargs.set_tol(body_v.GetInteger("tol"));
        kwargs.set_alpha_1(body_v.GetInteger("alpha_1"));
        kwargs.set_alpha_2(body_v.GetInteger("alpha_2"));
        kwargs.set_lambda_1(body_v.GetInteger("lambda_1"));
        kwargs.set_lambda_2(body_v.GetInteger("lambda_2"));
        kwargs.set_alpha_init(body_v.GetInteger("alpha_init"));
        kwargs.set_lambda_init(body_v.GetInteger("lambda_init"));
        kwargs.set_compute_score(body_v.GetBool("compute_score"));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_copy_x(body_v.GetBool("copy_x"));
        kwargs.set_verbose(body_v.GetBool("verbose"));
        kwargs.set_n_iter(body_v.GetInteger("n_iter"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->BayesianRidgeEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return BayesianRidgeReply();
    }
}

TweedieRegressorReply LinearServiceClient::tweedieRegressorResponse(const string &request_msg)
{
    TweedieRegressorReply response;

    ClientContext context;

    TweedieRegressorRequest request;
    
    TweedieRegressorkwargs kwargs;

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
        kwargs.set_power(body_v.GetDouble("power"));
        kwargs.set_alpha(body_v.GetDouble("alpha"));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_link(static_cast<AIProto::Link>(body_v.GetInteger("link")));
        kwargs.set_solver(static_cast<AIProto::Solver>(body_v.GetInteger("solver")));
        kwargs.set_max_iter(body_v.GetInteger("max_iter"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_start(body_v.GetBool("start"));
        kwargs.set_verbose(body_v.GetInteger("verbose"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->TweedieRegressorEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return TweedieRegressorReply();
    }
}

SGDClassifierReply LinearServiceClient::sgdClassifierResponse(const string &request_msg)
{
    SGDClassifierReply response;

    ClientContext context;

    SGDClassifierRequest request;
    
    AIProto::SGDClassifierkwargs kwargs;

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
        kwargs.set_loss(static_cast<AIProto::Loss>(body_v.GetInteger("loss")));
        kwargs.set_penalty(static_cast<AIProto::Penalty>(body_v.GetInteger("penalty")));
        kwargs.set_alpha(body_v.GetDouble("alpha"));
        kwargs.set_l1_ratio(body_v.GetDouble("l1_ratio"));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_max_iter(body_v.GetInteger("max_iter"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_shuffle(body_v.GetDouble("set_shuffle"));
        kwargs.set_verbose(body_v.GetInteger("verbose"));
        kwargs.set_epsilon(body_v.GetDouble("epsilon"));
        kwargs.set_allocated_n_jobs(&Helper::ConvertJsonViewToAny(body_v, "n_jobs"));
        kwargs.set_allocated_random_state(&Helper::ConvertJsonViewToAny(body_v, "random_state"));
        kwargs.set_learning_rate(body_v.GetString("learning_rate"));
        kwargs.set_eta0(body_v.GetDouble("eta0"));
        kwargs.set_power_t(body_v.GetDouble("power_t"));
        kwargs.set_early_stopping(body_v.GetBool("early_stopping"));
        kwargs.set_validation_fraction(body_v.GetDouble("validation_fraction"));
        kwargs.set_n_iter_no_change(body_v.GetDouble("n_iter_no_change"));
        kwargs.set_allocated_class_weight(&Helper::ConvertJsonViewToAny(body_v, "class_weight"));
        kwargs.set_warm_start(body_v.GetBool("warm_start"));
        kwargs.set_allocated_average(&Helper::ConvertJsonViewToAny(body_v, "average"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->SGDClassifierEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return SGDClassifierReply();
    }
}

ElasticNetReply LinearServiceClient::elasticNetResponse(const string &request_msg)
{
    ElasticNetReply response;

    ClientContext context;

    ElasticNetRequest request;
    
    ElasticNetkwargs kwargs;

    JsonValue json(request_msg);

    if (!json.WasParseSuccessful()) {
        return response;
    }

    JsonView v = json.View();
    request.set_alpha(v.GetDouble("alpha"));
    request.set_allocated_x_drop_data(&Helper::ConvertJsonViewToAny(v, "x_drop_data"));
    request.set_allocated_y_drop_data(&Helper::ConvertJsonViewToAny(v, "y_drop_data"));
    request.set_allocated_size(&Helper::ConvertJsonViewToAny(v, "size"));
    request.set_allocated_random(&Helper::ConvertJsonViewToAny(v, "random"));
    request.set_allocated_key(&v.GetString("key"));
    request.set_allocated_sample_weight(&Helper::ConvertJsonViewToAny(v, "sample_weight"));

    JsonValue json_kwags(v.GetString("kwargs"));
    if (json_kwags.WasParseSuccessful()) {
        JsonView body_v = json_kwags.View();
        kwargs.set_l1_ratio(body_v.GetDouble("l1_ratio"));
        kwargs.set_fit_intercept(body_v.GetBool("fitIntercept"));
        kwargs.set_precompute(body_v.GetBool("precompute"));
        kwargs.set_max_iter(body_v.GetInteger("max_iter"));
        kwargs.set_copy_x(body_v.GetBool("copy_x"));
        kwargs.set_tol(body_v.GetDouble("tol"));
        kwargs.set_warm_start(body_v.GetBool("warm_start"));
        kwargs.set_positive(body_v.GetBool("positive"));
        kwargs.set_random_state(body_v.GetInteger("randomState"));
        kwargs.set_selection(body_v.GetString("selection"));
        request.set_allocated_kwargs(&kwargs);
    }

    Status status = stub_->ElasticNetEvent(&context, request, &response);

    if (status.ok()) {
        return response;
    } else {
        std::cerr << "RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        return ElasticNetReply();
    }
}
