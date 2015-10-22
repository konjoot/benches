package com.example;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class School {
  public Integer id;
  public String name;
  public String guid;

  public School() {
    this.id = null;
    this.name = null;
    this.guid = null;
  }

  public School(
    Integer id,
    String name,
    String guid
  ) {
    this.id = id;
    this.name = name;
    this.guid = guid;
  }
}
