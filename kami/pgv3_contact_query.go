package main

// import (
// 	"errors"
// 	"log"
// 	// "time"
// )

// const (
// 	PG_USERS_QUERY = `
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
// 		  limit ?
// 		  offset ?;`

// 	PG_PROFILES_QUERY = `
// 		select	p.id as profile__id,
// 		      	p.type as profile__type,
// 		      	p.user_id as profile__user_id,
// 		      	p.school_id as school__id,
// 		      	s.short_name as school__name,
// 		      	s.guid as school__guid,
// 		      	p.class_unit_id as class_unit__id,
// 		      	cu.name as class_unit__name,
// 		      	p.enlisted_on as class_unit__enlisted_on,
// 		      	p.left_on as class_unit__left_on,
// 		      	c.subject_id as subject__id,
// 		      	sb.name as subject__name
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
// 		    and p.user_id = any(?::integer[])
// 		  order by p.user_id, p.id;`
// )

// func NewPGContactQuery(page int, perPage int) *PgContactQuery {
// 	if page < 1 {
// 		page = 1
// 	}
// 	if perPage < 1 {
// 		perPage = 1
// 	}

// 	return &PgContactQuery{
// 		limit:      perPage,
// 		offset:     perPage * (page - 1),
// 		collection: NewContactList(perPage),
// 	}
// }

// type PgContactQuery struct {
// 	limit      int
// 	offset     int
// 	collection ContactList
// }

// func (cq *PgContactQuery) All() []*Contact {
// 	if !cq.fillUsers() {
// 		return NewContactList(0).Items
// 	}

// 	if err := cq.fillDependentData(); err != nil {
// 		log.Print(err)
// 	}

// 	return cq.collection.Items
// }

// func (cq *PgContactQuery) fillUsers() (ok bool) {
// 	var err error

// 	defer func() {
// 		if err != nil {
// 			log.Print(err)
// 		}

// 		if cq.collection.Any() {
// 			ok = true
// 		}
// 	}()

// 	db, err := PgDBConn()
// 	if err != nil {
// 		return
// 	}

// 	_, err = db.Query(&cq.collection.Items, PG_USERS_QUERY, cq.limit, cq.offset)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func (cq *PgContactQuery) fillDependentData() (err error) {
// 	db, err := PgDBConn()
// 	if err != nil {
// 		return
// 	}

// 	rows := make([]*ProfileRow, 0, len(cq.collection.Items))

// 	_, err = db.Query(&rows, PG_PROFILES_QUERY, cq.collection.IntIds())
// 	if err != nil {
// 		return
// 	}

// 	current := cq.collection.Next()

// 	if current == nil {
// 		return errors.New("Empty collection")
// 	}

// 	for _, row := range rows {

// 		for *current.Id != *row.Profile.UserId {
// 			if next := cq.collection.Next(); next != nil {
// 				current = next
// 			} else {
// 				break
// 			}
// 		}

// 		if *current.Id != *row.Profile.UserId {
// 			continue
// 		}

// 		if row.ClassUnit.Id != nil {
// 			row.Profile.ClassUnit = row.ClassUnit
// 		}

// 		if row.School.Id != nil {
// 			row.Profile.School = row.School
// 		}

// 		if lastPr := current.LastProfile(); lastPr == nil {
// 			current.Profiles = append(current.Profiles, row.Profile)
// 		} else if *lastPr.Id != *row.Profile.Id {
// 			current.Profiles = append(current.Profiles, row.Profile)
// 		}

// 		if row.Subject.Id != nil {
// 			current.LastProfile().Subjects = append(
// 				current.LastProfile().Subjects,
// 				row.Subject,
// 			)
// 		}
// 	}

// 	return
// }
