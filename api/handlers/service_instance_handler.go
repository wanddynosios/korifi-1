package handlers

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"code.cloudfoundry.org/korifi/api/apierrors"
	"code.cloudfoundry.org/korifi/api/payloads"
	"code.cloudfoundry.org/korifi/api/routing"

	"code.cloudfoundry.org/korifi/api/presenter"

	"code.cloudfoundry.org/korifi/api/repositories"

	"code.cloudfoundry.org/korifi/api/authorization"

	"github.com/go-chi/chi"

	"github.com/go-logr/logr"
)

const (
	ServiceInstancesPath = "/v3/service_instances"
	ServiceInstancePath  = "/v3/service_instances/{guid}"
)

//counterfeiter:generate -o fake -fake-name CFServiceInstanceRepository . CFServiceInstanceRepository
type CFServiceInstanceRepository interface {
	CreateServiceInstance(context.Context, authorization.Info, repositories.CreateServiceInstanceMessage) (repositories.ServiceInstanceRecord, error)
	ListServiceInstances(context.Context, authorization.Info, repositories.ListServiceInstanceMessage) ([]repositories.ServiceInstanceRecord, error)
	GetServiceInstance(context.Context, authorization.Info, string) (repositories.ServiceInstanceRecord, error)
	DeleteServiceInstance(context.Context, authorization.Info, repositories.DeleteServiceInstanceMessage) error
}

type ServiceInstanceHandler struct {
	serverURL           url.URL
	serviceInstanceRepo CFServiceInstanceRepository
	spaceRepo           SpaceRepository
	decoderValidator    *DecoderValidator
}

func NewServiceInstanceHandler(
	serverURL url.URL,
	serviceInstanceRepo CFServiceInstanceRepository,
	spaceRepo SpaceRepository,
	decoderValidator *DecoderValidator,
) *ServiceInstanceHandler {
	return &ServiceInstanceHandler{
		serverURL:           serverURL,
		serviceInstanceRepo: serviceInstanceRepo,
		spaceRepo:           spaceRepo,
		decoderValidator:    decoderValidator,
	}
}

//nolint:dupl
func (h *ServiceInstanceHandler) serviceInstanceCreateHandler(r *http.Request) (*routing.Response, error) {
	authInfo, _ := authorization.InfoFromContext(r.Context())
	logger := logr.FromContextOrDiscard(r.Context()).WithName("service-instance-handler.service-instance-create")

	var payload payloads.ServiceInstanceCreate
	if err := h.decoderValidator.DecodeAndValidateJSONPayload(r, &payload); err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "failed to decode payload")
	}

	spaceGUID := payload.Relationships.Space.Data.GUID
	_, err := h.spaceRepo.GetSpace(r.Context(), authInfo, spaceGUID)
	if err != nil {
		return nil, apierrors.LogAndReturn(
			logger,
			apierrors.AsUnprocessableEntity(err, "Invalid space. Ensure that the space exists and you have access to it.", apierrors.NotFoundError{}, apierrors.ForbiddenError{}),
			"Failed to fetch namespace from Kubernetes",
			"spaceGUID", spaceGUID,
		)
	}

	serviceInstanceRecord, err := h.serviceInstanceRepo.CreateServiceInstance(r.Context(), authInfo, payload.ToServiceInstanceCreateMessage())
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "Failed to create service instance", "Service Instance Name", serviceInstanceRecord.Name)
	}

	return routing.NewHandlerResponse(http.StatusCreated).WithBody(presenter.ForServiceInstance(serviceInstanceRecord, h.serverURL)), nil
}

func (h *ServiceInstanceHandler) serviceInstanceListHandler(r *http.Request) (*routing.Response, error) {
	authInfo, _ := authorization.InfoFromContext(r.Context())
	logger := logr.FromContextOrDiscard(r.Context()).WithName("service-instance-handler.service-instance-list")

	if err := r.ParseForm(); err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "Unable to parse request query parameters")
	}

	for k := range r.Form {
		if strings.HasPrefix(k, "fields[") || k == "per_page" {
			r.Form.Del(k)
		}
	}

	listFilter := new(payloads.ServiceInstanceList)
	err := payloads.Decode(listFilter, r.Form)
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "Unable to decode request query parameters")
	}

	serviceInstanceList, err := h.serviceInstanceRepo.ListServiceInstances(r.Context(), authInfo, listFilter.ToMessage())
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "Failed to list service instance")
	}

	return routing.NewHandlerResponse(http.StatusOK).WithBody(presenter.ForServiceInstanceList(serviceInstanceList, h.serverURL, *r.URL)), nil
}

func (h *ServiceInstanceHandler) serviceInstanceDeleteHandler(r *http.Request) (*routing.Response, error) {
	authInfo, _ := authorization.InfoFromContext(r.Context())
	logger := logr.FromContextOrDiscard(r.Context()).WithName("service-instance-handler.service-instance-delete")

	serviceInstanceGUID := chi.URLParam(r, "guid")

	serviceInstance, err := h.serviceInstanceRepo.GetServiceInstance(r.Context(), authInfo, serviceInstanceGUID)
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, apierrors.ForbiddenAsNotFound(err), "failed to get service instance")
	}

	err = h.serviceInstanceRepo.DeleteServiceInstance(r.Context(), authInfo, repositories.DeleteServiceInstanceMessage{
		GUID:      serviceInstanceGUID,
		SpaceGUID: serviceInstance.SpaceGUID,
	})
	if err != nil {
		return nil, apierrors.LogAndReturn(logger, err, "error when deleting service instance", "guid", serviceInstanceGUID)
	}

	return routing.NewHandlerResponse(http.StatusNoContent), nil
}

func (h *ServiceInstanceHandler) RegisterRoutes(router *chi.Mux) {
	router.Method("POST", ServiceInstancesPath, routing.Handler(h.serviceInstanceCreateHandler))
	router.Method("GET", ServiceInstancesPath, routing.Handler(h.serviceInstanceListHandler))
	router.Method("DELETE", ServiceInstancePath, routing.Handler(h.serviceInstanceDeleteHandler))
}
