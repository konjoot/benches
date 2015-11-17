module Main
  class Contact
    attr_reader :profiles

    def initialize(
      id:,
      email:,
      first_name:,
      last_name:,
      middle_name:,
      date_of_birth:,
      sex:
    )
      @hash = {}
      @hash[:Id] = id if id
      @hash[:Email] = email if email
      @hash[:FirstName] = first_name if first_name
      @hash[:LastName] = last_name if last_name
      @hash[:MiddleName] = middle_name if middle_name
      @hash[:DateOfBirth] = date_of_birth if date_of_birth
      @hash[:Sex] = sex if sex
      @profiles = []
    end

    def id
      @hash[:Id]
    end

    def to_hash
      @hash[:Profiles] = @profiles if @profiles[0]
      @hash
    end

    def to_json
      Oj.dump(to_hash)
    end
  end
end