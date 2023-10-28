from typing import Any
from flask import jsonify, make_response, status, Response

def generatePortMulti(
        key: str, 
        default_value: Any, 
        expression_value: Any
    ) -> Any:
    if key == "MULTI":
        return default_value
    return expression_value

def checkAPI(request, requestClass)->Response:
    if not isinstance(request.json, requestClass):
        return make_response(
            jsonify(
                {"error": "Invalid input"}
            ), 
            status.HTTP_400_BAD_REQUEST
        )
    return None