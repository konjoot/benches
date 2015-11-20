package main

import (
	// "database/sql"
	// "gopkg.in/pg.v3"
	"time"

	// "errors"
	// "gopkg.in/guregu/null.v3"
	"log"
)

func NewContactQuery(page int, perPage int) *ContactQuery {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 1
	}

	return &ContactQuery{
		limit:      perPage,
		offset:     perPage * (page - 1),
		collection: NewContactList(0),
	}
}

type ContactQuery struct {
	limit      int
	offset     int
	collection ContactList
}

type ContactRow struct {
	Id          *int32
	Email       *string
	FirstName   *string
	LastName    *string
	MiddleName  *string
	DateOfBirth *time.Time
	Sex         *int32
}

type ProfileRow struct {
	Id            *int64
	Type          *string
	UserId        int64
	ClassUnitId   *int64
	ClassUnitName *string
	EnlistedOn    *time.Time
	LeftOn        *time.Time
	SchoolId      *int64
	SchoolName    *string
	SchoolGuid    *string
	SubjectId     *int64
	SubjectName   *string
}

func (cq *ContactQuery) All() []*Contact {
	if !cq.fillUsers() {
		return NewContactList(0).Items()
	}

	// if err := cq.fillDependentData(); err != nil {
	// 	log.Print(err)
	// }

	return cq.collection.Items()
}

func (cq *ContactQuery) fillUsers() (ok bool) {
	var err error

	defer func() {
		if err != nil {
			log.Print(err)
		}

		if cq.collection.Any() {
			ok = true
		}
	}()

	sql := cq.selectUsersStmt()

	conn, err := DbPool()
	if err != nil {
		return
	}
	// defer conn.Close()

	rows, err := conn.Query(sql, cq.limit, cq.offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		contact := NewContact()
		err = rows.Scan(
			&contact.Id,
			&contact.Email,
			&contact.FirstName,
			&contact.LastName,
			&contact.MiddleName,
			&contact.DateOfBirth,
			&contact.Sex,
		)

		if err != nil {
			log.Print(err)
		}

		// contact := &Contact{
		// 	Id:          null.IntFromPtr(row.Id),
		// 	Email:       null.StringFromPtr(row.Email),
		// 	FirstName:   null.StringFromPtr(row.FirstName),
		// 	LastName:    null.StringFromPtr(row.LastName),
		// 	MiddleName:  null.StringFromPtr(row.MiddleName),
		// 	DateOfBirth: null.TimeFromPtr(row.DateOfBirth),
		// 	Sex:         null.IntFromPtr(row.Sex),
		// }

		cq.collection.Append(contact)
	}

	return
}

// func (cq *ContactQuery) fillUsers() (ok bool) {
// 	var err error

// 	defer func() {
// 		if err != nil {
// 			log.Print(err)
// 		}

// 		if cq.collection.Any() {
// 			ok = true
// 		}
// 	}()

// 	ps, err := cq.selectUsersStmt()
// 	if err != nil {
// 		return
// 	}
// 	defer ps.Close()

// 	rows := make([]ContactRow, 0)

// 	_, err = ps.Query(&rows, cq.limit, cq.offset)
// 	if err != nil {
// 		return
// 	}

// 	for _, row := range rows {
// 		contact := &Contact{
// 			Id:          null.IntFromPtr(row.Id),
// 			Email:       null.StringFromPtr(row.Email),
// 			FirstName:   null.StringFromPtr(row.FirstName),
// 			LastName:    null.StringFromPtr(row.LastName),
// 			MiddleName:  null.StringFromPtr(row.MiddleName),
// 			DateOfBirth: null.TimeFromPtr(row.DateOfBirth),
// 			Sex:         null.IntFromPtr(row.Sex),
// 		}
// 		cq.collection.Append(contact)
// 	}
// 	// defer rows.Close()

// 	// for rows.Next() {
// 	// 	contact := NewContact()
// 	// 	rows.Scan(
// 	// 		&contact.Id,
// 	// 		&contact.Email,
// 	// 		&contact.FirstName,
// 	// 		&contact.LastName,
// 	// 		&contact.MiddleName,
// 	// 		&contact.DateOfBirth,
// 	// 		&contact.Sex,
// 	// 	)

// 	// 	cq.collection.Append(contact)
// 	// }

// 	return
// }

// func (cq *ContactQuery) fillDependentData() (err error) {
// 	ps, err := cq.selectDependentDataStmt()
// 	if err != nil {
// 		return
// 	}
// 	defer ps.Close()

// 	rows := make([]ProfileRow, 0)
// 	_, err = ps.Query(&rows, cq.collection.Ids())
// 	if err != nil {
// 		return
// 	}
// 	// defer rows.Close()

// 	// var userId int64

// 	current := cq.collection.Next()

// 	if current == nil {
// 		return errors.New("Empty collection")
// 	}

// 	for _, row := range rows {

