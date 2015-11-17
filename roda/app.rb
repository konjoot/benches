require "oj"
require "roda"
require "sequel"

require "./db"
require "./class_unit"
require "./contact"
require "./contact_query"
require "./profile"
require "./school"
require "./subject"

module Main
  class App < Roda
    use Rack::Session::Cookie, secret: "v3rYstr0n9S3cr3t;)"
    Oj.default_options = {:mode => :compat }

    plugin :indifferent_params
    plugin :default_headers,
      "Content-Type" => "application/json; charset=utf-8"

    route do |r|
      # /contacts
      r.get "contacts" do
        Oj.dump ContactQuery.new(params).all()
      end
    end
  end
end
