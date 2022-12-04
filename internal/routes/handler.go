package routes

import (
	"encoding/json"
	"github.com/calebtracey/rugby-crawler-api/internal/facade"
	_ "github.com/calebtracey/rugby-crawler-api/internal/routes/statik"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Health check
	r.Handle("/health", h.HealthCheck()).Methods(http.MethodGet)

	r.Handle("/leaderboard", h.LeaderboardHandler()).Methods(http.MethodPost)
	r.Handle("/leaderboards", h.AllLeaderboardsHandler()).Methods(http.MethodGet)

	staticFs, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(staticFs)
	sh := http.StripPrefix("/swagger-ui/", staticServer)
	r.PathPrefix("/swagger-ui/").Handler(sh)

	return r
}

func (h *Handler) LeaderboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		apiRequest := leaderboard.RequestFromJSON(r.Body)
		apiResponse := h.Service.CrawlLeaderboardData(r.Context(), apiRequest)

		statusCode := apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.LeaderboardData))
		apiResponse.Message.AddMessageDetails(startTime)

		if err := apiResponse.ResponseToJSON(w); err != nil {
			log.Errorf("failed to marshal response: %s", err.RootCause)
			apiResponse.Message.ErrorLog = response.ErrorLogs{*err}
			statusCode = http.StatusInternalServerError
		}
		response.WriteHeader(w, statusCode)
	}
}

func (h *Handler) AllLeaderboardsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		apiResponse := h.Service.CrawlAllLeaderboardData(r.Context())

		statusCode := apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.LeaderboardData))
		apiResponse.Message.AddMessageDetails(startTime)

		if err := apiResponse.ResponseToJSON(w); err != nil {
			log.Errorf("failed to marshal response: %s", err.RootCause)
			apiResponse.Message.ErrorLog = response.ErrorLogs{*err}
			statusCode = http.StatusInternalServerError
		}
		response.WriteHeader(w, statusCode)
	}
}

func (h *Handler) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		if err != nil {
			log.Errorln(err.Error())
			return
		}
	}
}
