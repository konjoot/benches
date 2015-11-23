package main

import (
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

	conn, err := DBConn()
	if err != nil {
		return
	}

	rows, err := conn.Query(cq.selectUsersStmt(), cq.limit, cq.offset)
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
	conn, err := DBConn()
	if err != nil {
		return
	}

	rows, err := conn.Query(cq.selectDependentDataStmt(), cq.collection.Ids())
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
		school := NewSchool()
		classUnit := NewClassUnit()
		subject := NewSubject()
		rows.Scan(
			&profile.Id,
			&profile.Type,
			&userId,
			&school.Id,
			&school.Name,
			&school.Guid,
			&classUnit.Id,
			&classUnit.Name,
			&classUnit.EnlistedOn,
			&classUnit.LeftOn,
			&subject.Id,
			&subject.Name,
		)

		if school.Id != nil {
			profile.School = school
		}

		if classUnit.Id != nil {
			profile.ClassUnit = classUnit
		}

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

func (cq *ContactQuery) selectUsersStmt() string {
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

func (cq *ContactQuery) selectDependentDataStmt() string {
	return `
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
		  order by p.user_id, p.id`
}
