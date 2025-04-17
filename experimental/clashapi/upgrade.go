package clashapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func upgradeRouter(server *Server) http.Handler {
	r := chi.NewRouter()
	r.Post("/ui", updateUI(server))
	return r
}

func updateUI(server *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if server.externalUI != "" {
			err := server.downloadExternalUI()
			if err != nil {
				server.logger.Error("download external ui error: ", err)
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, newError(err.Error()))
				return
			}
			render.JSON(w, r, render.M{"status": "ok"})
		}
		render.NoContent(w, r)
	}
}
