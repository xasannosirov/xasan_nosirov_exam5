package v1

import (
	_ "admin-api-gateway/api/docs"
	"admin-api-gateway/api/models"
	clientproto "admin-api-gateway/genproto/client_service"
	jobproto "admin-api-gateway/genproto/job_service"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"strconv"
	"time"
)

// @Summary 		Create Job
// @Description 	This API for create a new job
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           Job body models.Job true "Job Model"
// @Success 		201 {object} string
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/job [POST]
func (h HandlerV1) CreateJob(c *gin.Context) {
	var (
		body        models.Job
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	if err := c.ShouldBindJSON(&body); err != nil {
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

	response, err := h.Service.JobService().CreateJob(ctx, &jobproto.Job{
		Name:           body.Name,
		Salary:         body.Salary,
		Level:          body.Level,
		LocationType:   body.LocationType,
		EmploymentType: body.EmploymentType,
		Address:        body.Address,
		Company:        body.Company,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary 		Update Job
// @Description 	This API for update a job
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           Job body models.Job true "Job Model"
// @Success 		200 {object} models.Job
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/job [PUT]
func (h HandlerV1) UpdateJob(c *gin.Context) {
	var (
		body        models.Job
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	if err := c.ShouldBindJSON(&body); err != nil {
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

	response, err := h.Service.JobService().UpdateJob(ctx, &jobproto.Job{
		Id:             body.ID,
		Name:           body.Name,
		Salary:         body.Salary,
		Level:          body.Level,
		LocationType:   body.LocationType,
		EmploymentType: body.EmploymentType,
		Address:        body.Address,
		Company:        body.Company,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

// @Summary 		Delete Job
// @Description 	This API for delete a job
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           id path string true "Job ID"
// @Success 		200 {object} models.Job
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/job/{id} [DELETE]
func (h HandlerV1) DeleteJob(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	jobID := c.Param("id")

	_, err = h.Service.JobService().DeleteJob(ctx, &jobproto.JobWithGUID{
		JobId: jobID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Status{Status: true})
}

// @Summary 		Get Job
// @Description 	This API for get a job
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           id path string true "Job ID"
// @Success 		200 {object} models.Job
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/job/{id} [GET]
func (h HandlerV1) GetJob(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	jobID := c.Param("id")

	response, err := h.Service.JobService().GetJob(ctx, &jobproto.JobWithGUID{
		JobId: jobID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		List Jobs
// @Description 	This API for get a list jobs
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           page query string true "Page"
// @Param 			limit query string true "Limit"
// @Success 		200 {object} []models.Job
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/jobs/active [GET]
func (h HandlerV1) ListJobs(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	page := c.Query("page")
	limt := c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	limitInt, err := strconv.Atoi(limt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	response, err := h.Service.JobService().GetAllJobs(ctx, &jobproto.ListRequest{
		Page:  uint64(pageInt),
		Limit: uint64(limitInt),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	var listJobs []*models.Job
	for _, job := range response.Jobs {
		listJobs = append(listJobs, &models.Job{
			ID:             job.Id,
			Name:           job.Name,
			Salary:         job.Salary,
			Level:          job.Level,
			LocationType:   job.LocationType,
			EmploymentType: job.EmploymentType,
			Address:        job.Address,
			Company:        job.Company,
		})
	}

	c.JSON(http.StatusOK, listJobs)
}

// @Summary 		List Deleted Jobs
// @Description 	This API for get list deleted jobs
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           page query string true "Page"
// @Param 			limit query string true "Limit"
// @Success 		200 {object} []models.Job
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/jobs/deleted [GET]
func (h HandlerV1) ListDeletedJobs(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	duration, err := time.ParseDuration(h.Config.Context.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	page := c.Query("page")
	limt := c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	limitInt, err := strconv.Atoi(limt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	response, err := h.Service.JobService().GetAllDeletedJobs(ctx, &jobproto.ListRequest{
		Page:  uint64(pageInt),
		Limit: uint64(limitInt),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	var listJobs []*models.Job
	for _, job := range response.Jobs {
		listJobs = append(listJobs, &models.Job{
			ID:             job.Id,
			Name:           job.Name,
			Salary:         job.Salary,
			Level:          job.Level,
			LocationType:   job.LocationType,
			EmploymentType: job.EmploymentType,
			Address:        job.Address,
			Company:        job.Company,
		})
	}

	c.JSON(http.StatusOK, listJobs)
}

// @Summary 		Add Client to Job
// @Description 	This API for add client to job
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           ClientJob body models.ClientJobs true "Client Job Model"
// @Success 		201 {object} []models.Status
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/job/add-client [POST]
func (h HandlerV1) AddClientToJob(c *gin.Context) {
	var (
		body        models.ClientJobs
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	if err := c.ShouldBindJSON(&body); err != nil {
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

	response, err := h.Service.JobService().AddClientJob(ctx, &jobproto.ClientJobs{
		ClientId:  body.ClientID,
		JobId:     body.JobID,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary 		Remove Client from Job
// @Description 	This API for remove client from job
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           ClientJob body models.ClientJobs true "Client Job Model"
// @Success 		200 {object} []models.Status
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/job/remove-client [DELETE]
func (h HandlerV1) RemoveClientFromJob(c *gin.Context) {
	var (
		body        models.ClientJobs
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	if err := c.ShouldBindJSON(&body); err != nil {
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

	_, err = h.Service.JobService().DeleteClientJob(ctx, &jobproto.ClientJobs{
		ClientId:  body.ClientID,
		JobId:     body.JobID,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.Status{Status: true})
}

// @Summary 		Get Clients with Job
// @Description 	This API for get clients with job-id
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           ClientJobRequest body models.ClientJobRequest true "Client Job Request"
// @Success 		200 {object} models.ClientWithJobs
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/jobs/client-jobs [POST]
func (h HandlerV1) GetClientsWithJob(c *gin.Context) {
	var (
		body        models.ClientJobRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	if err := c.ShouldBindJSON(&body); err != nil {
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

	clientJobs, err := h.Service.JobService().GetClientJobs(ctx, &jobproto.ClientJobRequest{
		ClientId: body.ClientID,
		Page:     uint64(body.Page),
		Limit:    uint64(body.Limit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	var response models.ClientWithJobs
	client, err := h.Service.ClientService().GetClient(ctx, &clientproto.ClientWithGUID{
		Guid: body.ClientID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	response.Client = models.Client{
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
	}
	for _, clientjob := range clientJobs.ClientJobs {
		job, err := h.Service.JobService().GetJob(ctx, &jobproto.JobWithGUID{
			JobId: clientjob.JobId,
		})
		if err != nil {
			log.Println("job deleted", clientjob.JobId)
			continue
		}
		startDate, err := time.Parse(time.RFC3339, clientjob.StartDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Error{
				Message: err.Error(),
			})
			return
		}
		endDate, err := time.Parse(time.RFC3339, clientjob.EndDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Error{
				Message: err.Error(),
			})
			return
		}

		response.Jobs = append(response.Jobs, models.ResponseJob{
			ID:             job.Id,
			Name:           job.Name,
			Salary:         job.Salary,
			Level:          job.Level,
			LocationType:   job.LocationType,
			EmploymentType: job.EmploymentType,
			Address:        job.Address,
			Company:        job.Company,
			StartDate:      startDate,
			EndDate:        endDate,
		})
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Get Jobs with Client
// @Description 	This API for get jobs with client-id
// @Tags 			jobs
// @Accept 			json
// @Produce 		json
// @Param           ClientJobRequest body models.ClientJobRequest true "Client Job Request"
// @Success 		200 {object} models.JobWithClients
// @Failure 		400 {object} models.Error
// @Failure 		401 {object} models.Error
// @Failure 		403 {object} models.Error
// @Failure 		500 {object} models.Error
// @Router 			/v1/jobs/job-clients [POST]
func (h HandlerV1) GetJobsWithClient(c *gin.Context) {
	var (
		body        models.ClientJobRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	if err := c.ShouldBindJSON(&body); err != nil {
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

	jobClients, err := h.Service.JobService().GetJobClients(ctx, &jobproto.ClientJobRequest{
		JobId: body.JobID,
		Page:  uint64(body.Page),
		Limit: uint64(body.Limit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	var response models.JobWithClients
	job, err := h.Service.JobService().GetJob(ctx, &jobproto.JobWithGUID{
		JobId: body.JobID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	startDate, err := time.Parse(time.RFC3339, jobClients.ClientJobs[0].StartDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}
	endDate, err := time.Parse(time.RFC3339, jobClients.ClientJobs[0].EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	response.Job = models.ResponseJob{
		ID:             job.Id,
		Name:           job.Name,
		Salary:         job.Salary,
		Level:          job.Level,
		LocationType:   job.LocationType,
		EmploymentType: job.EmploymentType,
		Address:        job.Address,
		Company:        job.Company,
		StartDate:      startDate,
		EndDate:        endDate,
	}

	for _, clientInfo := range jobClients.ClientJobs {
		client, err := h.Service.ClientService().GetClient(ctx, &clientproto.ClientWithGUID{
			Guid: clientInfo.ClientId,
		})
		if err != nil {
			log.Println("client deleted", clientInfo.ClientId)
			continue
		}
		response.Clients = append(response.Clients, models.Client{
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
