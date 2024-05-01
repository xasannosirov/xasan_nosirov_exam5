package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"job-service/internal/entity"
	"job-service/internal/infrastructure/repository"
	"job-service/internal/pkg/otlp"
	"job-service/internal/pkg/postgres"
	"time"
)

const (
	jobTableName       = "jobs"
	clientJobTableName = "client_jobs"
	jobsSpanRepoPrefix = "jobsRepo"
)

type jobRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewJobsRepo(db *postgres.PostgresDB) repository.Jobs {
	return &jobRepo{
		tableName: jobTableName,
		db:        db,
	}
}

func (p *jobRepo) jobsSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"name",
			"salary",
			"level",
			"location_type",
			"employment_type",
			"address",
			"company",
			"created_at",
			"updated_at",
		).From(p.tableName)
}

func (p *jobRepo) clientJobsSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.Select(
		"client_id",
		"job_id",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	).From(clientJobTableName)
}

func (p jobRepo) CreateJob(ctx context.Context, job *entity.Job) (*entity.Job, error) {

	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "CreateJob")
	defer span.End()

	data := map[string]any{
		"id":              job.GUID,
		"name":            job.Name,
		"salary":          job.Salary,
		"level":           job.Level,
		"location_type":   job.LocationType,
		"employment_type": job.EmploymentType,
		"address":         job.Address,
		"company":         job.Company,
		"created_at":      job.CreatedAt,
		"updated_at":      job.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}

	return job, nil
}

func (p jobRepo) UpdateJob(ctx context.Context, job *entity.Job) (*entity.Job, error) {

	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "UpdateJob")
	defer span.End()

	clauses := map[string]any{
		"name":            job.Name,
		"salary":          job.Salary,
		"level":           job.Level,
		"location_type":   job.LocationType,
		"employment_type": job.EmploymentType,
		"address":         job.Address,
		"company":         job.Company,
		"updated_at":      job.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", job.GUID)).
		ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return nil, p.db.Error(fmt.Errorf("no sql rows"))
	}

	return job, nil
}

func (p jobRepo) DeleteJob(ctx context.Context, guid string) error {

	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "DeleteJob")
	defer span.End()

	clauses := map[string]any{
		"deleted_at": time.Now().Format(time.RFC3339),
	}

	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", guid)).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" delete")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}

func (p jobRepo) GetJob(ctx context.Context, params map[string]string) (*entity.Job, error) {

	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "GetJob")
	defer span.End()

	var (
		job entity.Job
	)

	queryBuilder := p.jobsSelectQueryPrefix()

	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
	}
	queryBuilder = queryBuilder.Where("deleted_at IS NULL")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "get"))
	}

	if err = p.db.QueryRow(ctx, query, args...).Scan(
		&job.GUID,
		&job.Name,
		&job.Salary,
		&job.Level,
		&job.LocationType,
		&job.EmploymentType,
		&job.Address,
		&job.Company,
		&job.CreatedAt,
		&job.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &job, nil
}

func (p jobRepo) GetAllJobs(ctx context.Context, limit uint64, offset uint64, filter map[string]string) ([]*entity.Job, error) {

	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "ListJobs")
	defer span.End()

	var (
		jobs []*entity.Job
	)
	queryBuilder := p.jobsSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	queryBuilder = queryBuilder.Where("deleted_at IS NULL")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	jobs = make([]*entity.Job, 0)
	for rows.Next() {
		var job entity.Job
		if err = rows.Scan(
			&job.GUID,
			&job.Name,
			&job.Salary,
			&job.Level,
			&job.LocationType,
			&job.EmploymentType,
			&job.Address,
			&job.Company,
			&job.CreatedAt,
			&job.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}

		jobs = append(jobs, &job)
	}

	return jobs, nil
}

