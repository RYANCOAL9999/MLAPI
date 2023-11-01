cmake_minimum_required(VERSION 3.8)

project(quotes C CXX)

include(../cmake/common.cmake)

# Proto file
get_filename_component(ls_proto "../../../../../../../AIProto/lib/enum_expression.proto" ABSOLUTE)
get_filename_component(ls_proto "../../../../../../../AIProto/lib/np_library.proto" ABSOLUTE)
get_filename_component(svm_proto "../../../../../../../AIProto/svm_expression.proto" ABSOLUTE)
get_filename_component(svm_proto_path "${svm_proto}" PATH)

# Generated sources
set(enum_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.cc")
set(enum_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.h")
set(np_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.cc")
set(np_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.h")
set(svm_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/svm_expression.pb.cc")
set(svm_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/svm_expression.pb.h")
set(svm_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/svm_expression.grpc.pb.cc")
set(svm_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/svm_expression.grpc.pb.h")

add_custom_command(
    OUTPUT "${enum_proto_srcs}" "${enum_proto_hdrs}" "${np_proto_srcs}" "${np_proto_srcs}" "${svm_proto_srcs}" "${svm_proto_hdrs}" "${svm_grpc_srcs}" "${svm_grpc_hdrs}"
    COMMAND ${_PROTOBUF_PROTOC}
    ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
    --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
    -I "${svm_proto_path}"
    --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
    "${svm_proto}"
    DEPENDS "${svm_proto}"
)

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# svm_grpc_proto
add_library(
    svm_grpc_proto
    ${enum_proto_srcs}
    ${enum_proto_hdrs} 
    ${np_proto_srcs} 
    ${np_proto_srcs}
    ${svm_grpc_srcs}
    ${svm_grpc_hdrs}
    ${svm_proto_srcs}
    ${svm_proto_hdrs}
)

target_link_libraries(
    svm_grpc_proto
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
#   svm_grpc_proto
#   ${_REFLECTION}
#   ${_GRPC_GRPCPP}
#   ${_PROTOBUF_LIBPROTOBUF}
# )

# Targets route_guide_(client|server)
foreach(
    _target
    svmServiceClient
    # svmServiceServer
)
    add_executable(
        ${_target}
        "${_target}.cc"
    )
    target_link_libraries(
        ${_target}
        svm_grpc_proto
        # helper
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF}
    )
endforeach()














