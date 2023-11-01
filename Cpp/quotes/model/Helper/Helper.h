#ifndef HELPER_H
#define HELPER_H

#include <google/protobuf/any.pb.h>
#include <aws/core/utils/json/JsonSerializer.h>

class Helper {
public:
    static google::protobuf::Any ConvertJsonViewToAny(const Aws::Utils::Json::JsonView& jsonView, const std::string& typeUrl);
};

#endif // HELPER_H