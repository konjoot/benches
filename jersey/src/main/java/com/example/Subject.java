package com.example;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Subject {
  public Integer id;
  public String name;

  public Subject() {}

  public Subject(
    String id,
    String name
  ) {
    this.id = id == null ? null : new Integer(id);
    this.name = name;
  }
}
