module Main
  class School < Jsonable
    attr_reader :id, :guid, :name

    def initialize(id:, guid:, name:)
      @id = id
      @guid = guid
      @name = name
    end
  end
end
