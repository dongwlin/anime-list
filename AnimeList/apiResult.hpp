#pragma once

#include <nlohmann/json.hpp>
#include <string>

std::string api_result(std::string msg = "fail", nlohmann::json data = nlohmann::json::array(), int code = 500)
{
	nlohmann::json result;
	result["code"] = code;
	result["msg"] = msg;
	result["data"] = data;
	return result.dump();
}

std::string api_success(nlohmann::json data = nlohmann::json::array(), std::string msg = "success", int code = 200)
{
	return api_result(msg, data, code);
}

std::string api_error(std::string msg = "fail", nlohmann::json data = nlohmann::json::array(),  int code = 500)
{
	return api_result(msg, data, code);
}

std::string api_error_clinet(std::string msg = "client error", nlohmann::json data = nlohmann::json::array(), int code = 400)
{
	return api_result(msg, data, code);
}

std::string api_error_404(std::string msg = "not found", nlohmann::json data = nlohmann::json::array(), int code = 404)
{
	return api_result(msg, data, code);
}