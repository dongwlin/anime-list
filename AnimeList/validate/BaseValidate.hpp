#pragma once
#include <map>
#include <vector>
#include <string>
#include <regex>
#include <algorithm>
#include <stdexcept>

/**
 * 验证器(针对字符串的)
 */
class BaseValidate {
protected:
    /**
     * 验证规则
     */
    std::map<std::string, std::vector<std::string>> rule;
    /**
     * 错误信息
     */
    std::map<std::string, std::string> message;
    /**
     * 默认错误信息
     */
    std::map<std::string, std::string> defaultMessage = {
        {"required", "The :attribute field is required."},
        {"min", "The :attribute field must be at least :value."},
        {"max", "The :attribute field may not be greater than :value."},
        {"email", "The :attribute field must be a valid email address."},
        {"url", "The :attribute field must be a valid URL."},
        {"regex", "The :attribute field format is invalid."},
        {"between", "The :attribute field must be between :value1 and :value2."},
        {"invalid_rule", "The :rule validation rule for :attribute filed is not supported."},
        {"invalid_type", "The :attribute filed must be a valid data type."},
    };

    // 获取错误信息
    std::string getMessage(const std::string& key, const std::string& attribute, const std::string& value1 = "", const std::string& value2 = "")
    {
        // 优先使用用户自定义的错误信息
        auto iter = message.find(attribute + "." + key);
        if (iter != message.end())
        {
            return iter->second;
        }
        // 默认错误信息
        std::string msg = "";
        if (key == "required" || key == "email" || key == "url" || key == "regex" || key == "invalid_type")
        {
            msg = defaultMessage[key];
            return replace(msg, ":attribute", attribute);
        }
        else if (key == "min" || key == "max")
        {
            msg = defaultMessage[key];
            msg = replace(msg, ":attribute", attribute);
            msg = replace(msg, ":value", value1);
            return msg;
        }
        else if (key == "between")
        {
            msg = defaultMessage[key];
            msg = replace(msg, ":attribute", attribute);
            msg = replace(msg, ":value1", value1);
            msg = replace(msg, ":value2", value2);
            return msg;
        }
        else if (key == "invalid_rule")
        {
            msg = defaultMessage["invalid_rule"];
            msg = replace(msg, ":rule", value1);
            msg = replace(msg, ":attribute", attribute);
            return msg;
        }
        msg = defaultMessage["invalid_rule"];
        msg = replace(msg, ":rule", key);
        msg = replace(msg , ":attribute", attribute);
        return msg;
    }

    // 获取错误信息
    std::string getMessage(const std::string& key, const std::string& attribute, const int& value1, const int& value2 = 0)
    {
        std::string strValue1 = std::to_string(value1);
        std::string strValue2 = std::to_string(value2);
        return getMessage(key, attribute, strValue1, strValue2);
    }

    // 字符串替换
    std::string replace(std::string str, const std::string& from, const std::string& to)
    {
        size_t pos = str.find(from);
        if (pos != std::string::npos)
        {
            str.replace(pos, from.length(), to);
        }
        return str;
    }

public:
    BaseValidate() = default;

    BaseValidate(const std::map<std::string, std::vector<std::string>>& rule, const std::map<std::string, std::string>& message) : rule(rule), message(message) {}

    virtual ~BaseValidate() = default;

    // 设置验证规则
    void setRule(const std::map<std::string, std::vector<std::string>>& rule)
    {
        this->rule = rule;
    }

    // 设置错误信息
    void setMessage(const std::map<std::string, std::string>& message)
    {
        this->message = message;
    }

    // 验证
    bool check(const std::map<std::string, std::string>& input, std::map<std::string, std::vector<std::string>>& errors)
    {
        bool success = true;
        for (const auto& [field, rules] : rule) {
            auto iter = input.find(field);
            if (iter == input.end()) {
                // required约束的字段不存在时
                if (std::find(rules.begin(), rules.end(), "required") != rules.end())
                {
                    success = false;
                    errors[field].push_back(getMessage("required", field));
                }
                continue;
            }
            auto& input_value = iter->second;
            for (const auto& rule : rules) {
                if (rule == "required") {
                    // 必填项
                    if (input_value.empty()) {
                        success = false;
                        errors[field].push_back(getMessage("required", field));
                    }
                }
                else if (rule.substr(0, 6) == "regex:") {
                    // 正则表达式验证
                    try {
                        std::regex pattern(rule.substr(6));
                        if (!std::regex_match(input_value, pattern)) {
                            success = false;
                            errors[field].push_back(getMessage("regex", field));
                        }
                    }
                    catch (const std::regex_error& e) {
                        // 语法错误或正则表达式太复杂时
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, rule));
                    }
                }
                else if (rule.substr(0, 4) == "min:") {
                    // 最小长度验证
                    try {
                        int min_length = std::stoi(rule.substr(4));
                        if (input_value.length() < min_length) {
                            success = false;
                            errors[field].push_back(getMessage("min", field, min_length));
                        }
                    }
                    catch (const std::invalid_argument& e) {
                        // 字符串不能转化为整数时
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, "mix"));
                    }
                    catch (const std::out_of_range& e) {
                        // 字符串转换后超出int范围
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, "mix"));
                    }      
                }
                else if (rule.substr(0, 4) == "max:") {
                    // 最大长度验证
                    try {
                        int max_length = std::stoi(rule.substr(4));
                        if (input_value.length() > max_length) {
                            success = false;
                            errors[field].push_back(getMessage("max", field, max_length));
                        }
                    }
                    catch (const std::invalid_argument& e) {
                        // 字符串不能转化为整数时
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, "max"));
                    }
                    catch (const std::out_of_range& e) {
                        // 字符串转换后超出int范围
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, "max"));
                    }
                }
                else if (rule == "email") {
                    // 邮箱格式验证
                    try {
                        std::regex pattern(R"(\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b)");
                        if (!std::regex_match(input_value, pattern)) {
                            success = false;
                            errors[field].push_back(getMessage("email", field));
                        }
                    }
                    catch (const std::regex_error& e) {
                        // 语法错误或正则表达式太复杂时
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, rule));
                    }
                }
                else if (rule == "url") {
                    // URL 格式验证
                    
                    try {
                        std::regex pattern(R"(https?://\S+)");
                        if (!std::regex_match(input_value, pattern)) {
                            success = false;
                            errors[field].push_back(getMessage("url", field));
                        }
                    }
                    catch (const std::regex_error& e) {
                        // 语法错误或正则表达式太复杂时
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, rule));
                    }
                }
                else if (rule.substr(0, 8) == "between:") {
                    // 整数范围验证
                    size_t colon_pos = rule.find(':');
                    size_t comma_pos = rule.find(',');
                    // :和,不存在时
                    if (colon_pos == std::string::npos || comma_pos == std::string::npos)
                    {
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, "between"));
                    }
                    try {
                        int min_value = std::stoi(rule.substr(colon_pos + 1, comma_pos));
                        int max_value = std::stoi(rule.substr(comma_pos + 1));
                        if (std::stoi(input_value) < min_value || std::stoi(input_value) > max_value) {
                            success = false;
                            errors[field].push_back(getMessage("between", field, min_value, max_value));
                        }
                    }
                    catch (const std::invalid_argument& e) {
                        // 字符串不能转化为整数时
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, "between"));
                    }
                    catch (const std::out_of_range& e) {
                        // 字符串转换后超出int范围
                        success = false;
                        errors[field].push_back(getMessage("invalid_rule", field, "between"));
                    }
                }
            }
        }
        return success;
    }
};