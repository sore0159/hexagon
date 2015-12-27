package hexagon

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Scan impliments Scanner for use in SQL queries
func (c *Coord) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Bad value scanned to hexagon coord:", value))
	}
	parts := strings.Split(string(bytes), ",")
	x, err := strconv.Atoi(parts[0][1:])
	if err != nil {
		return errors.New(fmt.Sprint("bad hexagon coord scan value:", parts, err))
	}
	y, err := strconv.Atoi(parts[1][:len(parts[1])-1])
	if err != nil {
		return errors.New(fmt.Sprint("bad hexagon coord scan value:", parts, err))
	}
	c[0], c[1] = x, y
	return nil
}

// Value impliments Valuer for SQL queries
func (c Coord) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%d,%d)", c[0], c[1]), nil
}

func (c Coord) SQLStr() string {
	return fmt.Sprintf("POINT(%d,%d)", c[0], c[1])
}

func CoordList2Sql(list []Coord) string {
	if len(list) == 0 {
		return "ARRAY[]::point[]"
	}
	listStr := "ARRAY["
	parts := make([]string, len(list))
	for i, c := range list {
		parts[i] = fmt.Sprintf("POINT(%d, %d)", c[0], c[1])
	}
	listStr += strings.Join(parts, ", ") + "]"
	return listStr
}

func Sql2CoordList(bytes []byte) (list []Coord, ok bool) {
	listStr := string(bytes)
	if listStr == "{}" {
		return []Coord{}, true
	}
	parts := strings.Split(listStr, ",")
	if len(parts)%2 != 0 {
		return nil, false
	}
	var odd bool
	var x, y int
	list = []Coord{}
	for _, part := range parts {
		var err error
		if odd {
			subParts := strings.Split(part, ")")
			if len(subParts) != 2 {
				return nil, false
			}
			y, err = strconv.Atoi(subParts[0])
			if err != nil {
				return nil, false
			}
			list = append(list, Coord{x, y})
		} else {
			subParts := strings.Split(part, "(")
			if len(subParts) != 2 {
				return nil, false
			}
			x, err = strconv.Atoi(subParts[1])
			if err != nil {
				return nil, false
			}
		}
		odd = !odd
	}
	return list, true
}
