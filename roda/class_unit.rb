require "oj"

module Main
  class ClassUnit

    def initialize(id:, name:)
      @hash = {}
      @hash[:Id] = id if id
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