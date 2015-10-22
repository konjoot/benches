package com.example;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class ClassUnit {
  public Integer id;
  public String name;
  public String enlistedOn;
  public String leftOn;

  public ClassUnit() {
    this.id = null;
    this.name = null;
    this.enlistedOn = null;
    this.leftOn = null;
  }

  public ClassUnit(
    Integer id,
    String name,
    String enlistedOn,
    String leftOn
  ) {
    this.id = id;
    this.name = name;
    this.enlistedOn = enlistedOn;
    this.leftOn = leftOn;
  }
}
