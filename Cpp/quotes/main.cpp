#include <iostream>
#include "crow.h"
#include <aws/lambda-runtime/runtime.h>
#include <model/generalService/generalService.h>
// #include <aws/core/utils/json/JsonSerializer.h>
// #include <aws/core/utils/memory/stl/SimpleStringStream.h>

using namespace std;
using namespace crow;
using namespace aws::lambda_runtime;

void setupRoutes(
    crow::SimpleApp& app, 
    GeneralServiceClient& gsclient
) 
{
    CROW_ROUTE(app, "/api/head")
    .methods("GET"_method)
    ([&gsclient](
        const crow::request& req, 
        crow::response& res
    ) 
    {
        string request_msg = "";
        DataFrame response = gsclient.headerResponse(request_msg);
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
        string request_msg = "";
        DataFrame response = gsclient.infoResponse(request_msg);
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
        string request_msg = "";
        DataFrame response = gsclient.describlerResponse(request_msg);
        res.set_header("Content-Type", "application/json");
        res.write(response.SerializeAsString());
        res.end();
    });




}

invocation_response my_handler(invocation_request const& request)
{
    string server_address_1 = "localhost:50051";
    GeneralServiceClient gsclient(server_address_1);

    SimpleApp app;

    setupRoutes(
        app, 
        gsclient
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