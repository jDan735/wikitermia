#include <nlohmann/json.hpp>
#include <curl/curl.h>
#include <iostream>
#include <sstream>

using namespace nlohmann;
using namespace std;

ostringstream stream;

size_t write_data(char *ptr, size_t size, size_t nmemb, void *userdata) {
    ostringstream *stream = (ostringstream*)userdata;
    size_t count = size * nmemb;
    stream->str("");
    stream->write(ptr, count);
    return count;
}

class Wikipedia {
    public:
        string download(string url){
            CURL *curl;
            CURLcode res;

            curl = curl_easy_init();
            if(curl) {
                curl_easy_setopt(curl, CURLOPT_CAINFO, "cacert.pem");
                curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_data);
                curl_easy_setopt(curl, CURLOPT_WRITEDATA, &stream);
                curl_easy_setopt(curl, CURLOPT_URL, url.c_str());

                curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);

                res = curl_easy_perform(curl);

                if(res != CURLE_OK)
                    cout << "curl_easy_perform() failed:" << curl_easy_strerror(res) << endl;

                curl_easy_cleanup(curl);
                return stream.str();
            }
            return "ban";
        }

        void setLang(string newLang) {
            lang = newLang;
        }

        json search(string query, string limit) {
            request = download("https://" + lang + ".wikipedia.org/w/api.php?action=query&prop=extracts&list=search&srsearch=" + query + "&srlimit=" + limit + "&srprop=size&format=json");

            auto data = json::parse(request)["query"]["search"];
            json arr = json::array();
            int l = data.size();

            for (int i = 0; i <= l; ++i) {
                arr += data[i]["title"];
            }

            return arr;
        }

        string getPage(string query) {
            request = download("https://" + lang + ".wikipedia.org/w/api.php?action=query&prop=extracts&explaintext&titles=" + query + "&exsentences=2&format=json");
            auto data = json::parse(request)["query"]["pages"];

            for (auto it = data.begin(); it != data.end(); ++it) {
                return it.value()["extract"];
            }

            return "ban";
        }

    private:
        string lang;
        string request;
};
