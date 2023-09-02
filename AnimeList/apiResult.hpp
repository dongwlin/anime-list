#pragma once

#include <nlohmann/json.hpp>
#include <string>

using json = nlohmann::json;

std::string api_result(std::string msg = "fail", json data = json::array(), int code = 500)
{
	json result;
	result["code"] = code;
	result["msg"] = msg;
	result["data"] = data;
	return result.dump();
}

std::string api_success(json data = json::array(), std::string msg = "success", int code = 200)
{
	return api_result(msg, data, code);
}

std::string api_error(std::string msg = "fail", json data = json::array(),  int code = 500)
{
	return api_result(msg, data, code);
}

std::string api_bad_request(std::string msg = "bad request", json data = json::array(), int code = 400)
{
	return api_result(msg, data, code);
}

std::string api_error_404(std::string msg = "not found", json data = json::array(), int code = 404)
{
	return api_result(msg, data, code);
}