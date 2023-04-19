#pragma once

#include <nlohmann/json.hpp>

void divisionAnime(nlohmann::json animeList, nlohmann::json& animeListA, nlohmann::json& animeListB)
{
	bool next = true;
	for (auto [key, value] : animeList.items())
	{
		if (!value["status"])
		{
			continue;
		}
		if (next)
		{
			next = !next;
			animeListA.push_back(value);
		}
		else
		{
			next = !next;
			animeListB.push_back(value);
		}
	}
}