package postgresql

import (
	"client-service/internal/entity"
	"client-service/internal/pkg/otlp"
	"client-service/internal/pkg/postgres"
	"context"
	"database/sql"
	"fmt"
	"time"

	"client-service/internal/infrastructure/repository"
	"github.com/Masterminds/squirrel"
)

const (
	clientTableName       = "clients"
	clientsSpanRepoPrefix = "clientRepo"
)

type clientRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewClientsRepo(db *postgres.PostgresDB) repository.Clients {
	return &clientRepo{
		tableName: clientTableName,
		db:        db,
	}
}

// this is add select fields to query builder
func (p *clientRepo) clientsSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"first_name",
			"last_name",
			"age",
			"gender",
			"phone_number",
			"address",
			"email",
			"password",
			"status",
			"refresh",
			"created_at",
			"updated_at",
			"deleted_at",
		).From(p.tableName)
}

func (p clientRepo) CreateClient(ctx context.Context, client *entity.Client) (*entity.Client, error) {

	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "CreateClient")
	defer span.End()

	data := map[string]any{
		"id":           client.GUID,
		"first_name":   client.FirstName,
		"last_name":    client.LastName,
		"age":          client.Age,
		"gender":       client.Gender,
		"phone_number": client.PhoneNumber,
		"address":      client.Address,
		"email":        client.Email,
		"password":     client.Password,
		"status":       client.Status,
		"refresh":      client.Refresh,
		"created_at":   client.CreatedAt,
		"updated_at":   client.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}

	return client, nil
}

func (p clientRepo) UpdateClient(ctx context.Context, client *entity.Client) (*entity.Client, error) {

	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "UpdateClient")
	defer span.End()

	clauses := map[string]any{
		"first_name":   client.FirstName,
		"last_name":    client.LastName,
		"age":          client.Age,
		"gender":       client.Gender,
		"phone_number": client.PhoneNumber,
		"address":      client.Address,
		"email":        client.Email,
		"status":       client.Status,
		"updated_at":   client.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", client.GUID)).
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

	return client, nil
}

func (p clientRepo) DeleteClient(ctx context.Context, guid string) error {

	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "DeleteClient")
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

func (p clientRepo) GetClient(ctx context.Context, params map[string]string) (*entity.Client, error) {

	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "GetClient")
	defer span.End()

	var (
		client entity.Client
	)

	queryBuilder := p.clientsSelectQueryPrefix()

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
	var (
		nullAge         sql.NullInt32
		nullGender      sql.NullString
		nullPhoneNumber sql.NullString
		nullAddress     sql.NullString
		nullDeletedAt   sql.NullTime
	)
	if err = p.db.QueryRow(ctx, query, args...).Scan(
		&client.GUID,
		&client.FirstName,
		&client.LastName,
		&nullAge,
		&nullGender,
		&nullPhoneNumber,
		&nullAddress,
		&client.Email,
		&client.Password,
		&client.Status,
		&client.Refresh,
		&client.CreatedAt,
		&client.UpdatedAt,
		&nullDeletedAt,
	); err != nil {
		return nil, p.db.Error(err)
	}
	if nullAge.Valid {
		client.Age = uint64(nullAge.Int32)
	}
	if nullGender.Valid {
		client.Gender = nullGender.String
	}
	if nullPhoneNumber.Valid {
		client.PhoneNumber = nullPhoneNumber.String
	}
	if nullAddress.Valid {
		client.Address = nullAddress.String
	}
	if nullDeletedAt.Valid {
		client.DeletedAt = nullDeletedAt.Time
	}

	return &client, nil
}

func (p clientRepo) GetAllClients(ctx context.Context, limit uint64, offset uint64, filter map[string]string) ([]*entity.Client, error) {

	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "ListClients")
	defer span.End()

	var (
		clients []*entity.Client
	)
	queryBuilder := p.clientsSelectQueryPrefix()

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

	clients = make([]*entity.Client, 0)
	for rows.Next() {
		var (
			client          entity.Client
			nullAge         sql.NullInt32
			nullGender      sql.NullString
			nullPhoneNumber sql.NullString
			nullAddress     sql.NullString
			nullDeletedAt   sql.NullTime
		)
		if err = rows.Scan(
			&client.GUID,
			&client.FirstName,
			&client.LastName,
			&nullAge,
			&nullGender,
			&nullPhoneNumber,
			&nullAddress,
			&client.Email,
			&client.Password,
			&client.Status,
			&client.Refresh,
			&client.CreatedAt,
			&client.UpdatedAt,
			&nullDeletedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		if nullAge.Valid {
			client.Age = uint64(nullAge.Int32)
		}
		if nullGender.Valid {
			client.Gender = nullGender.String
		}
		if nullPhoneNumber.Valid {
			client.PhoneNumber = nullPhoneNumber.String
		}
		if nullAddress.Valid {
			client.Address = nullAddress.String
		}
		if nullDeletedAt.Valid {
			client.DeletedAt = nullDeletedAt.Time
		}

		clients = append(clients, &client)
	}

	return clients, nil
}

