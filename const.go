package gonnachan

var (
	serversGelboru  = []string{ServerGelbooru, ServerSafebooru, ServerRule34, ServerTBIB}
	serversKonachan = []string{ServerKonachan, ServerYandere}
	serversSankaku  = []string{ServerSankaku}
)

//add order:favcount

const (
	//ServerKonachan Host
	ServerKonachan = "http://konachan.com/"
	//ServerYandere Host
	ServerYandere = "https://yande.re/"
	//ServerGelbooru Host
	ServerGelbooru = "https://gelbooru.com/"
	//ServerRule34 Host
	ServerRule34 = "https://rule34.xxx/"
	//ServerSafebooru Host
	ServerSafebooru = "https://safebooru.org/"
	//ServerTBIB Host
	ServerTBIB = "https://tbib.org/"
	//ServerSankaku Host
	ServerSankaku = "https://capi-beta.sankakucomplex.com/"
	//Common query to get post result in similar pages
	endpointGelboru  = "index.php?page=dapi&s=post&q=index&json=1&"
	endpointKonachan = "post.json?"
	endpointSankaku  = "post/index.json?page=1&"
	//Enum Alternative
	typeUnsupported = -1
	typeKonachan    = 0
	typeGelboru     = 1
	typeSankaku     = 2

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
