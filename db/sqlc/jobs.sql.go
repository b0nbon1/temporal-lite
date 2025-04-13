// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: jobs.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createJob = `-- name: CreateJob :one
INSERT INTO jobs (
  name,
  schedule,
  type,
  config,
  status,
  retries,
  max_retries
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING id, name, schedule, type, config, status, retries, max_retries, created_at, updated_at
`

type CreateJobParams struct {
	Name       string        `json:"name"`
	Schedule   pgtype.Text   `json:"schedule"`
	Type       pgtype.Text   `json:"type"`
	Config     []byte        `json:"config"`
	Status     NullJobStatus `json:"status"`
	Retries    pgtype.Int4   `json:"retries"`
	MaxRetries pgtype.Int4   `json:"max_retries"`
}

func (q *Queries) CreateJob(ctx context.Context, arg CreateJobParams) (Job, error) {
	row := q.db.QueryRow(ctx, createJob,
		arg.Name,
		arg.Schedule,
		arg.Type,
		arg.Config,
		arg.Status,
		arg.Retries,
		arg.MaxRetries,
	)
	var i Job
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Schedule,
		&i.Type,
		&i.Config,
		&i.Status,
		&i.Retries,
		&i.MaxRetries,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteJob = `-- name: DeleteJob :exec
DELETE FROM jobs
WHERE id = $1
`

func (q *Queries) DeleteJob(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteJob, id)
	return err
}

const getJob = `-- name: GetJob :one
SELECT id, name, schedule, type, config, status, retries, max_retries, created_at, updated_at FROM jobs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetJob(ctx context.Context, id pgtype.UUID) (Job, error) {
	row := q.db.QueryRow(ctx, getJob, id)
	var i Job
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Schedule,
		&i.Type,
		&i.Config,
		&i.Status,
		&i.Retries,
		&i.MaxRetries,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listJobs = `-- name: ListJobs :many
SELECT id, name, schedule, type, config, status, retries, max_retries, created_at, updated_at FROM jobs
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListJobsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListJobs(ctx context.Context, arg ListJobsParams) ([]Job, error) {
	rows, err := q.db.Query(ctx, listJobs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Job{}
	for rows.Next() {
		var i Job
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Schedule,
			&i.Type,
			&i.Config,
			&i.Status,
			&i.Retries,
			&i.MaxRetries,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPendingJobs = `-- name: ListPendingJobs :many
SELECT id, name, schedule, type, config, status, retries, max_retries, created_at, updated_at FROM jobs
where status = 'pending' AND 
      (schedule IS NOT NULL OR schedule != '')
ORDER BY created_at DESC
`

func (q *Queries) ListPendingJobs(ctx context.Context) ([]Job, error) {
	rows, err := q.db.Query(ctx, listPendingJobs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Job{}
	for rows.Next() {
		var i Job
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Schedule,
			&i.Type,
			&i.Config,
			&i.Status,
			&i.Retries,
			&i.MaxRetries,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateJobStatus = `-- name: UpdateJobStatus :one
UPDATE jobs
SET
  status = $2,
  retries = $3,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, schedule, type, config, status, retries, max_retries, created_at, updated_at
`

type UpdateJobStatusParams struct {
	ID      pgtype.UUID   `json:"id"`
	Status  NullJobStatus `json:"status"`
	Retries pgtype.Int4   `json:"retries"`
}

func (q *Queries) UpdateJobStatus(ctx context.Context, arg UpdateJobStatusParams) (Job, error) {
	row := q.db.QueryRow(ctx, updateJobStatus, arg.ID, arg.Status, arg.Retries)
	var i Job
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Schedule,
		&i.Type,
		&i.Config,
		&i.Status,
		&i.Retries,
		&i.MaxRetries,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
