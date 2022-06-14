package global

const (
	API_PING              = "/ping"
	API_INSERT_USER       = "/register"
	API_UPDATE_USER       = "/edit"
	API_ALL_USER          = "/all-user"
	API_LOGIN             = "/login"
	API_REFRESH_TOKEN     = "/refresh"
	API_UPDATE_VOTES      = "/vote"
	API_INSERT_POST       = "/twits"
	API_VIEW_TOP_TWITS    = "/twits"
	API_DELETE_POST       = "/twits/{postID}"
	API_VIEW_VOTES        = "/twits/{postID}"
	API_VIEW_TWITS_PAGES  = "/twits/pages/{pages}"
	API_GET_ALL_PAGES     = "/pages"
	API_NEW_BLOGPAGE      = "/pages"
	API_SEARCH_USER_PAGES = "/pages/search"
	API_UPDATE_BLOGPAGE   = "/user/pages"
	API_GET_USER_PAGES    = "/user/pages/{userID}"
	API_BIRTHDAY_WEEK     = "/birthday"
	API_IS_BIRTHDAY       = "/birthday/{owner}"
)
