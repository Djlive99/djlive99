package stream

import (
	"database/sql"
	goerrors "errors"
	"fmt"

	"npm/internal/database"
	"npm/internal/entity"
	"npm/internal/errors"
	"npm/internal/logger"
	"npm/internal/model"
)

// GetByID finds a auth by ID
func GetByID(id int) (Model, error) {
	var m Model
	err := m.LoadByID(id)
	return m, err
}

// Create will create a Auth from this model
func Create(host *Model) (int, error) {
	if host.ID != 0 {
		return 0, goerrors.New("Cannot create stream when model already has an ID")
	}

	host.Touch(true)

	db := database.GetInstance()
	// nolint: gosec
	result, err := db.NamedExec(`INSERT INTO `+fmt.Sprintf("`%s`", tableName)+` (
		created_on,
		modified_on,
		user_id,
		provider,
		name,
		domain_names,
		expires_on,
		meta,
		is_deleted
	) VALUES (
		:created_on,
		:modified_on,
		:user_id,
		:provider,
		:name,
		:domain_names,
		:expires_on,
		:meta,
		:is_deleted
	)`, host)

	if err != nil {
		return 0, err
	}

	last, lastErr := result.LastInsertId()
	if lastErr != nil {
		return 0, lastErr
	}

	return int(last), nil
}

// Update will Update a Host from this model
func Update(host *Model) error {
	if host.ID == 0 {
		return goerrors.New("Cannot update stream when model doesn't have an ID")
	}

	host.Touch(false)

	db := database.GetInstance()
	// nolint: gosec
	_, err := db.NamedExec(`UPDATE `+fmt.Sprintf("`%s`", tableName)+` SET
		created_on = :created_on,
		modified_on = :modified_on,
		user_id = :user_id,
		provider = :provider,
		name = :name,
		domain_names = :domain_names,
		expires_on = :expires_on,
		meta = :meta,
		is_deleted = :is_deleted
	WHERE id = :id`, host)

	return err
}

// List will return a list of hosts
func List(pageInfo model.PageInfo, filters []model.Filter) (ListResponse, error) {
	var result ListResponse
	var exampleModel Model

	defaultSort := model.Sort{
		Field:     "name",
		Direction: "ASC",
	}

	db := database.GetInstance()
	if db == nil {
		return result, errors.ErrDatabaseUnavailable
	}

	// Get count of items in this search
	query, params := entity.ListQueryBuilder(exampleModel, tableName, &pageInfo, defaultSort, filters, getFilterMapFunctions(), true)
	countRow := db.QueryRowx(query, params...)
	var totalRows int
	queryErr := countRow.Scan(&totalRows)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		logger.Debug("%s -- %+v", query, params)
		return result, queryErr
	}

	// Get rows
	var items []Model
	query, params = entity.ListQueryBuilder(exampleModel, tableName, &pageInfo, defaultSort, filters, getFilterMapFunctions(), false)
	err := db.Select(&items, query, params...)
	if err != nil {
		logger.Debug("%s -- %+v", query, params)
		return result, err
	}

	result = ListResponse{
		Items:  items,
		Total:  totalRows,
		Limit:  pageInfo.Limit,
		Offset: pageInfo.Offset,
		Sort:   pageInfo.Sort,
		Filter: filters,
	}

	return result, nil
}
