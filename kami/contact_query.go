package main

import (
	// "errors"
	"database/sql"
	// "encoding/binary"
	"log"
	// "time"
	// "bytes"
	// "strconv"

	// "github.com/jmoiron/sqlx"
)

const (
	USERS_QUERY = `
		select	id,
		      	email,
		      	first_name,
		      	last_name,
		      	middle_name,
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

	// if err := cq.fillDependentData(); err != nil {
	// 	log.Print(err)
	// }

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
	var row = make([]sql.RawBytes, 6)
	// var id int32
	// var email string
	// var firstName string
	// var lastName string
	// var middleName string
	// var dateOfBirth time.Time
	// var sex int32
	for rows.Next() {
		contact = NewContact()
		// 	// rows.MapScan(row)
		// 	// row = contactPool.Get().([]interface{})
		err = rows.Scan(
			&row[0],
			&row[1],
			&row[2],
			&row[3],
			&row[4],
			&row[5],
		)
		if err != nil {
			log.Print("this err")
			return
		}
		if row[0] != nil {
			contact.Id.Set([]byte(row[0]))
		}
		if row[1] != nil {
			contact.Email.Set([]byte(row[1]))
		}
		if row[2] != nil {
			contact.FirstName.Set([]byte(row[2]))
		}
		if row[3] != nil {
			contact.LastName.Set([]byte(row[3]))
		}
		if row[4] != nil {
			contact.MiddleName.Set([]byte(row[4]))
		}
		// if dateOfBirth, ok := row[5].(time.Time); ok {
		// 	contact.DateOfBirth = &dateOfBirth
		// }
		if row[5] != nil {
			contact.Sex.Set([]byte(row[5]))
		}
		// 	// rows.Scan(
		// 	// 	&contact.Id,
		// 	// 	&contact.Email,
		// 	// 	&contact.FirstName,
		// 	// 	&contact.LastName,
		// 	// 	&contact.MiddleName,
		// 	// 	&contact.DateOfBirth,
		// 	// 	&contact.Sex,
		// 	// )
		cq.collection.Items = append(cq.collection.Items, contact)
		// cq.collection.Append(*contact)
	}

	return
}

// func (cq *ContactQuery) fillDependentData() (err error) {
// 	db, err := DBConn()
// 	if err != nil {
// 		return
// 	}

// 	query, args, err := sqlx.In(PROFILES_QUERY, cq.collection.Ids())
// 	if err != nil {
// 		return
// 	}

// 	query = db.Rebind(query)

// 	rows, err := db.Queryx(query, args...)
// 	if err != nil {
// 		return
// 	}
// 	defer rows.Close()

// 	current := cq.collection.Next()

// 	if current == nil {
// 		return errors.New("Empty collection")
// 	}

// 	var (
// 		profile   *Profile
// 		classUnit *ClassUnit
// 		school    *School
// 		subject   *Subject
// 	)

// 	for rows.Next() {
// 		profile = NewProfile()
// 		classUnit = NewClassUnit()
// 		school = NewSchool()
// 		subject = NewSubject()

// 		rows.Scan(
// 			&profile.Id,
// 			&profile.Type,
// 			&profile.UserId,
// 			&school.Id,
// 			&school.Name,
// 			&school.Guid,
// 			&classUnit.Id,
// 			&classUnit.Name,
// 			&classUnit.EnlistedOn,
// 			&classUnit.LeftOn,
// 			&subject.Id,
// 			&subject.Name,
// 		)

// 		for *current.Id != *profile.UserId {
// 			if next := cq.collection.Next(); next != nil {
// 				current = next
// 			} else {
// 				break
// 			}
// 		}

// 		if *current.Id != *profile.UserId {
// 			continue
// 		}

// 		if classUnit.Id != nil {
// 			profile.ClassUnit = classUnit
// 		}

// 		if school.Id != nil {
// 			profile.School = school
// 		}

// 		if lastPr := current.LastProfile(); lastPr == nil {
// 			current.Profiles = append(current.Profiles, profile)
// 		} else if *lastPr.Id != *profile.Id {
// 			current.Profiles = append(current.Profiles, profile)
// 		}

// 		if subject.Id != nil {
// 			current.LastProfile().Subjects = append(
// 				current.LastProfile().Subjects,
// 				subject,
// 			)
// 		}
// 	}

// 	return
// }
