require "./jsonable"

module Main
  class Contact < Jsonable
    attr_reader :id,
                :email,
                :first_name,
                :last_name,
                :middle_name,
                :date_of_birth,
                :sex

    def initialize(attrs)
      attrs.each do |key, val|
        instance_variable_set("@#{key}", val)
      end
    end
  end
end