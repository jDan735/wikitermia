#include <nlohmann/json.hpp>
#include <iostream>
#include "main.h"

using namespace std;
using namespace nlohmann;

int main(int argc, char *argv[]){
    Wikipedia wiki;

    if (argc < 2) {
        cout << "Enter your wikipedia page request!" << endl;
        return 1;
    }

    wiki.setLang("en");
    //json res = wiki.search("Бан", "10");

    cout << wiki.getPage(argv[1]) << endl;

    return 0;
}
