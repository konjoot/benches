package com.example;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.Statement;
import java.sql.ResultSet;
import java.sql.SQLException;

import org.apache.commons.pool2.ObjectPool;
import org.apache.commons.pool2.impl.GenericObjectPool;
import org.apache.commons.dbcp2.ConnectionFactory;
import org.apache.commons.dbcp2.PoolableConnection;
import org.apache.commons.dbcp2.PoolingDataSource;
import org.apache.commons.dbcp2.PoolableConnectionFactory;
import org.apache.commons.dbcp2.DriverManagerConnectionFactory;

public class DBConn {
  // db settings
  public static final String DBHOST = "localhost";
  public static final String DATABASE = "lms2_development";
  public static final String DBUSER = "lms";
  public static final String DBPASS = "";

  public static DataSource dataSource;

  public static Connection get() {
    if (dataSource == null) {
      try {

        Class.forName("org.postgresql.Driver");
        dataSource = setupDataSource();

      } catch (ClassNotFoundException e) {

        System.err.println(e.getMessage());

      }
    }

    Connection conn = null;

    try {

      if (dataSource != null) {
        conn = dataSource.getConnection();
      }

    } catch(SQLException e) {

      System.err.println(e.getMessage());

    }

    return conn;
  }

  private static DataSource setupDataSource() {
    String connectURI = "jdbc:postgresql://"+DBHOST+"/"+DATABASE+"?user="+DBUSER+"&password="+DBPASS+"&ssl=false";

    ConnectionFactory connectionFactory =
      new DriverManagerConnectionFactory(connectURI, null);

    PoolableConnectionFactory poolableConnectionFactory =
      new PoolableConnectionFactory(connectionFactory, null);

    ObjectPool<PoolableConnection> connectionPool =
      new GenericObjectPool<>(poolableConnectionFactory);

    poolableConnectionFactory.setPool(connectionPool);

    return new PoolingDataSource<>(connectionPool);
  }
}
