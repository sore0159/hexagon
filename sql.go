package hexagon

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type NullCoord struct {
	Coord Coord `json:"coord"`
	Valid bool  `json:"valid"`
}

func NewNullCoord() *NullCoord {
	return &NullCoord{}
}

func (nc NullCoord) IsCoord(c Coord) bool {
	return nc.Valid && nc.Coord == c
}
func (nc NullCoord) Eq(nc2 NullCoord) bool {
	return nc.Valid && nc2.Valid && nc.Coord == nc2.Coord
}

func (nc NullCoord) Value() (driver.Value, error) {
	if !nc.Valid {
		return nil, nil
	}
	return nc.Coord.Value()
}
func (nc *NullCoord) Scan(value interface{}) error {
	if value == nil {
		nc.Valid, nc.Coord = false, Coord{0, 0}
		return nil
	}
	nc.Valid = true
	return nc.Coord.Scan(value)
}

// Scan impliments Scanner for use in SQL queries
func (c *Coord) Scan(value interface{}) error {
	if value == nil {
		return errors.New("Bad value scanned to hexagon coord: NULL")
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Bad value scanned to hexagon coord:", value))
	}
	parts := strings.Split(string(bytes), ",")
	if len(parts) < 2 || len(parts[0]) < 2 || len(parts[1]) < 2 {
		return errors.New(fmt.Sprint("bad hexagon coord scan value:", parts, "TOO SHORT"))
	}
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

type CoordList []Coord

func (cl CoordList) Value() (driver.Value, error) {
	if len(cl) == 0 {
		return "{}", nil
	}
	parts := make([]string, len(cl))
	for i, c := range cl {
		parts[i] = fmt.Sprintf("\"(%d,%d)\"", c[0], c[1])
	}
	str := fmt.Sprintf("{%s}", strings.Join(parts, ", "))
	return str, nil
}

func (cl *CoordList) Scan(src interface{}) error {
	if src == nil {
		return errors.New("Bad value scanned to hexagon coordlist: NULL")
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Bad value scanned to hexagon coordlist: %v", src)
	}
	str := string(bytes)
	if str == "{}" {
		*cl = []Coord{}
		return nil
	}
	//str = strings.Trim(str, "{}")
	parts := strings.Split(str, "(")[1:]
	list := make([]Coord, len(parts))
	for i, part := range parts {
		pt2 := strings.Split(part, ")")[0]
		err := (&(list[i])).Scan([]byte(fmt.Sprintf("(%s)", pt2)))
		if err != nil {
			return err
		}
	}
	*cl = list
	return nil
}
