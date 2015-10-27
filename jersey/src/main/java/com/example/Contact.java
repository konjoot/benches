package com.example;

import java.util.ArrayList;
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
  public ArrayList<Profile> profiles;

  public Contact() {
    this.id = null;
    this.email = null;
    this.firstName = null;
    this.lastName = null;
    this.middleName = null;
    this.dateOfBirth = null;
    this.sex = null;
    this.profiles = null;
  }

  public Contact(
    String id,
    String email,
    String firstName,
    String lastName,
    String middleName,
    String dateOfBirth,
    String sex
  ) {
    this.id = id == null ? null : new Integer(id);
    this.email = email;
    this.firstName = firstName;
    this.lastName = lastName;
    this.middleName = middleName;
    this.dateOfBirth = dateOfBirth;
    this.sex = sex == null ? null : new Integer(sex);
    this.profiles = null;
  }

  public Profile lastProfile() {
    if (profiles == null) { return null; }

    int size = profiles.size();

    if (size == 0) { return null; }

    return profiles.get(size - 1);
  }

  public void addProfile(Profile profile) {
    if (profiles == null) {
      profiles = new ArrayList<Profile>();
    }

    profiles.add(profile);
  }
}
