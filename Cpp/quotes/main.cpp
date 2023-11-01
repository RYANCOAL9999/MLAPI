#include <iostream>
#include "crow.h"
#include <aws/lambda-runtime/runtime.h>
// #include <aws/core/utils/json/JsonSerializer.h>
// #include <aws/core/utils/memory/stl/SimpleStringStream.h>
#include <model/generalService/generalServiceClient.h>
#include <model/linearService/linearServiceClient.h>
#include <model/neighborsService/neighborsServiceClient.h>
#include <model/polynomialService/polynomialServiceClient.h>
#include <model/svmService/svmServiceClient.h>

using namespace std;
using namespace crow;
using namespace aws::lambda_runtime;

void setupRoutes(
    crow::SimpleApp& app, 
    GeneralServiceClient& gsclient,
    LinearServiceClient& lsclient,
    NeighborsServiceClient& nsclient,
    PolynomialServiceClient& psclient,
    SVMServiceClient& svmClient
) 
{
    CROW_ROUTE(app, "/api/head")
    .methods("GET"_method)
    ([&gsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        DataFrame response = gsclient.headerResponse("");
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/info")
    .methods("GET"_method)
    ([&gsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        DataFrame response = gsclient.infoResponse("");
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/describle")
    .methods("GET"_method)
    ([&gsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        DataFrame response = gsclient.describlerResponse("");
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/linearSVC")
    .methods("POST"_method)
    ([&svmClient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::LinearReply response = svmClient.linearSVCResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/linearSVR")
    .methods("POST"_method)
    ([&svmClient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::LinearReply response = svmClient.linearSVRResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/svc")
    .methods("POST"_method)
    ([&svmClient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::SVCReply response = svmClient.svcResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/polynomialFeatures")
    .methods("POST"_method)
    ([&psclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::PolynomialFeaturesReply response = psclient.polynomialFeaturesResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/polynomialFeaturesFitTransform")
    .methods("POST"_method)
    ([&psclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::PolynomialFeaturesFitTransformReply response = psclient.polynomialFeaturesFitTransformResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/nearestNeighbors")
    .methods("POST"_method)
    ([&nsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::NearestNeighborsReply response = nsclient.nearestNeighborsResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/kdTree")
    .methods("POST"_method)
    ([&nsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::KDTreeReply response = nsclient.kdTreeResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/nearestCentroid")
    .methods("POST"_method)
    ([&nsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::NearestCentroidReply response = nsclient.nearestCentroidResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/linearRegression")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::LinearRegressionReply response = lsclient.linearRegressionResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/ridge")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::LinearRidgeReply response = lsclient.ridgeResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/ridgeCV")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::LinearRidgeCVReply response = lsclient.ridgeCVResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/lasso")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::LassoExpressionReply response = lsclient.lassoResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/lassoLars")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::LassoLarsLassoExpressionReply response = lsclient.lassoLarsResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/bayesianRidge")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::BayesianRidgeReply response = lsclient.bayesianRidgeResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/tweedieRegressor")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::TweedieRegressorReply response = lsclient.tweedieRegressorResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });


    CROW_ROUTE(app, "/api/sgdClassifier")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::SGDClassifierReply response = lsclient.sgdClassifierResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });

    CROW_ROUTE(app, "/api/elasticNet")
    .methods("POST"_method)
    ([&lsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        AIProto::ElasticNetReply response = lsclient.elasticNetResponse(req.body);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });
}

invocation_response my_handler(invocation_request const& request)
{
    string server_address_1 = "localhost:50051";
    string server_address_2 = "localhost:50052";
    string server_address_3 = "localhost:50053";
    string server_address_4 = "localhost:50054";

    GeneralServiceClient gsclient(server_address_1);
    LinearServiceClient lsclient(server_address_1);
    NeighborsServiceClient nsclient(server_address_2);
    PolynomialServiceClient psclient(server_address_3);
    SVMServiceClient svmClient(server_address_4);

    SimpleApp app;

    setupRoutes(
        app, 
        gsclient,
        lsclient,
        nsclient,
        psclient,
        svmClient
    );

    app.port(8080).multithreaded().run();

    return invocation_response::success(
        "Marning Learning API starting Success",
        "application/json"
    );
}

int main()
{
    run_handler(my_handler);
    return 0;
}