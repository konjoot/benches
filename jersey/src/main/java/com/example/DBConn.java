package com.example;

import java.sql.*;

public class DBConn {
  // db settings
  public static final String DBHOST = "localhost";
  public static final String DATABASE = "lms2_development";
  public static final String DBUSER = "lms";
  public static final String DBPASS = "";

  public DBConn() {}

  public Connection get() {
    Connection conn = null;

    try
    {
      Class.forName("org.postgresql.Driver");
      String url = "jdbc:postgresql://" + DBHOST + "/" + DATABASE;
      conn = DriverManager.getConnection(url, DBUSER, DBPASS);
    }
    catch (ClassNotFoundException|SQLException e)
    {
      System.err.println(e.getMessage());
    }

    return conn;
  }
}