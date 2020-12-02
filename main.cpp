#include <iostream>
#include <curl/curl.h>

using namespace std;

int main(void){
    CURL *curl;
    CURLcode res;

    curl = curl_easy_init();
    if(curl) {
        curl_easy_setopt(curl, CURLOPT_CAINFO, "cacert.pem");
        curl_easy_setopt(curl, CURLOPT_URL, "http://ru.wikipedia.org/w/api.php?action=query&prop=extracts&list=search&srsearch=%D0%BA%D0%B0%D1%86&srlimit=10&srprop=size&format=json");

        /* example.com is redirected, so we tell libcurl to follow redirection */
        curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);

        /* Perform the request, res will get the return code */
        res = curl_easy_perform(curl);
        /* Check for errors */
        if(res != CURLE_OK)
            cout << "curl_easy_perform() failed:" << curl_easy_strerror(res) << endl;

        /* always cleanup */
        curl_easy_cleanup(curl);
    }

    return 0;
}