func (p jobRepo) GetAllDeletedJobs(ctx context.Context, limit uint64, offset uint64, filter map[string]string) ([]*entity.Job, error) {
	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "ListDeletedJobs")
	defer span.End()

	var (
		jobs []*entity.Job
	)
	queryBuilder := p.jobsSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	queryBuilder = queryBuilder.Where("deleted_at IS NOT NULL")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	jobs = make([]*entity.Job, 0)
	for rows.Next() {
		var job entity.Job
		if err = rows.Scan(
			&job.GUID,
			&job.Name,
			&job.Salary,
			&job.Level,
			&job.LocationType,
			&job.EmploymentType,
			&job.Address,
			&job.Company,
			&job.CreatedAt,
			&job.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}

		jobs = append(jobs, &job)
	}

	return jobs, nil
}

func (p jobRepo) GetClientJobs(ctx context.Context, limit uint64, offset uint64, filter map[string]string) ([]*entity.ClientJob, error) {
	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "GetClientJobs")
	defer span.End()

	var (
		clientJobs []*entity.ClientJob
	)
	queryBuilder := p.clientJobsSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	queryBuilder = queryBuilder.Limit(limit)
	queryBuilder = queryBuilder.Offset(offset)
	for key, value := range filter {
		if key == "client_id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
	}

	queryBuilder = queryBuilder.Where("deleted_at IS NULL")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", clientJobTableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			nullEndDate sql.NullString
			clientJob   entity.ClientJob
		)
		if err = rows.Scan(
			&clientJob.ClientID,
			&clientJob.JobID,
			&clientJob.StartDate,
			&nullEndDate,
			&clientJob.CreatedAt,
			&clientJob.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		if nullEndDate.Valid {
			endDate, err := time.Parse(time.RFC3339, nullEndDate.String)
			if err != nil {
				return nil, p.db.Error(err)
			}
			clientJob.EndDate = endDate
		}

		clientJobs = append(clientJobs, &clientJob)
	}

	return clientJobs, nil
}

func (p jobRepo) GetJobClients(ctx context.Context, limit uint64, offset uint64, filter map[string]string) ([]*entity.ClientJob, error) {
	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "ListJobClients")
	defer span.End()

	var (
		clientJobs []*entity.ClientJob
	)
	queryBuilder := p.clientJobsSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	queryBuilder = queryBuilder.Limit(limit)
	queryBuilder = queryBuilder.Offset(offset)
	for key, value := range filter {
		if key == "job_id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
	}

	queryBuilder = queryBuilder.Where("deleted_at IS NULL")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", clientJobTableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			nullEndDate sql.NullString
			clientJob   entity.ClientJob
		)
		if err = rows.Scan(
			&clientJob.ClientID,
			&clientJob.JobID,
			&clientJob.StartDate,
			&nullEndDate,
			&clientJob.CreatedAt,
			&clientJob.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		if nullEndDate.Valid {
			endDate, err := time.Parse(time.RFC3339, nullEndDate.String)
			if err != nil {
				return nil, p.db.Error(err)
			}
			clientJob.EndDate = endDate
		}

		clientJobs = append(clientJobs, &clientJob)
	}

	return clientJobs, nil
}

func (p jobRepo) AddClientJob(ctx context.Context, clientJob *entity.ClientJob) (*entity.Response, error) {
	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "AddClientJob")
	span.End()

	data := map[string]any{
		"client_id":  clientJob.ClientID,
		"job_id":     clientJob.JobID,
		"start_date": clientJob.StartDate,
		"end_date":   clientJob.EndDate,
	}
	query, args, err := p.db.Sq.Builder.Insert(clientJobTableName).SetMap(data).ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", clientJobTableName, "create"))
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return &entity.Response{Status: false}, p.db.Error(err)
	}

	return &entity.Response{Status: true}, nil
}

func (p jobRepo) DeleteClientJob(ctx context.Context, clientJob *entity.ClientJob) error {
	ctx, span := otlp.Start(ctx, jobsSpanRepoPrefix+"_grpc-repository", "DeleteClientJob")
	defer span.End()

	clauses := map[string]any{
		"deleted_at": time.Now().Format(time.RFC3339),
	}

	sqlStr, args, err := p.db.Sq.Builder.
		Update(clientJobTableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("client_id", clientJob.ClientID)).
		Where(p.db.Sq.Equal("job_id", clientJob.JobID)).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, clientJobTableName+" delete")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	fmt.Println(commandTag.RowsAffected())

	return nil
}
