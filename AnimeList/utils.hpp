#pragma once

#include <Windows.h>
#include <string>
#include <xstring>
#include <utility>
#include <sstream>
#include <iomanip>
#include <chrono>

namespace utils
{

    std::string UTF8ToGBK(const std::string& strUTF8)
    {
        int len = MultiByteToWideChar(CP_UTF8, 0, strUTF8.c_str(), -1, NULL, 0);
        TCHAR* wszGBK = new TCHAR[len + 1];
        memset(wszGBK, 0, len * 2 + 2);
        MultiByteToWideChar(CP_UTF8, 0, strUTF8.c_str(), -1, wszGBK, len);

        len = WideCharToMultiByte(CP_ACP, 0, wszGBK, -1, NULL, 0, NULL, NULL);
        char* szGBK = new char[len + 1];
        memset(szGBK, 0, len + 1);
        WideCharToMultiByte(CP_ACP, 0, wszGBK, -1, szGBK, len, NULL, NULL);
        std::string strTemp(szGBK);
        delete[]szGBK;
        delete[]wszGBK;
        return strTemp;
    }

    std::string GBKToUTF8(const std::string& strGBK)
    {
        int len = MultiByteToWideChar(CP_ACP, 0, strGBK.c_str(), -1, NULL, 0);
        TCHAR* wszUTF8 = new TCHAR[len + 1];
        memset(wszUTF8, 0, len * 2 + 2);
        MultiByteToWideChar(CP_ACP, 0, strGBK.c_str(), -1, wszUTF8, len);

        len = WideCharToMultiByte(CP_UTF8, 0, wszUTF8, -1, NULL, 0, NULL, NULL);
        char* szUTF8 = new char[len + 1];
        memset(szUTF8, 0, len + 1);
        WideCharToMultiByte(CP_UTF8, 0, wszUTF8, -1, szUTF8, len, NULL, NULL);
        std::string strTemp(szUTF8);
        delete[]szUTF8;
        delete[]wszUTF8;
        return strTemp;
    }

    std::wstring stringToWString_ansi(const std::string& str)
    {
        size_t strLen = str.size();
        int num = MultiByteToWideChar(CP_ACP, 0, str.c_str(), -1, nullptr, 0);
        wchar_t* wide = new wchar_t[num];
        MultiByteToWideChar(CP_ACP, 0, str.c_str(), -1, wide, num);
        std::wstring w_str(wide);
        delete[] wide;
        return w_str;
    }

    std::wstring stringToWString_utf8(const std::string& str)
    {
        size_t strLen = str.size();
        int num = MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, nullptr, 0);
        wchar_t* wide = new wchar_t[num];
        MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, wide, num);
        std::wstring w_str(wide);
        delete[] wide;
        return w_str;
    }

    enum CodeType 
    {
        ANSI = 0,
        UTF8
    };

    void open(std::string folderName, CodeType codeType = ANSI)
    {
        std::wstring tmp;
        switch (codeType)
        {
        case ANSI:
            tmp = stringToWString_ansi(folderName);
            break;
        case UTF8:
            tmp = stringToWString_utf8(folderName);
            break;
        default:
            break;
        }

        if (tmp.empty())
        {
            std::cout << "Error: Code Type not exist." << std::endl;
            exit(1);
        }

        ShellExecute(nullptr, L"open", tmp.c_str(), nullptr, nullptr, SW_SHOWNORMAL);
    }

    std::pair<std::string, std::string> argExtraction(const std::string& argStr)
    {
        std::string tmp = argStr.substr(12, argStr.size());
        size_t opEnd = tmp.find_first_of('/');
        if (opEnd != std::string::npos)
        {
            std::string op = tmp.substr(0, tmp.find_first_of('/'));
            tmp = tmp.substr(opEnd + 1, tmp.size());
            return std::pair<std::string, std::string>(op, tmp);
        }
        return std::pair<std::string, std::string>("", "");
    }

    std::string url_decode(const std::string& str) {
        std::stringstream ss;
        ss << std::hex << std::setfill('0');
        for (auto i = str.begin(), end = str.end(); i != end; ++i) {
            if (*i == '+') {
                ss << ' ';
            }
            else if (*i == '%' && std::distance(i, end) > 2 && std::isxdigit(*(i + 1)) && std::isxdigit(*(i + 2))) {
                ss << static_cast<char>(std::stoi(std::string(i + 1, i + 3), nullptr, 16));
                i += 2;
            }
            else {
                ss << *i;
            }
        }
        return ss.str();
    }

    bool to_bool(std::string str)
    {
        if (str == "true")
        {
            return true;
        }
        else if (str == "false")
        {
            return false;
        }

        if (str == "1")
        {
            return true;
        }
        else if (str == "0")
        {
            return false;
        }

        return false;
    }

    std::string getTimestamp()
    {
        auto ms = std::chrono::steady_clock::now().time_since_epoch();
        return std::to_string(ms.count());
    }
}
