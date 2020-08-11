package druid

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gerfg/go-druid/druid/filters"
)

type scanQuery struct {
	QueryType  string
	DataSource struct {
		Type string
		Name string
	}
	Intervals    []Interval
	ResultFormat *string
	Columns      *[]string
	Limit        *int64
	Filters      *filters.FilterAggregator

	json string
}

func NewScanQuery(
	dataSourceType string,
	dataSourceName string,
	intervals []Interval,
	resultFormat string,
	columns []string,
	limit int64,
	filters filters.FilterAggregator,
) *scanQuery {
	return &scanQuery{
		QueryType: "scan",
		DataSource: struct {
			Type string
			Name string
		}{Type: dataSourceType, Name: dataSourceName},
		Intervals:    intervals,
		ResultFormat: &resultFormat,
		Columns:      &columns,
		Limit:        &limit,
		Filters:      &filters,
	}
}

func (s *scanQuery) JSON() string {
	return s.json
}

func (s *scanQuery) CreateNativeQuery() error {
	s.header()

	if s.ResultFormat != nil {
		if err := s.resultFormat(); err != nil {
			return err
		}
	}

	if s.Columns != nil {
		s.columns()
	}

	if s.Limit != nil {
		s.limit(false)
	} else {
		s.limit(true)
	}

	if s.Filters != nil {
		s.filters()
	}

	s.json += "}"

	return nil
}

func (s *scanQuery) header() {
	s.json = fmt.Sprintf(
		`{"queryType": "scan", "dataSource": {"type": "%s","name": "%s"}, "intervals": {"type": "intervals", "intervals": [`,
		s.DataSource.Type, s.DataSource.Name,
	)

	for _, interval := range s.Intervals {
		s.json += fmt.Sprintf(`"%s/%s", `, interval.begin.Format(time.RFC3339), interval.end.Format(time.RFC3339))
	}

	s.json = strings.TrimSuffix(s.json, ", ")

	s.json += `]}`
}

func (s *scanQuery) resultFormat() error {
	if *s.ResultFormat != "list" && *s.ResultFormat != "compactedList" {
		return errors.New("Invalid resultFormat, must be: list or compactedList.")
	}
	s.json += fmt.Sprintf(`, "resultFormat": "%s"`, *s.ResultFormat)
	return nil
}

func (s *scanQuery) columns() {
	s.json += `, "columns": ["__time", `

	for _, col := range *s.Columns {
		s.json += fmt.Sprintf(`"%s", `, col)
	}

	s.json = strings.TrimSuffix(s.json, ", ")
	s.json += `]`
}

func (s *scanQuery) limit(defaultLimit bool) {
	if defaultLimit {
		s.json += `, "limit": 10`
	} else {
		s.json += fmt.Sprintf(`, "limit": %d`, *s.Limit)
	}
}

func (s *scanQuery) filters() {
	filterString, err := s.Filters.ToFilter()
	if err != nil {
		return
	}

	s.json += filterString
}
