package v1

import (
	_ "admin-api-gateway/api/docs"
	"context"
	"net/http"
	"strconv"
	"time"

	"admin-api-gateway/api/models"
	clientproto "admin-api-gateway/genproto/client_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary 	Create Client
// @Description This API for create a new client
// @Tags 		clients
// @Accept 		json
// @Produce 	json
// @Param 		Client body models.Client true "Client Model"
// @Success 	201 {object} string
// @Failure 	400 {object} models.Error
// @Failure    	401 {object} models.Error
// @Failure     403 {object} models.Error
// @Failure 	500 {object} models.Error
// @Router 		/v1/client [POST]
func (h HandlerV1) CreateClient(c *gin.Context) {
	var (
		body        models.Client
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	response, err := h.Service.ClientService().CreateClient(ctx, &clientproto.Client{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Age:         uint32(body.Age),
		Gender:      body.Gender,
		PhoneNumber: body.PhoneNumber,
		Address:     body.Address,
		Email:       body.Email,
		Password:    body.Password,
		Status:      body.Status,
		Refresh:     body.Refresh,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Guid)
}

// @Summary 	Update Client
// @Description This API for update a client
// @Tags 		clients
// @Accept 		json
// @Produce 	json
// @Param 		Client body models.Client true "Client Model"
// @Success 	200 {object} models.Client
// @Failure 	400 {object} models.Error
// @Failure    	401 {object} models.Error
// @Failure     403 {object} models.Error
// @Failure 	500 {object} models.Error
// @Router 		/v1/client [PUT]
func (h HandlerV1) UpdateClient(c *gin.Context) {
	var (
		body        models.Client
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	response, err := h.Service.ClientService().UpdateClient(ctx, &clientproto.Client{
		Id:          body.Id,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Age:         uint32(body.Age),
		Gender:      body.Gender,
		PhoneNumber: body.PhoneNumber,
		Address:     body.Address,
		Email:       body.Email,
		Password:    body.Password,
		Status:      body.Status,
		Refresh:     body.Refresh,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 	Delete Client
// @Description This API for delete a client
// @Tags 		clients
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Client ID"
// @Success 	200 {object} models.Status
// @Failure 	400 {object} models.Error
// @Failure    	401 {object} models.Error
// @Failure     403 {object} models.Error
// @Failure 	500 {object} models.Error
// @Router 		/v1/client/{id} [DELETE]
func (h HandlerV1) DeleteClient(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	clientID := c.Param("id")

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	_, err = h.Service.ClientService().DeleteClient(ctx, &clientproto.ClientWithGUID{
		Guid: clientID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Status{Status: true})
}

// @Summary 	Get Client
// @Description This API for get a client
// @Tags 		clients
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Client ID"
// @Success 	200 {object} models.Client
// @Failure 	400 {object} models.Error
// @Failure    	401 {object} models.Error
// @Failure     403 {object} models.Error
// @Failure 	500 {object} models.Error
// @Router 		/v1/client/{id} [GET]
func (h HandlerV1) GetClient(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	clientID := c.Param("id")

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	response, err := h.Service.ClientService().GetClient(ctx, &clientproto.ClientWithGUID{
		Guid: clientID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 	List Clients
// @Description This API for get a list of clients
// @Tags 		clients
// @Accept 		json
// @Produce 	json
// @Param 		page query uint64 true "Page"
// @Param 		limit query uint64 true "Limit"
// @Success 	200 {object} []models.Client
// @Failure 	400 {object} models.Error
// @Failure    	401 {object} models.Error
// @Failure     403 {object} models.Error
// @Failure 	500 {object} models.Error
// @Router 		/v1/clients/active [GET]
func (h HandlerV1) ListClients(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	page := c.Query("page")
	limit := c.Query("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	listClients, err := h.Service.ClientService().GetAllClients(ctx, &clientproto.ListRequest{
		Page:  int64(pageInt),
		Limit: int64(limitInt),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	var response []*models.Client
	for _, client := range listClients.Clients {
		response = append(response, &models.Client{
			Id:          client.Id,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			Age:         uint64(client.Age),
			Gender:      client.Gender,
			PhoneNumber: client.PhoneNumber,
			Address:     client.Address,
			Email:       client.Email,
			Password:    client.Password,
			Status:      client.Status,
			Refresh:     client.Refresh,
		})
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 	List Deleted Clients
// @Description This API for get a list of deleted clients
// @Tags 		clients
// @Accept 		json
// @Produce 	json
// @Param 		page query uint64 true "Page"
// @Param 		limit query uint64 true "Limit"
// @Success 	200 {object} []models.Client
// @Failure 	400 {object} models.Error
// @Failure    	401 {object} models.Error
// @Failure     403 {object} models.Error
// @Failure 	500 {object} models.Error
// @Router 		/v1/clients/deleted [GET]
func (h HandlerV1) ListDeletedClients(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	page := c.Query("page")
	limit := c.Query("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	listClients, err := h.Service.ClientService().GetAllDeletedClients(ctx, &clientproto.ListRequest{
		Page:  int64(pageInt),
		Limit: int64(limitInt),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	var response []*models.Client
	for _, client := range listClients.Clients {
		response = append(response, &models.Client{
			Id:          client.Id,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			Age:         uint64(client.Age),
			Gender:      client.Gender,
			PhoneNumber: client.PhoneNumber,
			Address:     client.Address,
			Email:       client.Email,
			Password:    client.Password,
			Status:      client.Status,
			Refresh:     client.Refresh,
		})
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 	List Hidden Clients
// @Description This API for get a list of hidden clients
// @Tags 		clients
// @Accept 		json
// @Produce 	json
// @Param 		page query uint64 true "Page"
// @Param 		limit query uint64 true "Limit"
// @Success 	200 {object} []models.Client
// @Failure 	400 {object} models.Error
// @Failure    	401 {object} models.Error
// @Failure     403 {object} models.Error
// @Failure 	500 {object} models.Error
// @Router 		/v1/clients/hidden [GET]
func (h HandlerV1) ListHiddenClients(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	page := c.Query("page")
	limit := c.Query("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	listClients, err := h.Service.ClientService().GetAllHiddenClients(ctx, &clientproto.ListRequest{
		Page:  int64(pageInt),
		Limit: int64(limitInt),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	var response []*models.Client
	for _, client := range listClients.Clients {
		response = append(response, &models.Client{
			Id:          client.Id,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			Age:         uint64(client.Age),
			Gender:      client.Gender,
			PhoneNumber: client.PhoneNumber,
			Address:     client.Address,
			Email:       client.Email,
			Password:    client.Password,
			Status:      client.Status,
			Refresh:     client.Refresh,
		})
	}
	c.JSON(http.StatusOK, response)
}
