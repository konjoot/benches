package com.example;

import java.util.Date;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Contact {
  public Integer id;
  public String email;
  public String firstName;
  public String lastName;
  public String middleName;
  public String dateOfBirth;
  public Integer sex;

  public Contact() {
    this.id = null;
    this.email = null;
    this.firstName = null;
    this.lastName = null;
    this.middleName = null;
    this.dateOfBirth = null;
    this.sex = null;
  }

  public Contact(
    Integer id,
    String email,
    String firstName,
    String lastName,
    String middleName,
    String dateOfBirth,
    Integer sex
  ) {
    this.id = id;
    this.email = email;
    this.firstName = firstName;
    this.lastName = lastName;
    this.middleName = middleName;
    this.dateOfBirth = dateOfBirth;
    this.sex = sex;
  }
}
