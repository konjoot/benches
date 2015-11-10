package com.example;

import org.eclipse.jetty.server.Server;
import org.glassfish.jersey.jetty.JettyHttpContainerFactory;
import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.moxy.json.MoxyJsonConfig;

import javax.ws.rs.ext.ContextResolver;

import java.io.IOException;
import java.net.URI;
import java.util.HashMap;
import java.util.Map;

/**
 * Main class.
 *
 */
public class Main {
  // Base URI the Grizzly HTTP server will listen on
  public static final String BASE_URI = "http://localhost:8080/";

  /**
   * Starts Jetty HTTP server exposing JAX-RS resources defined in this application.
   * @return Jetty HTTP server.
   */
  public static Server startServer() {
    // create a resource config that scans for JAX-RS resources and providers
    // in com.example package
    final ResourceConfig rc =
      new ResourceConfig()
        .packages("com.example")
        .packages("org.glassfish.jersey.examples.jsonmoxy")
        .register(createMoxyJsonResolver());

    // create and start a new instance of grizzly http server
    // exposing the Jersey application at BASE_URI
    return JettyHttpContainerFactory.createServer(URI.create(BASE_URI), rc);
  }

  /**
   * Main method.
   * @param args
   * @throws IOException
   */
  public static void main(String[] args) throws Exception {
    try {
      final Server server = startServer();
      System.out.println(String.format("Jersey app started with WADL available at "
            + "%sapplication.wadl\nHit enter to stop it...", BASE_URI));
      System.in.read();
      server.stop();
    } catch (Exception e) {
      System.err.println(e.getMessage());
    }
  }

  public static ContextResolver<MoxyJsonConfig> createMoxyJsonResolver() {
    final MoxyJsonConfig moxyJsonConfig = new MoxyJsonConfig();
    Map<String, String> namespacePrefixMapper = new HashMap<String, String>(1);
    namespacePrefixMapper.put("http://www.w3.org/2001/XMLSchema-instance", "xsi");
    moxyJsonConfig.setNamespacePrefixMapper(namespacePrefixMapper).setNamespaceSeparator(':');
    return moxyJsonConfig.resolver();
  }
}

