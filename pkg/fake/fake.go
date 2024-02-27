// Package fake provides the ability to create fake vehicle data.
package fake

import (
	"fmt"
	"time"

	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/brianvoe/gofakeit/v7/source"
)

// ElaticStatus will create a model.DataPoint with fake data.
func ElaticStatus() (*model.DataPoint, error) {
	var dataPoint model.DataPoint
	err := gofakeit.Struct(&dataPoint)
	if err != nil {
		return nil, fmt.Errorf("error creating fake data point: %w", err)
	}
	return &dataPoint, nil
}

// Setup will setup the fake package.
func Setup() {
	gofakeit.GlobalFaker = gofakeit.NewFaker(source.NewCrypto(), true)
	gofakeit.AddFuncLookup("timeRangeStr", gofakeit.Info{
		Category:    "timeRangeStr",
		Description: "A random time string between the given years",
		Example:     "2010-01-01T00:00:00Z",
		Output:      "string",
		Params: []gofakeit.Param{
			{
				Field:   "startYear",
				Type:    "int",
				Default: "2010",
			},
			{
				Field:   "endYear",
				Type:    "int",
				Default: "2020",
			},
		},
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			startYear, err := info.GetInt(m, "startYear")
			if err != nil {
				return "", fmt.Errorf("error getting startYear: %w", err)
			}
			endYear, err := info.GetInt(m, "endYear")
			if err != nil {
				return "", fmt.Errorf("error getting endYear: %w", err)
			}

			date := f.DateRange(time.Date(startYear, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(endYear, 1, 1, 0, 0, 0, 0, time.UTC))
			return date.Format(time.RFC3339), nil
		},
	})
}
