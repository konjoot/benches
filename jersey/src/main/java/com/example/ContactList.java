package com.example;

import java.util.ArrayList;

public class ContactList extends ArrayList<Contact> {

  public ContactList() {}

  public Integer[] ids() {
    Integer[] ids = new Integer[size()];

    int i = 0;

    for(Contact contact : this)
    {
      ids[i++] = contact.id;
    }

    return ids;
  }
}
