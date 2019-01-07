package gonnachan

var (
	serversGelboru  = []string{ServerGelbooru, ServerSafebooru, ServerRule34}
	serversKonachan = []string{ServerKonachan, ServerYandere}
)

const (
	//ServerKonachan APIendPoint
	ServerKonachan = "http://konachan.com/"
	//ServerYandere APIendPoint
	ServerYandere = "https://yande.re/"
	//ServerGelbooru APIendPoint
	ServerGelbooru = "https://gelbooru.com/"
	//ServerSafebooru APIendPoint
	ServerSafebooru = "https://safebooru.org/"
	//ServerRule34 APIendPoint
	ServerRule34 = "https://rule34.xxx/"

	//Common query to get post result in similar pages
	endpointGelboru  = "index.php?page=dapi&s=post&q=index&json=1&"
	endpointKonachan = "post.json?"
	//Enum Alternative
	typeUnsupported = -1
	typeKonachan    = 0
	typeGelboru     = 1

	//Danbooru APIendPoint
	//Danbooru = "https://danbooru.donmai.us/posts/1.json"
	//Im too lazy to fix this at the moment.

	//ToDo
	//chan.sankakucomplex.com
	//idol.sankakucomplex.com

	//e-shuushuu.net

	//anime-pictures.net

	//danbooru.donmai.us

	//www.zerochan.net
)

const (
	//Rating

	//RatingSafe PG
	RatingSafe = "s"
	//RatingQuestionable +16
	RatingQuestionable = "q"
	//RatingExplicit +18
	RatingExplicit = "e"
)

const (
	//ErrNoResults No results found
	ErrNoResults = 1
	//ErrEmpty No answer from server
	ErrEmpty = 2
)
