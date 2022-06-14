package controller

import (
	"net/http"
	"strconv"

	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/blogpost/dto"
	blogpostServicePkg "github.com/deedima3/yearbook-backend/internal/blogpost/service/api"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/gorilla/mux"
)

type BlogpostController struct {
	router *mux.Router
	bs     blogpostServicePkg.BlogpostService
}

func (bc *BlogpostController) viewUpvoteDownvote(rw http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	postIDVar := routerVar["postID"]
	postIDConv, _ := strconv.ParseUint(postIDVar, 10, 64)

	postVote, err := bc.bs.ViewUpvoteDownvote(r.Context(), postIDConv)
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("internal", "server error"),
		))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, postVote).ToJSON(rw)
}

func (bc *BlogpostController) deletePost(rw http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	postIDVar := routerVar["postID"]
	postIDConv, _ := strconv.ParseUint(postIDVar, 10, 64)

	deleteBlogPost := bc.bs.DeletePostByID(r.Context(), postIDConv)
	if deleteBlogPost != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			sicgolib.NewErrorResponseValue("server error", deleteBlogPost.Error()),
		))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, "Deleted").ToJSON(rw)
}

func (bc *BlogpostController) createPost(rw http.ResponseWriter, r *http.Request) {
	postRequest := new(dto.BlogPostRequestBody)

	if err := postRequest.FromJSON(r.Body); err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("request body", "invalid json format"),
		))
	}

	blogID, _ := bc.bs.CreatePost(r.Context(), *postRequest)
	sicgolib.NewBaseResponse(201, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, blogID).ToJSON(rw)
}

func (bc *BlogpostController) viewTopTwits(rw http.ResponseWriter, r *http.Request) {
	topTwits, err := bc.bs.ViewTopTwits(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("internal", "server error"),
		))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, topTwits).ToJSON(rw)
}

func (bc *BlogpostController) updateVotes(rw http.ResponseWriter, r *http.Request) {
	voteRequest := new(dto.BlogPostVotesRequestBody)

	if err := voteRequest.FromJSON(r.Body); err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("request body", "invalid json format"),
		))
	}
	msg, _ := bc.bs.UpdateVotes(r.Context(), *voteRequest)
	sicgolib.NewBaseResponse(201, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, msg).ToJSON(rw)
}

func (bc *BlogpostController) getTwitsPerPages(rw http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	pagesVar := routerVar["pages"]
	pagesConv, _ := strconv.ParseUint(pagesVar, 10, 64)

	twits, err := bc.bs.GetTwitsPerPages(r.Context(), pagesConv)
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			http.StatusBadRequest,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("internal", "server error"),
		))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, twits).ToJSON(rw)
}

func (bc *BlogpostController) InitializeController() {
	//Add your routes here
	bc.router.HandleFunc(global.API_INSERT_POST, bc.createPost).Methods(http.MethodPost, http.MethodOptions)
	bc.router.HandleFunc(global.API_DELETE_POST, bc.deletePost).Methods(http.MethodDelete, http.MethodOptions)
	bc.router.HandleFunc(global.API_VIEW_VOTES, bc.viewUpvoteDownvote).Methods(http.MethodGet, http.MethodOptions)
	bc.router.HandleFunc(global.API_VIEW_TOP_TWITS, bc.viewTopTwits).Methods(http.MethodGet, http.MethodOptions)
	bc.router.HandleFunc(global.API_UPDATE_VOTES, bc.updateVotes).Methods(http.MethodPatch, http.MethodOptions)
	bc.router.HandleFunc(global.API_VIEW_TWITS_PAGES, bc.getTwitsPerPages).Methods(http.MethodGet, http.MethodOptions)
}

func ProvideBlogpostController(router *mux.Router, bs blogpostServicePkg.BlogpostService) *BlogpostController {
	return &BlogpostController{router: router, bs: bs}
}
