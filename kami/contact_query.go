package main

import (
	"database/sql"
	"errors"
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
		collection: NewContactList(perPage),
	}
}

type ContactQuery struct {
	limit      int
	offset     int
	collection ContactList
	conn       *sql.DB
}

func (cq *ContactQuery) All() []*Contact {
	if !cq.fillUsers() {
		return NewContactList(0).Items()
	}

	if err := cq.fillDependentData(); err != nil {
		log.Print(err)
	}

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

	ps, err := cq.selectUsersStmt()
	if err != nil {
		return
	}
	defer ps.Close()

	rows, err := ps.Query(cq.limit, cq.offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		contact := NewContact()
		rows.Scan(
			&contact.Id,
			&contact.Email,
			&contact.FirstName,
			&contact.LastName,
			&contact.MiddleName,
			&contact.DateOfBirth,
			&contact.Sex,
		)

		cq.collection.Append(contact)
	}

	return
}

func (cq *ContactQuery) fillDependentData() (err error) {
	ps, err := cq.selectDependentDataStmt()
	if err != nil {
		return
	}
	defer ps.Close()

	rows, err := ps.Query(cq.collection.Ids())
	if err != nil {
		return
	}
	defer rows.Close()

	var userId *int32

	current := cq.collection.Next()

	if current == nil {
		return errors.New("Empty collection")
	}

	for rows.Next() {

		profile := NewProfile()
		subject := NewSubject()
		rows.Scan(
			&profile.Id,
			&profile.Type,
			&userId,
			&profile.School.Id,
			&profile.School.Name,
			&profile.School.Guid,
			&profile.ClassUnit.Id,
			&profile.ClassUnit.Name,
			&profile.ClassUnit.EnlistedOn,
			&profile.ClassUnit.LeftOn,
			&subject.Id,
			&subject.Name,
		)

		for *current.Id != *userId {
			if next := cq.collection.Next(); next != nil {
				current = next
			} else {
				break
			}
		}

		if *current.Id != *userId {
			continue
		}

		if lastPr := current.LastProfile(); lastPr == nil {
			current.Profiles = append(current.Profiles, profile)
		} else if *lastPr.Id != *profile.Id {
			current.Profiles = append(current.Profiles, profile)
		}

		if subject.Id != nil {
			current.LastProfile().Subjects = append(
				current.LastProfile().Subjects,
				subject,
			)
		}
	}

	return
}

func (cq *ContactQuery) selectUsersStmt() (*sql.Stmt, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	return db.Prepare(`
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
		  offset $2`)
}

func (cq *ContactQuery) selectDependentDataStmt() (*sql.Stmt, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	return db.Prepare(`
		select	p.id,
		      	p.type,
		      	p.user_id,
		      	p.school_id,
		      	s.short_name,
		      	s.guid,
		      	p.class_unit_id,
		      	cu.name,
		      	p.enlisted_on,
		      	p.left_on,
		      	c.subject_id,
		      	sb.name
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
		  order by p.user_id, p.id`)
}
