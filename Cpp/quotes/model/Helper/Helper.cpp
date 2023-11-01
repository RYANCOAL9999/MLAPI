#include "Helper.h"

google::protobuf::Any Helper::ConvertJsonViewToAny(const Aws::Utils::Json::JsonView &jsonView, const std::string &typeUrl)
{
    google::protobuf::Any any;
    any.set_type_url(typeUrl);

    // Serialize the JSON data to a string using the provided jsonView.
    std::string jsonBytes = jsonView.WriteReadable();
    any.set_value(jsonBytes);

    return any;
}
