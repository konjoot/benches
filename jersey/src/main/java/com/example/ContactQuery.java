package com.example;

import java.sql.*;

import com.example.ClassUnit;
import com.example.Contact;
import com.example.ContactList;
import com.example.Profile;
import com.example.School;
import com.example.Subject;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.NoSuchElementException;


public class ContactQuery {
  private int page;
  private int perPage;
  private ContactList collection;
  private Connection conn;

  public ContactQuery(int page, int perPage) {
    this.page = page <= 0 ? 1 : page ;
    this.perPage = perPage;
    this.conn = new DBConn().get();
    this.collection = new ContactList();
  }

  public ContactList all() {
    if (fillUsers()) { fillDependentData(); }

    return collection;
  }

  private boolean fillUsers() {
    PreparedStatement ps = null;
    ResultSet rs = null;

    try{
      ps = selectUsersStmt();
      rs = ps.executeQuery();

      while ( rs.next() )
      {
        Contact contact = new Contact(
          rs.getString("id") == null ? null : rs.getInt("id"), // I don\t know how to do it better :(
          rs.getString("email"),
          rs.getString("first_name"),
          rs.getString("last_name"),
          rs.getString("middle_name"),
          rs.getString("date_of_birth"),
          rs.getString("sex") == null ? null : rs.getInt("sex") // the same
        );

        collection.add(contact);
      }
    }

    catch (SQLException e)
    {
      System.err.println(e.getMessage());
      return false;
    }
    finally
    {
      close(rs, ps);
    }

    return !collection.isEmpty();
  }

  private void fillDependentData() {
    PreparedStatement ps = null;
    ResultSet rs = null;

    try{
      ps = selectDependentDataStmt();
      rs = ps.executeQuery();

      Iterator<Contact> i = collection.iterator();
      Contact current = i.next();
      int userId;

      while ( rs.next() )
      {
        userId = rs.getInt("user_id");

        while (userId != current.id && i.hasNext())
        {
          current = i.next();
        }

        Profile profile = new Profile(
          rs.getInt("id"),
          rs.getString("type")
        );

        int classUnitId = rs.getInt("class_unit_id");

        if (classUnitId > 0){
          profile.classUnit = new ClassUnit(
            classUnitId,
            rs.getString("name"),
            rs.getString("enlisted_on"),
            rs.getString("left_on")
          );
        }

        int schoolId = rs.getInt("school_id");

        if (schoolId > 0){
          profile.school = new School(
            schoolId,
            rs.getString("short_name"),
            rs.getString("guid")
          );
        }

        int subjectId = rs.getInt("subject_id");

        while (subjectId > 0 && profile.id == rs.getInt("id"))
        {
          Subject subject = new Subject(
            subjectId,
            rs.getString("subject_name")
          );

          if (profile.subjects == null) {
            profile.subjects = new ArrayList<Subject>();
          }

          profile.subjects.add(subject);

          rs.next();
        }

        if (rs.getInt("id") != profile.id){
          rs.previous();
        }

        if (current.profiles == null) {
          current.profiles = new ArrayList<Profile>();
        }

        current.profiles.add(profile);
      }
    }
    catch (Exception e)
    {
      System.err.println(e.getMessage());
    }
    finally
    {
      close(rs, ps);
    }
  }

  private PreparedStatement selectUsersStmt() throws SQLException {
    if (conn == null) { return null; }

    PreparedStatement ps = conn.prepareStatement(
      "select id,"
          + " email,"
          + " first_name,"
          + " last_name,"
          + " middle_name,"
          + " date_of_birth,"
          + " sex"
    + " from users"
    + " where deleted_at is null"
    + " order by id"
    + " limit ?"
    + " offset ?");

    ps.setInt(1, limit());
    ps.setInt(2, offset());

    return ps;
  }

  private PreparedStatement selectDependentDataStmt() throws SQLException {
    if (conn == null) { return null; }

    PreparedStatement ps = conn.prepareStatement(
      "select p.id,"
          + " p.type,"
          + " p.user_id,"
          + " p.class_unit_id,"
          + " p.enlisted_on,"
          + " p.left_on,"
          + " cu.name,"
          + " p.school_id,"
          + " sc.guid,"
          + " sc.short_name,"
          + " c.subject_id,"
          + " sb.name as subject_name"
    + " from profiles p"
    + " left outer join class_units cu"
      + " on cu.id = p.class_unit_id"
      + " and cu.deleted_at is null"
    + " left outer join schools sc"
      + " on sc.id = p.school_id"
      + " and sc.deleted_at is null"
    + " left outer join competences c"
      + " on p.id = c.profile_id"
    + " left outer join subjects sb"
      + " on sb.id = c.subject_id"
    + " where p.deleted_at is null"
      + " and p.user_id = any(?)"
    + " order by p.user_id",
    ResultSet.TYPE_SCROLL_INSENSITIVE,
    ResultSet.CONCUR_READ_ONLY);

    Array array = conn.createArrayOf("integer", collection.ids());

    ps.setArray(1, array);

    return ps;
  }

  private int limit() {
    return perPage;
  }

  private int offset() {
    return perPage * (page - 1);
  }

  private void close(AutoCloseable... c) {
    for (AutoCloseable ac : c)
    {
      try
      {
        ac.close();
      }
      catch (Exception e)
      {
        System.err.println(e.getMessage());
      }
    }
  }
}
