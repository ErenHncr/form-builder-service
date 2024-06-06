package types

import (
	"regexp"
	"strings"
)

type Sorting struct {
	Field string       `json:"field"`
	Order SortingOrder `json:"order"`
}

type SortingOrder string

const (
	OrderAsc  = "+"
	OrderDesc = "-"
)

var SortingOrderMap = map[SortingOrder]string{
	OrderAsc:  OrderAsc,
	OrderDesc: OrderDesc,
}

func (s SortingOrder) String() string {
	return SortingOrderMap[s]
}

func NewSorting(sort string) []Sorting {
	var sortingSlice []Sorting

	if sort == "" {
		return sortingSlice
	}

	sortSlice := strings.Split(sort, ",")
	for _, value := range sortSlice {
		var sorting Sorting
		sorting.Order = getSortingOrder(value)
		sorting.Field = sanitizeSortingField(value)

		if sorting.Field != "" && sorting.Order != "" {
			sortingSlice = append(sortingSlice, sorting)
		}
	}

	return sortingSlice
}

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`)

func sanitizeSortingField(val string) string {
	for _, value := range SortingOrderMap {
		val = strings.ReplaceAll(val, value, "")
		val = strings.TrimSpace(val)
	}

	if isLetter.FindString(val) == "" {
		return ""
	}

	return val
}

func getSortingOrder(value string) SortingOrder {
	isAscending := strings.HasPrefix(value, OrderAsc)
	isDescending := strings.HasPrefix(value, OrderDesc)

	var sortingOrder SortingOrder = OrderAsc
	if isAscending && isDescending {
		sortingOrder = OrderAsc
	} else if isDescending {
		sortingOrder = OrderDesc
	}

	return sortingOrder
}
