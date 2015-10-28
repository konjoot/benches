require "roda"

class App < Roda
  use Rack::Session::Cookie, secret: "v3rYstr0n9S3cr3t;)"

  route do |r|
    # /contacts
    r.get "contacts" do
      "There will be contacts"
    end
  end
end
