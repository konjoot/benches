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

  public Profile() {}

  public Profile(
    String id,
    String type,
    String classUnitId,
    String classUnitName,
    String enlistedOn,
    String leftOn,
    String schoolId,
    String schoolName,
    String schoolGuid
  ) {
    this.id = id == null ? null : new Integer(id);
    this.type = type;
    this.subjects = null;

    ClassUnit cu = new ClassUnit(
      classUnitId,
      classUnitName,
      enlistedOn,
      leftOn
    );

    School s = new School(
      schoolId,
      schoolName,
      schoolGuid
    );

    this.classUnit = cu.id == null ? null : cu;
    this.school = s.id == null ? null : s;
  }

  public void addSubject(Subject subject) {
    if (subjects == null) {
      subjects = new ArrayList<Subject>();
    }

    subjects.add(subject);
  }
}
