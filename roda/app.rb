require "roda"

require "./contact_query"

module Main
  class App < Roda
    use Rack::Session::Cookie, secret: "v3rYstr0n9S3cr3t;)"

    plugin :indifferent_params
    plugin :default_headers,
      'Content-Type' => 'application/json; charset=utf-8'

    route do |r|
      # /contacts
      r.get "contacts" do
        ContactQuery.new(params).all().to_json
      end
    end
  end
end
