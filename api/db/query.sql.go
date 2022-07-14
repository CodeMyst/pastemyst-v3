// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const createPaste = `-- name: CreatePaste :one
insert into pastes (
    id, created_at, title
) values (
    $1, $2, $3
)
returning id, created_at, title
`

type CreatePasteParams struct {
	ID        string
	CreatedAt time.Time
	Title     string
}

func (q *Queries) CreatePaste(ctx context.Context, arg CreatePasteParams) (Paste, error) {
	row := q.db.QueryRowContext(ctx, createPaste, arg.ID, arg.CreatedAt, arg.Title)
	var i Paste
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Title)
	return i, err
}

const createPasty = `-- name: CreatePasty :one
insert into pasties (
    id, paste_id, title, content
) values (
    $1, $2, $3, $4
)
returning id, paste_id, title, content
`

type CreatePastyParams struct {
	ID      string
	PasteID string
	Title   string
	Content string
}

func (q *Queries) CreatePasty(ctx context.Context, arg CreatePastyParams) (Pasty, error) {
	row := q.db.QueryRowContext(ctx, createPasty,
		arg.ID,
		arg.PasteID,
		arg.Title,
		arg.Content,
	)
	var i Pasty
	err := row.Scan(
		&i.ID,
		&i.PasteID,
		&i.Title,
		&i.Content,
	)
	return i, err
}

const existsPaste = `-- name: ExistsPaste :one
select exists(select 1 from pastes where id = $1)
`

func (q *Queries) ExistsPaste(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsPaste, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsUserByProvider = `-- name: ExistsUserByProvider :one
select exists(select 1 from users where provider_name = $1 and provider_id = $2)
`

type ExistsUserByProviderParams struct {
	ProviderName string
	ProviderID   string
}

func (q *Queries) ExistsUserByProvider(ctx context.Context, arg ExistsUserByProviderParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsUserByProvider, arg.ProviderName, arg.ProviderID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getPaste = `-- name: GetPaste :one
select id, created_at, title from pastes
where id = $1 limit 1
`

func (q *Queries) GetPaste(ctx context.Context, id string) (Paste, error) {
	row := q.db.QueryRowContext(ctx, getPaste, id)
	var i Paste
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Title)
	return i, err
}

const getPasteCount = `-- name: GetPasteCount :one
select count(*) from pastes
`

func (q *Queries) GetPasteCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPasteCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPastePasties = `-- name: GetPastePasties :many
select id, paste_id, title, content from pasties
where paste_id = $1
`

func (q *Queries) GetPastePasties(ctx context.Context, pasteID string) ([]Pasty, error) {
	rows, err := q.db.QueryContext(ctx, getPastePasties, pasteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pasty
	for rows.Next() {
		var i Pasty
		if err := rows.Scan(
			&i.ID,
			&i.PasteID,
			&i.Title,
			&i.Content,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByProvider = `-- name: GetUserByProvider :one
select id, created_at, username, avatar_url, provider_name, provider_id from users
where provider_name = $1 and provider_id = $2 limit 1
`

type GetUserByProviderParams struct {
	ProviderName string
	ProviderID   string
}

func (q *Queries) GetUserByProvider(ctx context.Context, arg GetUserByProviderParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByProvider, arg.ProviderName, arg.ProviderID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Username,
		&i.AvatarUrl,
		&i.ProviderName,
		&i.ProviderID,
	)
	return i, err
}
