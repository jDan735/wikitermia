package wikitermia

var LANG string = "en"
var HOST string = "https://" + LANG + ".wikipedia.org/w/api.php?"

func SetLang(lang string){
	LANG = lang
    HOST = "https://" + LANG + ".wikipedia.org/w/api.php?"
}
