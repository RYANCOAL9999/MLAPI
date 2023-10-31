cmake_minimum_required(VERSION 3.8)

project(quotes C CXX)

include(../cmake/common.cmake)

# Proto file
get_filename_component(gs_proto "../../../../../../../AIProto/general_expression.proto" ABSOLUTE)
get_filename_component(gs_proto_path "${gs_proto}" PATH)

# Generated sources
set(gs_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/general_expression.pb.cc")
set(gs_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/general_expression.pb.h")
set(gs_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/general_expression.grpc.pb.cc")
set(gs_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/general_expression.grpc.pb.h")

add_custom_command(
    OUTPUT "${gs_proto_srcs}" "${gs_proto_hdrs}" "${gs_grpc_srcs}" "${gs_grpc_hdrs}"
    COMMAND ${_PROTOBUF_PROTOC}
    ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
    --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
    -I "${gs_proto_path}"
    --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
    "${gs_proto}"
    DEPENDS "${gs_proto}"
)

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# gs_grpc_proto
add_library(
    gs_grpc_proto
    ${gs_grpc_srcs}
    ${gs_grpc_hdrs}
    ${gs_proto_srcs}
    ${gs_proto_hdrs}
)

target_link_libraries(
    gs_grpc_proto
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF}
)

# add_library(
#   helper
#   "helper.h"
#   "helper.cc"
# )
# target_link_libraries(
#   helper
#   gs_grpc_proto
#   ${_REFLECTION}
#   ${_GRPC_GRPCPP}
#   ${_PROTOBUF_LIBPROTOBUF}
# )

# Targets route_guide_(client|server)
foreach(
    _target
    generalServiceClient
    # generalServiceServer
)
    add_executable(
        ${_target}
        "${_target}.cc"
    )
    target_link_libraries(
        ${_target}
        gs_grpc_proto
        # helper
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF}
    )
endforeach()














