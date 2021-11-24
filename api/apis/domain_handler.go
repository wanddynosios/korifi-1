package apis

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"code.cloudfoundry.org/cf-k8s-controllers/api/payloads"
	"code.cloudfoundry.org/cf-k8s-controllers/api/presenter"

	"github.com/go-http-utils/headers"
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"k8s.io/client-go/rest"

	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	DomainListEndpoint = "/v3/domains"
)

//counterfeiter:generate -o fake -fake-name CFDomainRepository . CFDomainRepository

type CFDomainRepository interface {
	FetchDomain(context.Context, client.Client, string) (repositories.DomainRecord, error)
	FetchDomainList(context.Context, client.Client, repositories.DomainListMessage) ([]repositories.DomainRecord, error)
}

type DomainHandler struct {
	logger      logr.Logger
	serverURL   url.URL
	domainRepo  CFDomainRepository
	buildClient ClientBuilder
	k8sConfig   *rest.Config // TODO: this would be global for all requests, not what we want
}

func NewDomainHandler(
	logger logr.Logger,
	serverURL url.URL,
	domainRepo CFDomainRepository,
	buildClient ClientBuilder,
	k8sConfig *rest.Config) *DomainHandler {
	return &DomainHandler{
		logger:      logger,
		serverURL:   serverURL,
		domainRepo:  domainRepo,
		buildClient: buildClient,
		k8sConfig:   k8sConfig,
	}
}

func (h *DomainHandler) DomainListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		h.logger.Error(err, "Unable to parse request query parameters")
		writeUnknownErrorResponse(w)
		return
	}

	domainListFilter := new(payloads.DomainList)
	err := schema.NewDecoder().Decode(domainListFilter, r.Form)
	if err != nil {
		switch err.(type) {
		case schema.MultiError:
			multiError := err.(schema.MultiError)
			for _, v := range multiError {
				_, ok := v.(schema.UnknownKeyError)
				if ok {
					h.logger.Info("Unknown key used in Domain filter")
					writeUnknownKeyError(w, domainListFilter.SupportedFilterKeys())
					return
				}
			}
			h.logger.Error(err, "Unable to decode request query parameters")
			writeUnknownErrorResponse(w)
			return

		default:
			h.logger.Error(err, "Unable to decode request query parameters")
			writeUnknownErrorResponse(w)
			return
		}
	}

	client, err := h.buildClient(h.k8sConfig, r.Header.Get(headers.Authorization))
	if err != nil {
		h.logger.Error(err, "Unable to create Kubernetes client")
		writeUnknownErrorResponse(w)
		return
	}

	domainList, err := h.domainRepo.FetchDomainList(ctx, client, domainListFilter.ToMessage())
	if err != nil {
		h.logger.Error(err, "Failed to fetch domain(s) from Kubernetes")
		writeUnknownErrorResponse(w)
		return
	}

	responseBody, err := json.Marshal(presenter.ForDomainList(domainList, h.serverURL))
	if err != nil {
		h.logger.Error(err, "Failed to render response")
		writeUnknownErrorResponse(w)
		return
	}

	_, _ = w.Write(responseBody)
}

func (h *DomainHandler) RegisterRoutes(router *mux.Router) {
	router.Path(DomainListEndpoint).Methods("GET").HandlerFunc(h.DomainListHandler)
}
