package fiber

import (
	"strconv"

	"github.com/ganigeorgiev/fexpr"
	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/utils/errors"
)

type QueryFilters struct {
	Conditions []fexpr.ExprGroup
	Page       int
	Limit      int
}

const (
	defaultLimit = 12
	maxLimit     = 35
	startPage    = 0
)

func ParseQueryFilter(c *fiber.Ctx) (*QueryFilters, error) {
	var filter = QueryFilters{}
	var err error

	filterQry := c.Query("filter")
	if filterQry != "" {
		if filter.Conditions, err = fexpr.Parse(filterQry); err != nil {
			return nil, errors.New("invalid filter passed in query params")
		}
	}

	if filter.Page, filter.Limit, err = parsePagination(c); err != nil {
		return nil, err
	}

	return &filter, nil
}

func parsePagination(c *fiber.Ctx) (int, int, error) {
	pageQry := c.Query("start")
	limitQry := c.Query("limit")

	page, limit := startPage, defaultLimit
	var err error

	if pageQry != "" {
		if page, err = strconv.Atoi(pageQry); err != nil {
			return 0, 0, errors.New("invalid start query params passed")
		}
	}
	if limitQry != "" {
		if limit, err = strconv.Atoi(limitQry); err != nil {
			return 0, 0, errors.New("invalid limit query params passed")
		}
	}

	if limit > maxLimit {
		return 0, 0, errors.New("limit is greater then max limit allowed")
	}

	return page, limit, nil
}
