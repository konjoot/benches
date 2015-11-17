require "oj"

module Main
  class School

    def initialize(id:, guid:, name:)
      @hash = {}
      @hash[:Id] = id if id
      @hash[:Guid] = guid if guid
      @hash[:Name] = name if name
    end

    def to_hash
      @hash
    end

    def to_json
      Oj.dump(@hash)
    end
  end
end
