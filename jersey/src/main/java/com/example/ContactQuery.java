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
  private int limit;
  private int offset;
  private ContactList collection;
  private Connection conn;

  public ContactQuery(int page, int perPage) {
    page = page < 1 ? 1 : page;
    perPage = perPage < 1 ? 1 : perPage;

    this.limit = perPage;
    this.offset = perPage * (page - 1);
    this.conn = null;
    this.collection = new ContactList();
  }

  public ContactList all() {
    this.conn = new DBConn().get();

    if (fillUsers()) { fillDependentData(); }

    close(conn);

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
          rs.getString("id"),
          rs.getString("email"),
          rs.getString("first_name"),
          rs.getString("last_name"),
          rs.getString("middle_name"),
          rs.getString("date_of_birth"),
          rs.getString("sex")
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

        if (userId != current.id){
          continue;
        }

        Profile profile = new Profile(
          rs.getString("id"),
          rs.getString("type"),
          rs.getString("class_unit_id"),
          rs.getString("class_unit_name"),
          rs.getString("enlisted_on"),
          rs.getString("left_on"),
          rs.getString("school_id"),
          rs.getString("school_name"),
          rs.getString("school_guid")
        );

        if (profile.id == null) { continue; }

        Profile lastPr = current.lastProfile();

        if (lastPr == null) {
          current.addProfile(profile);
        } else if (!lastPr.id.equals(profile.id)) {
          current.addProfile(profile);
        }

        Subject subject = new Subject(
          rs.getString("subject_id"),
          rs.getString("subject_name")
        );

        if (subject.id != null) {
          current.lastProfile().addSubject(subject);
        }
      }
    }
    catch (Exception e)
    {
      e.printStackTrace();
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

    ps.setInt(1, limit);
    ps.setInt(2, offset);

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
          + " cu.name as class_unit_name,"
          + " p.school_id,"
          + " sc.guid as school_guid,"
          + " sc.short_name as school_name,"
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
    + " order by p.user_id, p.id");

    Array array = conn.createArrayOf("integer", collection.ids());

    ps.setArray(1, array);

    return ps;
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
