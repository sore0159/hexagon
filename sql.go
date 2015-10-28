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
	return fmt.Sprintf("(%d,%d)", c[0], c[1]), nil
}
