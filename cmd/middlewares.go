package main

import (
	"fmt"
	"github.com/ihamzapped/rss-aggregator/internal/auth"
	"github.com/ihamzapped/rss-aggregator/internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (api *ApiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondErr(w, http.StatusUnauthorized, fmt.Sprintf("%v", err))
			return
		}

		user, err := api.DB.GetUserByApiKey(r.Context(), key)

		if err != nil {
			respondErr(w, http.StatusNotFound, fmt.Sprintf("%v", err))
			return
		}

		handler(w, r, user)

	}
}