func (p clientRepo) GetAllDeletedClients(ctx context.Context, limit uint64, offset uint64, filter map[string]string) ([]*entity.Client, error) {
	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "ListDeletedClient")
	defer span.End()

	var (
		clients []*entity.Client
	)
	queryBuilder := p.clientsSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	// check is deleted
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

	clients = make([]*entity.Client, 0)
	for rows.Next() {
		var (
			client          entity.Client
			nullAge         sql.NullInt32
			nullGender      sql.NullString
			nullPhoneNumber sql.NullString
			nullAddress     sql.NullString
			nullDeletedAt   sql.NullTime
		)
		if err = rows.Scan(
			&client.GUID,
			&client.FirstName,
			&client.LastName,
			&nullAge,
			&nullGender,
			&nullPhoneNumber,
			&nullAddress,
			&client.Email,
			&client.Password,
			&client.Status,
			&client.Refresh,
			&client.CreatedAt,
			&client.UpdatedAt,
			&nullDeletedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		if nullAge.Valid {
			client.Age = uint64(nullAge.Int32)
		}
		if nullGender.Valid {
			client.Gender = nullGender.String
		}
		if nullPhoneNumber.Valid {
			client.PhoneNumber = nullPhoneNumber.String
		}
		if nullAddress.Valid {
			client.Address = nullAddress.String
		}
		if nullDeletedAt.Valid {
			client.DeletedAt = nullDeletedAt.Time
		}

		clients = append(clients, &client)
	}

	return clients, nil
}

func (p clientRepo) GetAllHiddenClients(ctx context.Context, limit uint64, offset uint64, filter map[string]bool) ([]*entity.Client, error) {
	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "ListHiddenClients")
	defer span.End()

	var (
		clients []*entity.Client
	)
	queryBuilder := p.clientsSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for key, value := range filter {
		if key == "status" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "list"))
	}

	fmt.Println(query)
	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			client          entity.Client
			nullAge         sql.NullInt32
			nullGender      sql.NullString
			nullPhoneNumber sql.NullString
			nullAddress     sql.NullString
			nullDeletedAt   sql.NullTime
		)
		if err = rows.Scan(
			&client.GUID,
			&client.FirstName,
			&client.LastName,
			&nullAge,
			&nullGender,
			&nullPhoneNumber,
			&nullAddress,
			&client.Email,
			&client.Password,
			&client.Status,
			&client.Refresh,
			&client.CreatedAt,
			&client.UpdatedAt,
			&nullDeletedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		if nullAge.Valid {
			client.Age = uint64(nullAge.Int32)
		}
		if nullGender.Valid {
			client.Gender = nullGender.String
		}
		if nullPhoneNumber.Valid {
			client.PhoneNumber = nullPhoneNumber.String
		}
		if nullAddress.Valid {
			client.Address = nullAddress.String
		}
		if nullDeletedAt.Valid {
			client.DeletedAt = nullDeletedAt.Time
		}

		clients = append(clients, &client)
	}

	return clients, nil
}

func (p clientRepo) UniqueEmail(ctx context.Context, request *entity.IsUnique) (*entity.Response, error) {
	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "UniqueEmail")
	defer span.End()

	var count uint64
	query := `SELECT COUNT(*) FROM clients WHERE email = $1`

	row := p.db.QueryRow(ctx, query, request.Email)
	if err := row.Scan(&count); err != nil {
		return &entity.Response{Status: false}, err
	}

	if count == 0 {
		return &entity.Response{Status: false}, nil
	}

	return &entity.Response{Status: true}, nil
}

func (p clientRepo) UpdateRefresh(ctx context.Context, request *entity.UpdateRefresh) (*entity.Response, error) {
	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "UpdateRefresh")
	defer span.End()

	clauses := map[string]any{
		"refresh":    request.RefreshToken,
		"updated_at": time.Now().Format(time.RFC3339),
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", request.ClientID)).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		return &entity.Response{Status: false}, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return &entity.Response{Status: false}, p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return &entity.Response{Status: false}, p.db.Error(fmt.Errorf("no sql rows"))
	}

	return &entity.Response{Status: true}, nil
}

func (p clientRepo) UpdatePassword(ctx context.Context, request *entity.UpdatePassword) (*entity.Response, error) {
	ctx, span := otlp.Start(ctx, clientsSpanRepoPrefix+"_grpc-repository", "UpdatePassword")
	defer span.End()

	clauses := map[string]any{
		"password":   request.NewPassword,
		"updated_at": time.Now().Format(time.RFC3339),
	}
	query, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", request.ClientID)).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		return &entity.Response{Status: false}, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	commandTag, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return &entity.Response{Status: false}, p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return &entity.Response{Status: false}, p.db.Error(fmt.Errorf("no sql rows"))
	}

	return &entity.Response{Status: true}, nil
}
