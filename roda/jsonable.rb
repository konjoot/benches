require "json"

module Main
  class Jsonable

    def to_json(*json)
      instance_variables.inject({}) do |h, key|
        json_key = key.to_s.gsub('@', '').capitalize
        value = instance_variable_get(key)
        h[json_key] = value if value
        h
      end.to_json(*json)
    end
  end
end