#pragma once

#include <fstream>
#include <nlohmann/json.hpp>
#include <string>

void divisionAnime(nlohmann::json animeList, nlohmann::json& animeListA, nlohmann::json& animeListB)
{
	for (auto [key, value] : animeList.items())
	{
		if (atoi(key.c_str()) % 2)
		{
			animeListB.push_back(value);
		}
		else
		{
			animeListA.push_back(value);
		}
	}
}