#pragma once

#include <string>
#include <iostream>
#include <filesystem>
#include <fstream>
#include <nlohmann/json.hpp>
#include <httplib/httplib.h>
#include <Windows.h>
#include <chrono>
#include <cstdlib>

#include "utils.hpp"
#include "apiResult.hpp"
#include "jsonProcess.hpp"

// ¾²Ä¬ÔËÐÐ
#ifndef _DEBUG
#pragma comment(linker, "/SUBSYSTEM:WINDOWS")
#pragma comment(linker, "/ENTRY:mainCRTStartup")
#endif // !_DEBUG

#ifdef _DEBUG
const int PORT = 9311;
#else
const int PORT = 9317;
#endif // _DEBUG
std::string ContentType = "application/json;charset=utf8";

void init()
{
	std::filesystem::path data = "./data";
	if (!std::filesystem::exists(data))
	{
#ifdef _DEBUG
		std::cout << "resource dir not exist" << std::endl;
#endif // _DEBUG

		std::filesystem::create_directory(data);

#ifdef _DEBUG
		std::cout << "resource dir created" << std::endl;
#endif // _DEBUG

		nlohmann::json animeObj;
		animeObj["id"] = utils::getTimestamp();
		animeObj["name"] = "example";
		animeObj["status"] = true;
		animeObj["type"] = -1;
		animeObj["dir"] = "";
		animeObj["url"] = "";
		animeObj["img"] = "none";
		animeObj["day"] = -1;
		nlohmann::json anime = nlohmann::json::array();
		anime.push_back(animeObj);

		std::filesystem::path animeFile = data / "anime.json";
		std::ofstream out(animeFile);
		out << std::setw(4) << anime << std::endl;
		out.close();
		
#ifdef _DEBUG
		std::cout << "anime.json created" << std::endl;
#endif // _DEBUG

	}
}

int main()
{
#ifdef _DEBUG
	std::cout << "START" << std::endl;
#endif // _DEBUG


#ifdef _DEBUG
	HANDLE hMutex = CreateMutex(nullptr, TRUE, L"AnimeListDebug");
#else
	HANDLE hMutex = CreateMutex(nullptr, TRUE, L"AnimeListRelease");
#endif // _DEBUG
	
	if (GetLastError() == ERROR_ALREADY_EXISTS)
	{
#ifdef _DEBUG
		std::cout << "AnimeList is already running" << std::endl;
#endif // _DEBUG

		utils::open("http://localhost:" + std::to_string(PORT));

		return 0;
	}


	init();

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
				ContentType
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
						ContentType
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
						ContentType
					);
				}
			}
		}
	);

	server.Get("/animeList",
		[](const httplib::Request& request, httplib::Response& response)
		{
			std::fstream fin("./data/anime.json");
			nlohmann::json anime = nlohmann::json::parse(fin);
			response.set_content(
				api_success(anime),
				ContentType
			);
			fin.close();
		}
	);

	server.Get("/animeListA",
		[](const httplib::Request& request, httplib::Response& response)
		{
			std::fstream fin("./data/anime.json");
			nlohmann::json anime = nlohmann::json::parse(fin);
			nlohmann::json animeListA = nlohmann::json::array();
			nlohmann::json animeListB = nlohmann::json::array();
			divisionAnime(anime, animeListA, animeListB);
			response.set_content(
				api_success(animeListA),
				ContentType
			);
			fin.close();
		}
	);

	server.Get("/animeListB",
		[](const httplib::Request& request, httplib::Response& response)
		{
			std::fstream fin("./data/anime.json");
			nlohmann::json anime = nlohmann::json::parse(fin);
			nlohmann::json animeListA = nlohmann::json::array();
			nlohmann::json animeListB = nlohmann::json::array();
			divisionAnime(anime, animeListA, animeListB);
			response.set_content(
				api_success(animeListB),
				ContentType
			);
			fin.close();
		}
	);

	server.Post("/add",
		[](const httplib::Request& request, httplib::Response& response)
		{
			std::ifstream fin("./data/anime.json");
			nlohmann::json animeList = nlohmann::json::parse(fin);
			fin.close();

			nlohmann::json animeObj;

			animeObj["id"] = utils::getTimestamp();
			if (request.has_file("name"))
			{
				animeObj["name"] = request.get_file_value("name").content;
			}
			else
			{
				animeObj["name"] = "example";
			}
			if (request.has_file("status") && request.get_file_value("status").content == "false")
			{
				animeObj["status"] = utils::to_bool(request.get_file_value("status").content);
			}
			else
			{
				animeObj["status"] = true;
			}
			if (request.has_file("type") && request.get_file_value("type").content == "0" || request.get_file_value("type").content == "1")
			{
				animeObj["type"] = atoi(request.get_file_value("type").content.c_str());
			}
			else
			{
				animeObj["type"] = -1;
			}
			if (request.has_file("dir"))
			{
				animeObj["dir"] = request.get_file_value("dir").content;
			}
			else
			{
				animeObj["dir"] = "";
			}
			if (request.has_file("url"))
			{
				animeObj["url"] = request.get_file_value("url").content;
			}
			else
			{
				animeObj["url"] = "";
			}
			if (request.has_file("img"))
			{
				animeObj["img"] = request.get_file_value("img").content;
			}
			else
			{
				animeObj["img"] = "none";
			}
			if (request.has_file("day") && 
				request.get_file_value("day").content == "0" ||
				request.get_file_value("day").content == "1" ||
				request.get_file_value("day").content == "2" ||
				request.get_file_value("day").content == "3" ||
				request.get_file_value("day").content == "4" ||
				request.get_file_value("day").content == "5" ||
				request.get_file_value("day").content == "6"
				)
			{
				animeObj["day"] = atoi(request.get_file_value("day").content.c_str());
			}
			else
			{
				animeObj["day"] = -1;

			}

			animeList.push_back(animeObj);

			std::ofstream fout("./data/anime.json");
			fout << std::setw(4) << animeList << std::endl;
			fout.close();

#ifdef _DEBUG
			std::cout << "add success" << std::endl;
#endif // _DEBUG

			response.set_content(
				api_success(),
				ContentType
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

	utils::open("http://localhost:" + std::to_string(PORT));

	server.listen("localhost", PORT);

	ReleaseMutex(hMutex);
	CloseHandle(hMutex);
}