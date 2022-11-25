package routes

import (
	"encoding/json"
	request2 "github.com/calebtracey/rugby-crawler-api/external/models/request"
	"github.com/calebtracey/rugby-crawler-api/external/models/response"
	"github.com/calebtracey/rugby-crawler-api/internal/facade"
	_ "github.com/calebtracey/rugby-crawler-api/internal/routes/statik"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Handler struct {
	Service facade.APIFacadeI
}

func (h *Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Health check
	r.Handle("/health", h.HealthCheck()).Methods(http.MethodGet)

	r.Handle("/competition", h.CompetitionHandler()).Methods(http.MethodPost)

	staticFs, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(staticFs)
	sh := http.StripPrefix("/swagger-ui/", staticServer)
	r.PathPrefix("/swagger-ui/").Handler(sh)

	return r
}

func (h *Handler) CompetitionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		var compResponse response.CrawlLeaderboardResponse
		var compRequest request2.CrawlLeaderboardRequest
		defer func() {
			status, _ := strconv.Atoi(compResponse.Message.Status)
			hn, _ := os.Hostname()
			compResponse.Message.HostName = hn
			compResponse.Message.TimeTaken = time.Since(startTime).String()
			_ = json.NewEncoder(writeHeader(w, status)).Encode(compResponse)
		}()
		body, bodyErr := readBody(r.Body)

		if bodyErr != nil {
			compResponse.Message.ErrorLog = errorLogs([]error{bodyErr}, "Unable to read psqlRequest body", http.StatusBadRequest)
			return
		}
		err := json.Unmarshal(body, &compRequest)
		if err != nil {
			compResponse.Message.ErrorLog = errorLogs([]error{err}, "Unable to parse psqlRequest", http.StatusBadRequest)
			return
		}

		compResponse = h.Service.CompetitionCrawlData(r.Context(), compRequest)
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
