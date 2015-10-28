package com.example;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.HeaderParam;
import javax.ws.rs.QueryParam;
import javax.ws.rs.DefaultValue;

import java.util.List;

import com.example.ContactQuery;


@Path("contacts")
public class ContactResource {

  @GET
  @Produces(MediaType.APPLICATION_JSON + ";charset=utf-8")
  public List<Contact> getCollection(
    @DefaultValue("100") @QueryParam("per_page") int perPage,
    @DefaultValue("1") @QueryParam("page") int page
  ) {
    return new ContactQuery(page, perPage).all();
  }
}
