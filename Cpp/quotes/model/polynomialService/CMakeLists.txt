cmake_minimum_required(VERSION 3.8)

project(quotes C CXX)

include(../cmake/common.cmake)

# Proto file
get_filename_component(ls_proto "../../../../../../../AIProto/lib/enum_expression.proto" ABSOLUTE)
get_filename_component(ls_proto "../../../../../../../AIProto/lib/np_library.proto" ABSOLUTE)
get_filename_component(ps_proto "../../../../../../../AIProto/polynomial_features.proto" ABSOLUTE)
get_filename_component(ps_proto_path "${ps_proto}" PATH)

# Generated sources
set(enum_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.cc")
set(enum_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.h")
set(np_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.cc")
set(np_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.h")
set(ps_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/polynomial_features.pb.cc")
set(ps_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/polynomial_features.pb.h")
set(ps_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/polynomial_features.grpc.pb.cc")
set(ps_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/polynomial_features.grpc.pb.h")

add_custom_command(
    OUTPUT "${enum_proto_srcs}" "${enum_proto_hdrs}" "${np_proto_srcs}" "${np_proto_srcs}" "${ps_proto_srcs}" "${ps_proto_hdrs}" "${ps_grpc_srcs}" "${ps_grpc_hdrs}"
    COMMAND ${_PROTOBUF_PROTOC}
    ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
    --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
    -I "${ps_proto_path}"
    --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
    "${ps_proto}"
    DEPENDS "${ps_proto}"
)

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# ps_grpc_proto
add_library(
    ps_grpc_proto
    ${enum_proto_srcs}
    ${enum_proto_hdrs} 
    ${np_proto_srcs} 
    ${np_proto_srcs}
    ${ps_grpc_srcs}
    ${ps_grpc_hdrs}
    ${ps_proto_srcs}
    ${ps_proto_hdrs}
)

target_link_libraries(
    ps_grpc_proto
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
#   ps_grpc_proto
#   ${_REFLECTION}
#   ${_GRPC_GRPCPP}
#   ${_PROTOBUF_LIBPROTOBUF}
# )

# Targets route_guide_(client|server)
foreach(
    _target
    polynomialServiceClient
    # poloynomialServiceServer
)
    add_executable(
        ${_target}
        "${_target}.cc"
    )
    target_link_libraries(
        ${_target}
        ps_grpc_proto
        # helper
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF}
    )
endforeach()














