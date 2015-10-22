package com.example;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Subject {
  public Integer id;
  public String name;

  public Subject() {
    this.id = null;
    this.name = null;
  }

  public Subject(
    Integer id,
    String name
  ) {
    this.id = id;
    this.name = name;
  }
}
