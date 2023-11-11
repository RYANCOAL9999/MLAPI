using System;
using System.Threading.Tasks;
using Amazon.Lambda.APIGatewayEvents;
using Amazon.Lambda.Core;
using DotNetServerless.Application.Requests;
using MediatR;
using Microsoft.Extensions.DependencyInjection;
using Newtonsoft.Json;

namespace DotNetServerless.Lambda.Functions
{
    public class GetHeadFunction
    {
        private readonly IServiceProvider _serviceProvider;

        public GetHeadFunction() : this(
            Startup
            .BuildContainer()
            .BuildServiceProvider()
        )
        {
        }

        public GetHeadFunction(IServiceProvider serviceProvider)
        {
            _serviceProvider = serviceProvider;
        }

        [LambdaSerializer(typeof(Amazon.Lambda.Serialization.Json.JsonSerializer))]
        public async Task<APIGatewayProxyResponse> Run(APIGatewayProxyRequest request)
        {
            var requestModel = JsonConvert.DeserializeObject<CreateItemRequest>(request.Body);

            /* think about how to make the _serviceProvide to call the RPC Server*/
            // var mediator = _serviceProvider.GetService<IMediator>();
            // var result = await mediator.Send(requestModel);

            var result = null;

            return new APIGatewayProxyResponse { StatusCode =  200,  Body = JsonConvert.SerializeObject(result)};

        }
    }

}
