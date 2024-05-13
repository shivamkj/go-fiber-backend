package attendance

import (
	"fmt"

	gofiber "github.com/gofiber/fiber/v2"
	. "github.com/qnify/api-server/utils/fiber"
)

type total struct {
	Absent  int `json:"absent"`
	HalfDay int `json:"halfDay"`
	Late    int `json:"late"`
}

func (m *module) getStudentTotal(c *gofiber.Ctx) error {
	id := c.Params("studentId")
	fmt.Println(id)

	const _qry = `
SELECT
 	COUNT(is_absent) AS absent,
  COUNT(is_half_day) AS half_day,
  COUNT(is_late) AS late
FROM attendance
WHERE student_id = $1;  
`
	row := m.db.QueryRowX(_qry, id)
	s := total{}
	if err := row.Scan(&s.Absent, &s.HalfDay, &s.Late); err != nil {
		return err
	}

	return SendResponse(c, s)
}

func (m *module) getSectionSummary(c *gofiber.Ctx) error {
	filter, err := ParseQueryFilter(c)
	if err != nil {
		return err
	}

	const _qry = `
SELECT
  COUNT(is_absent) AS absent,
  COUNT(is_half_day) AS half_day,
  COUNT(is_late) AS late
FROM attendance
WHERE student_id = $1;  
`
	rows, err := m.db.Listx(_qry, filter)
	if err != nil {
		return err
	}
	defer rows.Close()

	var courses []total
	for rows.Next() {
		s := total{}
		if err := rows.Scan(&s.Absent, &s.HalfDay, &s.Late); err != nil {
			return err
		}
		courses = append(courses, s)
	}

	return SendResponse(c, courses)
}
