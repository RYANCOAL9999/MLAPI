cmake_minimum_required(VERSION 3.8)

project(quotes C CXX)

include(../cmake/common.cmake)

# Proto file
get_filename_component(ls_proto "../../../../../../../AIProto/lib/enum_expression.proto" ABSOLUTE)
get_filename_component(ls_proto "../../../../../../../AIProto/lib/np_library.proto" ABSOLUTE)
get_filename_component(nbs_proto "../../../../../../../AIProto/nearest_neighbors.proto" ABSOLUTE)
get_filename_component(nbs_proto_path "${nbs_proto}" PATH)

# Generated sources
set(enum_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.cc")
set(enum_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.h")
set(np_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.cc")
set(np_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.h")
set(nbs_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/nearest_neighbors.pb.cc")
set(nbs_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/nearest_neighbors.pb.h")
set(nbs_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/nearest_neighbors.grpc.pb.cc")
set(nbs_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/nearest_neighbors.grpc.pb.h")

add_custom_command(
    OUTPUT "${enum_proto_srcs}" "${enum_proto_hdrs}" "${np_proto_srcs}" "${np_proto_srcs}" "${nbs_proto_srcs}" "${nbs_proto_hdrs}" "${nbs_grpc_srcs}" "${nbs_grpc_hdrs}"
    COMMAND ${_PROTOBUF_PROTOC}
    ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
    --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
    -I "${nbs_proto_path}"
    --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
    "${nbs_proto}"
    DEPENDS "${nbs_proto}"
)

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# nbs_grpc_proto
add_library(
    nbs_grpc_proto
    ${enum_proto_srcs}
    ${enum_proto_hdrs} 
    ${np_proto_srcs} 
    ${np_proto_srcs}
    ${nbs_grpc_srcs}
    ${nbs_grpc_hdrs}
    ${nbs_proto_srcs}
    ${nbs_proto_hdrs}
)

target_link_libraries(
    nbs_grpc_proto
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
#   nbs_grpc_proto
#   ${_REFLECTION}
#   ${_GRPC_GRPCPP}
#   ${_PROTOBUF_LIBPROTOBUF}
# )

# Targets route_guide_(client|server)
foreach(
    _target
    neighborsServiceClient
    # neighborsServiceServer
)
    add_executable(
        ${_target}
        "${_target}.cc"
    )
    target_link_libraries(
        ${_target}
        nbs_grpc_proto
        # helper
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF}
    )
endforeach()














