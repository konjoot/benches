package com.example;

import com.example.ClassUnit;
import com.example.Subject;
import com.example.School;
import java.util.ArrayList;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Profile {
  public Integer id;
  public String type;
  public ClassUnit classUnit;
  public School school;
  public ArrayList<Subject> subjects;

  public Profile() {
    this.id = null;
    this.type = null;
    this.classUnit = null;
    this.school = null;
    this.subjects = null;
  }

  public Profile(
    Integer id,
    String type
  ) {
    this.id = id;
    this.type = type;
    this.classUnit = null;
    this.school = null;
    this.subjects = null;
  }
}
