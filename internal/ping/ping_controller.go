package ping

import (
	"net/http"

	"github.com/SIC-Unud/sicgolib"
	"github.com/deedima3/yearbook-backend/internal/global"
	"github.com/gorilla/mux"
)

type PingController struct {
	router *mux.Router
	ps PingService
}

func (pc PingController) Ping(rw http.ResponseWriter, r *http.Request){
	res := pc.ps.Ping(r.Context())
	sicgolib.NewBaseResponse(200, res, nil, nil).SendResponse(&rw)
}

func (pc *PingController) InitializeController(){
	pc.router.HandleFunc(global.API_PING, pc.Ping).Methods(http.MethodGet)
}

func ProvidePingController(router *mux.Router, ps PingService) *PingController{
	return &PingController{router: router, ps: ps}
}