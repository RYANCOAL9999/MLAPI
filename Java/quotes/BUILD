load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_grpc_grpc_java//:java_grpc_library.bzl", "java_grpc_library")

proto_library(
    name = "any_proto",
    srcs = ["main/AIProto/google/protobuf/any.proto"],
)

proto_library(
    name = "enum_expression",
    srcs = ["main/AIProto/lib/enum_expression.proto"],
)

proto_library(
    name = "np_library",
    srcs = ["main/AIProto/lib/np_library.proto"],
    deps = [":any_proto"],
)

proto_library(
    name = "general_expression",
    srcs = ["main/AIProto/general_expression.proto"],
    deps = [":any_proto"],
)

proto_library(
    name = "file_handler",
    srcs = ["main/AIProto/file_handler.proto"],
)

proto_library(
    name = "linear_expression",
    srcs = ["main/AIProto/linear_expression.proto"],
    deps = [
        ":any_proto",
        ":np_library",
        ":enum_expression"
    ],
)

proto_library(
    name = "nearest_neighbors",
    srcs = ["main/AIProto/nearest_neighbors.proto"],
    deps = [
        ":any_proto",
        ":np_library",
        ":enum_expression"
    ],
)

proto_library(
    name = "polynomial_features",
    srcs = ["main/AIProto/polynomial_features.proto"],
    deps = [
        ":any_proto",
        ":np_library",
        ":enum_expression"
    ],
)

proto_library(
    name = "svm_expression",
    srcs = ["main/AIProto/svm_expression.proto"],
    deps = [
        ":any_proto",
        ":np_library",
        ":enum_expression"
    ],
)

java_proto_library(
    name = "general_expression_pb",
    deps = [":general_expression"],
)

java_proto_library(
    name = "file_handler_pb",
    deps = [":file_handler"],
)

java_proto_library(
    name = "linear_expression_pb",
    deps = [":linear_expression"],
)

java_proto_library(
    name = "nearest_neighbors_pb",
    deps = [":nearest_neighbors"],
)

java_proto_library(
    name = "polynomial_features_pb",
    deps = [":polynomial_features"],
)

java_proto_library(
    name = "svm_expression_pb",
    deps = [":svm_expression"],
)

java_grpc_library(
    name = "general_expression_grpc",
    srcs = [":general_expression"],
    deps = [":general_expression_pb"],
)

java_grpc_library(
    name = "file_handler_grpc",
    srcs = [":file_handler"],
    deps = [":file_handler_pb"],
)

java_grpc_library(
    name = "linear_expression_grpc",
    srcs = [":linear_expression"],
    deps = [":linear_expression_pb"],
)

java_grpc_library(
    name = "nearest_neighbors_grpc",
    srcs = [":nearest_neighbors"],
    deps = [":nearest_neighbors_pb"],
)

java_grpc_library(
    name = "polynomial_features_grpc",
    srcs = [":polynomial_features"],
    deps = [":polynomial_features_pb"],
)

java_grpc_library(
    name = "svm_expression_grpc",
    srcs = [":svm_expression"],
    deps = [":svm_expression_pb"],
)
