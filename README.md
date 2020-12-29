# 📚 Wikitermia
Wikipedia client for terminal

## 🔨 Building
For building enter this lines:
```bash
cmake .
make
./wikitermia ban
```
### ⚠️ Recommendation
* If you use `MinGW` write:
```bash
cmake . -G "MinGW Makefiles" 
```
* If you use `Haiku x86` write:
```bash
cmake . -CMAKE_CXX_COMPILER=g++-x86
```

## 🔨 Dependencies
### 📄 [nlohmann_json](https://github.com/nlohmann/json)
JSON for Modern C++
### 📥 [libcurl](https://github.com/curl/curl)
libcurl is the library curl is using to do its job. It is readily available to be used by your software.
