load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_grpc_grpc_java//:java_grpc_library.bzl", "java_grpc_library")

proto_library(
    name = "any_proto",
    srcs = ["main/AIProto/google/protobuf/any.proto"],
)

proto_library(
    name = "general_expression",
    srcs = ["main/AIProto/general_expression.proto"],
    deps = [":any_proto"],
)

java_proto_library(
    name = "general_expression_pb",
    deps = [":general_expression"],
)

java_grpc_library(
    name = "general_expression_grpc",
    srcs = [":general_expression"],
    deps = [":general_expression_pb"],
)
