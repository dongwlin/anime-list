#include <iostream>
#include <string>
#include <vector>
#include <BaseValidate.hpp>
#include <typeinfo>

// BaseValidate Sample

int main()
{
	// 定义验证规则
	std::map<std::string, std::vector<std::string>> rule = {
		{"name", {"required", "min:3", "max:10"}},
		{"email", {"required", "email"}},
		{"age", {"required", "between:18,70"}}
	};

	// 定义错误信息
	std::map<std::string, std::string> message = {
		{"name.required", "NAME IS REQUIRED."},
		//{"age.between", "AGE MUST BE BETWEEN 18 AND 70"}
	};

	// 实例化验证器
	BaseValidate validator;
	validator.setRule(rule);
	validator.setMessage(message);

	// 要验证的数据
	std::map<std::string, std::string> data1 = {
		{"name", "ZhangSan"},
		{"email", "ZhangSan@example.com"},
		{"age", "45"}
	};

	// 要验证的数据
	std::map<std::string, std::string> data2 = {
		{"name", "Li"},
		{"email", "Li@example.com"},
		{"age", "17"}
	};

	// 要验证的数据
	std::map<std::string, std::string> data3 = {
		{"name", "WangWu"},
		{"email", "WangWuexample"},
		{"age", "72"}
	};

	// 错误信息
	std::map < std::string, std::vector<std::string>> errors1;

	// 错误信息
	std::map < std::string, std::vector<std::string>> errors2;

	// 错误信息
	std::map < std::string, std::vector<std::string>> errors3;

	// 进行验证
	bool success1 = validator.check(data1, errors1);

	// 进行验证
	bool success2 = validator.check(data2, errors2);

	// 进行验证
	bool success3 = validator.check(data3, errors3);

	std::cout << "data1: " << std::endl;
	if (success1)
	{
		std::cout << "validation passed!" << std::endl;
	}
	else
	{
		std::cout << "validation failed:" << std::endl;
		for (const auto& [field, messages] : errors1)
		{
			std::cout << field << ":" << std::endl;
			for (const auto& message : messages)
			{
				std::cout << " " << message << std::endl;
			}
		}
	}

	std::cout << "data2: " << std::endl;
	if (success2)
	{
		std::cout << "validation passed!" << std::endl;
	}
	else
	{
		std::cout << "validation failed:" << std::endl;
		for (const auto& [field, messages] : errors2)
		{
			std::cout << field << ":" << std::endl;
			for (const auto& message : messages)
			{
				std::cout << " " << message << std::endl;
			}
		}
	}

	std::cout << "data3: " << std::endl;
	if (success3)
	{
		std::cout << "validation passed!" << std::endl;
	}
	else
	{
		std::cout << "validation failed:" << std::endl;
		for (const auto& [field, messages] : errors3)
		{
			std::cout << field << ":" << std::endl;
			for (const auto& message : messages)
			{
				std::cout << " " << message << std::endl;
			}
		}
	}

	return 0;

}