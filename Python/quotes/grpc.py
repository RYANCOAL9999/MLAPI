# import generated_pb # Import your gRPC generated files
# import generated_pb_grpc

def rpc(channel, **kwargs):
    print("Will try to greet world ...")
    with channel:
        # stub = generated_pb_grpc.GreeterStub(channel)
        # response = stub.SayHello(generated_pb.Request(**kwargs))
        pass
    print(
        "Greeter client received: " 
        # + response.message
    )


