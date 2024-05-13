package course

import (
	"github.com/qnify/api-server/utils/fiber"
)

const (
	_upsertColumns = "id, cid, name, locale, validity, price, discount_percent, is_public," +
		"is_open, description, thumbnail, starts_at, ends_at, created_at, updated_at"
	// _selectColumns  = "id," + _upsertColumns
	_courseByIdPg   = "SELECT " + _upsertColumns + " FROM course WHERE id=$1"
	_courseByIdMs   = "SELECT " + _upsertColumns + " FROM course WHERE id=?"
	_listCourse     = "SELECT " + _upsertColumns + " FROM course"
	_insertCourse   = "INSERT INTO course (" + _upsertColumns + ") "
	_insertCoursePg = _insertCourse + "VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id"
	_insertCourseMs = _insertCourse + "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_updateCoursePg = "UPDATE course SET name=$1, locale=$2, validity=$3, price=$4, discount_percent=$5, " +
		"is_public=$6, is_open=$7, description=$8, thumbnail=$9, starts_at=$10, ends_at=$11, created_at=$12, updated_at=$13 WHERE id=$14"
	_updateCourseMs = "UPDATE course SET name=?, locale=?, validity=?, price=?, discount_percent=?, " +
		"is_public=?, is_open=?, description=?, thumbnail=?, starts_at=?, ends_at=?, created_at=?, updated_at=? WHERE id=?"
	_deleteCoursePg = "DELETE FROM course WHERE id=$1"
	_deleteCourseMs = "DELETE FROM course WHERE id=?"
)

func (m *courseModule) courseById(id string) (*Course, error) {
	row := m.db.QueryRowX(_courseByIdPg, _courseByIdMs, id)
	c := Course{}
	err := row.Scan(
		&c.ID, &c.CID, &c.Name, &c.Locale, &c.Validity, &c.Price, &c.DiscountPercent, &c.IsPublic,
		&c.IsOpen, &c.Description, &c.Thumbnail, &c.StartsAt, &c.EndsAt, &c.CreatedAt, &c.UpdatedAt,
	)
	return &c, err
}

func (m *courseModule) listCourse(filter *fiber.QueryFilters) (*[]Course, error) {
	rows, err := m.db.Listx(_listCourse, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		c := Course{}
		if err := rows.Scan(
			&c.ID, &c.CID, &c.Name, &c.Locale, &c.Validity, &c.Price, &c.DiscountPercent, &c.IsPublic,
			&c.IsOpen, &c.Description, &c.Thumbnail, &c.StartsAt, &c.EndsAt, &c.CreatedAt, &c.UpdatedAt,
		); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	return &courses, nil
}

func (m *courseModule) newCourse(c *Course) (int, error) {
	return m.db.InsertX(_insertCoursePg, _insertCourseMs,
		c.ID, c.CID, c.Name, c.Locale, c.Validity, c.Price, c.DiscountPercent, c.IsPublic, c.IsOpen,
		c.Description, c.Thumbnail, c.StartsAt, c.EndsAt, c.CreatedAt, c.UpdatedAt,
	)
}

func (m *courseModule) setCourse(id string, c *Course) error {
	_, err := m.db.ExecX(_updateCoursePg, _updateCourseMs,
		c.Name, c.Locale, c.Validity, c.Price, c.DiscountPercent, c.IsPublic, c.IsOpen,
		c.Description, c.Thumbnail, c.StartsAt, c.EndsAt, c.CreatedAt, c.UpdatedAt, id,
	)
	return err
}

func (m *courseModule) delCourse(id string) error {
	_, err := m.db.ExecX(_deleteCoursePg, _deleteCourseMs, id)
	return err
}
