package com.example;

import com.example.Contact;
import java.sql.*;
import java.util.ArrayList;

public class ContactQuery {
  private int page;
  private int perPage;
  private ArrayList<Contact> collection;
  private Connection conn;

  public ContactQuery(int page, int perPage) {
    this.page = page <= 0 ? 1 : page ;
    this.perPage = perPage;
    this.conn = new DBConn().get();
    this.collection = new ArrayList<Contact>();
  }

  public ArrayList<Contact> all() {
    PreparedStatement ps = null;
    ResultSet rs = null;

    try{
      ps = prepareStatement();
      rs = ps.executeQuery();

      while ( rs.next() )
      {
        Contact contact = new Contact(
          rs.getString("id") == null ? null : rs.getInt("id"), // I don\t know how do it better :(
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
    }
    finally
    {
      close(ps);
      close(rs);
    }

    return collection;
  }

  private PreparedStatement prepareStatement() {
    if (conn == null) { return null; }

    PreparedStatement ps = null;

    try {
      ps = conn.prepareStatement(
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

    }
    catch (SQLException e) {
      System.err.println(e.getMessage());
    }

    return ps;
  }

  private int limit() {
    return perPage;
  }

  private int offset() {
    return perPage * (page - 1);
  }

  private void close(AutoCloseable c) {
    try
    {
      c.close();
    }
    catch (Exception e)
    {
      System.err.println(e.getMessage());
    }
  }
}
