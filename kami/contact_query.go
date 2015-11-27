package main

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	USERS_QUERY = `
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
		  limit ?
		  offset ?;`

	PROFILES_QUERY = `
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
		    and p.user_id in (?)
		  order by p.user_id, p.id;`
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
		return NewContactList(0).Items
	}

	if err := cq.fillDependentData(); err != nil {
		log.Print(err)
	}

	return cq.collection.Items
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

	db, err := DBConn()
	if err != nil {
		return
	}

	query := db.Rebind(USERS_QUERY)

	rows, err := db.Queryx(query, cq.limit, cq.offset)
	if err != nil {
		return
	}
	defer rows.Close()

	var contact *Contact
	for rows.Next() {
		contact = NewContact()
		rows.Scan(
			&contact.Id,
			&contact.Email,
			&contact.FirstName,
			&contact.LastName,
			&contact.MiddleName,
			&contact.DateOfBirth,
			&contact.Sex,
		)
		cq.collection.Items = append(cq.collection.Items, contact)
	}

	return
}

func (cq *ContactQuery) fillDependentData() (err error) {
	db, err := DBConn()
	if err != nil {
		return
	}

	query, args, err := sqlx.In(PROFILES_QUERY, cq.collection.Ids())
	if err != nil {
		return
	}

	query = db.Rebind(query)

	rows, err := db.Queryx(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	current := cq.collection.Next()

	if current == nil {
		return errors.New("Empty collection")
	}

	var (
		profile   *Profile
		classUnit *ClassUnit
		school    *School
		subject   *Subject
	)

	for rows.Next() {
		profile = NewProfile()
		classUnit = NewClassUnit()
		school = NewSchool()
		subject = NewSubject()

		rows.Scan(
			&profile.Id,
			&profile.Type,
			&profile.UserId,
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

		for *current.Id != *profile.UserId {
			if next := cq.collection.Next(); next != nil {
				current = next
			} else {
				break
			}
		}

		if *current.Id != *profile.UserId {
			continue
		}

		if classUnit.Id != nil {
			profile.ClassUnit = classUnit
		}

		if school.Id != nil {
			profile.School = school
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
