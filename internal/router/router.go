package router

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"grafana-dashboard/internal/cache"
	"grafana-dashboard/internal/domain/xerror"
	"grafana-dashboard/internal/dto"
	"grafana-dashboard/internal/middleware"
	"io/ioutil"
	"net/http"
)

func Body(req *http.Request, body any) error {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		return err
	}
	return nil
}

func SendError(w http.ResponseWriter, err error) {
	var e xerror.Error
	var res dto.Error
	if errors.As(err, &e) {
		res = dto.Error{
			Message: e.Message,
			Code:    e.Code,
		}
		w.WriteHeader(e.Code)
	}

	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(`{"code": 500,"message": "cause marshal error"}`))
	}

	_, _ = w.Write(b)
}

func Send(w http.ResponseWriter, code int, data any) {
	resByte, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(`{"code": 500,"message": "send marshal error"}`))
	}
	w.WriteHeader(code)
	_, _ = w.Write(resByte)
}

func New(cache cache.Cache[struct{}]) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.Metrics)

	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		})

		r.Post("/link", func(w http.ResponseWriter, r *http.Request) {
			var body dto.LinkDTO
			err := Body(r, &body)
			if err != nil {
				SendError(w, err)
				return
			}

			_, has := cache.Get(body.Id)
			if has {
				SendError(w, xerror.New(xerror.ErrConflict, "link id exist"))
				return
			}

			cache.Add(body.Id, struct{}{})

			Send(w, 200, dto.ResLink{
				Id:      body.Id,
				Facture: "eer0",
			})
		})

		r.Post("/unlink", func(w http.ResponseWriter, r *http.Request) {
			var body dto.LinkDTO
			err := Body(r, &body)
			if err != nil {
				SendError(w, err)
				return
			}

			_, has := cache.Get(body.Id)
			if !has {
				SendError(w, xerror.New(xerror.ErrNotFound, "link id not found"))
				return
			}

			cache.Delete(body.Id)

			Send(w, 200, dto.ResLink{
				Id:      body.Id,
				Facture: "eer0",
			})
		})
	})

	return r
}