// 		profile := &Profile{
// 			Id:   null.IntFromPtr(row.Id),
// 			Type: null.StringFromPtr(row.Type),
// 			ClassUnit: ClassUnit{
// 				Id:         null.IntFromPtr(row.ClassUnitId),
// 				Name:       null.StringFromPtr(row.ClassUnitName),
// 				EnlistedOn: null.TimeFromPtr(row.EnlistedOn),
// 				LeftOn:     null.TimeFromPtr(row.LeftOn),
// 			},
// 			School: School{
// 				Id:   null.IntFromPtr(row.SchoolId),
// 				Name: null.StringFromPtr(row.SchoolName),
// 				Guid: null.StringFromPtr(row.SchoolGuid),
// 			},
// 		}

// 		subject := &Subject{
// 			Id:   null.IntFromPtr(row.SubjectId),
// 			Name: null.StringFromPtr(row.SubjectName),
// 		}
// 		// rows.Scan(
// 		// 	&profile.Id,
// 		// 	&profile.Type,
// 		// 	&userId,
// 		// 	&profile.School.Id,
// 		// 	&profile.School.Name,
// 		// 	&profile.School.Guid,
// 		// 	&profile.ClassUnit.Id,
// 		// 	&profile.ClassUnit.Name,
// 		// 	&profile.ClassUnit.EnlistedOn,
// 		// 	&profile.ClassUnit.LeftOn,
// 		// 	&subject.Id,
// 		// 	&subject.Name,
// 		// )

// 		for current.GetId() != row.UserId {
// 			if next := cq.collection.Next(); next != nil {
// 				current = next
// 			} else {
// 				break
// 			}
// 		}

// 		if current.GetId() != row.UserId {
// 			continue
// 		}

// 		if lastPr := current.LastProfile(); lastPr == nil {
// 			current.Profiles = append(current.Profiles, profile)
// 		} else if lastPr.Id != profile.Id {
// 			current.Profiles = append(current.Profiles, profile)
// 		}

// 		if subject.Id.Valid {
// 			current.LastProfile().Subjects = append(
// 				current.LastProfile().Subjects,
// 				subject,
// 			)
// 		}
// 	}

// 	return
// }

// func (cq *ContactQuery) selectUsersStmt() (*pg.Stmt, error) {
// 	db, err := DBConn()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db.Prepare(`
// 		select	id,
// 		      	email,
// 		      	first_name,
// 		      	last_name,
// 		      	middle_name,
// 		      	date_of_birth,
// 		      	sex
// 		  from users
// 		  where deleted_at is null
// 		  order by id
// 		  limit $1
// 		  offset $2`)
// }

func (cq *ContactQuery) selectUsersStmt() (s string) {
	return `
		select	id,
		      	email,
		      	first_name,
		      	last_name,
		      	middle_name,
		      	date_of_birth,
		      	sex
		  from users
		  where deleted_at is null
		  order by id
		  limit $1
		  offset $2`
}

// func (cq *ContactQuery) selectDependentDataStmt() (*pg.Stmt, error) {
// 	db, err := DBConn()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db.Prepare(`
// 		select	p.id,
// 		      	p.type,
// 		      	p.user_id,
// 		      	p.school_id,
// 		      	s.short_name as school_name,
// 		      	s.guid as school_guid,
// 		      	p.class_unit_id,
// 		      	cu.name as class_unit_name,
// 		      	p.enlisted_on,
// 		      	p.left_on,
// 		      	c.subject_id,
// 		      	sb.name as subject_name
// 		  from profiles p
// 		  left outer join schools s
// 		    on s.id = p.school_id
// 		    and s.deleted_at is null
// 		  left outer join class_units cu
// 		    on cu.id = p.class_unit_id
// 		    and cu.deleted_at is null
// 		  left outer join competences c
// 		    on c.profile_id = p.id
// 		  left outer join subjects sb
// 		    on c.subject_id = sb.id
// 		  where p.deleted_at is null
// 		    and p.user_id = any($1::integer[])
// 		  order by p.user_id, p.id`)
// }

func (cq *ContactQuery) selectDependentDataStmt() (s string) {
	return `
		select	p.id,
		      	p.type,
		      	p.user_id,
		      	p.school_id,
		      	s.short_name as school_name,
		      	s.guid as school_guid,
		      	p.class_unit_id,
		      	cu.name as class_unit_name,
		      	p.enlisted_on,
		      	p.left_on,
		      	c.subject_id,
		      	sb.name as subject_name
		  from profiles p
		  left outer join schools s
		    on s.id = p.school_id
		    and s.deleted_at is null
		  left outer join class_units cu
		    on cu.id = p.class_unit_id
		    and cu.deleted_at is null
		  left outer join competences c
		    on c.profile_id = p.id
		  left outer join subjects sb
		    on c.subject_id = sb.id
		  where p.deleted_at is null
		    and p.user_id = any($1::integer[])
		  order by p.user_id, p.id`
}
