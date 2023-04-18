#pragma once

#include <string>
#include <iostream>
#include <filesystem>
#include <fstream>
#include <nlohmann/json.hpp>

#include "httplib.h"
#include "utils.hpp"
#include "apiResult.hpp"
#include "jsonProcess.hpp"

// ¾²Ä¬ÔËÐÐ
#ifndef _DEBUG
#pragma comment(linker, "/SUBSYSTEM:WINDOWS")
#pragma comment(linker, "/ENTRY:mainCRTStartup")
#endif // !_DEBUG

const int port = 9317;

void init()
{
	std::filesystem::path res = "./resource";
	if (!std::filesystem::exists(res))
	{
		std::filesystem::create_directory(res);
	}
}

int main()
{
#ifdef _DEBUG
	std::cout << "START" << std::endl;
#endif // _DEBUG

	std::filesystem::path www = "./www";
	if (!std::filesystem::exists(www))
	{
#ifdef _DEBUG
		std::cout << "www dir not exists" << std::endl;
#endif // _DEBUG

		std::filesystem::create_directory(www);

#ifdef _DEBUG
		std::cout << "www dir created" << std::endl;
#endif // _DEBUG
	}
	else
	{
#ifdef _DEBUG
		std::cout << "www dir exists" << std::endl;
#endif // _DEBUG
	}

	httplib::Server server;

	server.set_mount_point("/", www.generic_string());

	server.Get("/hi",
		[](const httplib::Request& request, httplib::Response& response)
		{
			response.set_content(
				api_success(),
				"text/plain;charset=utf8"
			);
		}
	);

	server.Get("/open",
		[](const httplib::Request& request, httplib::Response& response)
		{
			if (request.has_param("folder"))
			{
				std::filesystem::path folder = utils::UTF8ToGBK(request.get_param_value("folder"));
				if (std::filesystem::exists(folder))
				{
					utils::open(folder.generic_string());

					response.set_content(
						api_success(),
						"text/plain;charset=utf8"
					);
#ifdef _DEBUG
					std::cout << "open folder: " << folder.generic_string() << std::endl;
#endif // _DEBUG
				}
				else
				{
#ifdef _DEBUG
					std::cout << "Error: folder not exist." << std::endl;
#endif // _DEBUG
					response.set_content(
						api_error_404(),
						"text/plain;charset=utf8"
					);
				}
				
			}
		}
	);

	server.Get("/animeList",
		[](const httplib::Request& request, httplib::Response& response)
		{
			std::fstream fin("./resource/anime.json");
			nlohmann::json anime = nlohmann::json::parse(fin);
			response.set_content(
				api_success(anime),
				"text/plian;charset=utf8"
			);
		}
	);

	server.Get("/animeListA",
		[](const httplib::Request& request, httplib::Response& response)
		{
			std::fstream fin("./resource/anime.json");
			nlohmann::json anime = nlohmann::json::parse(fin);
			nlohmann::json animeListA = nlohmann::json::array();
			nlohmann::json animeListB = nlohmann::json::array();
			divisionAnime(anime, animeListA, animeListB);
			response.set_content(
				api_success(animeListA),
				"text/plian;charset=utf8"
			);
		}
	);

	server.Get("/animeListB",
		[](const httplib::Request& request, httplib::Response& response)
		{
			std::fstream fin("./resource/anime.json");
			nlohmann::json anime = nlohmann::json::parse(fin);
			nlohmann::json animeListA = nlohmann::json::array();
			nlohmann::json animeListB = nlohmann::json::array();
			divisionAnime(anime, animeListA, animeListB);
			response.set_content(
				api_success(animeListB),
				"text/plian;charset=utf8"
			);
		}
	);

	server.Get("/stop",
		[&](const httplib::Request& request, httplib::Response& response)
		{
			server.stop();

#ifdef _DEBUG
			std::cout << "END" << std::endl;
#endif // _DEBUG

		}
	);

	utils::open("http://localhost:" + std::to_string(port));

	server.listen("localhost", port);
}