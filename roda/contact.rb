require "./jsonable"

module Main
  class Contact < Jsonable
    attr_reader :id,
                :email,
                :first_name,
                :last_name,
                :middle_name,
                :date_of_birth,
                :sex,
                :profiles

    def initialize(
      id:,
      email:,
      first_name:,
      last_name:,
      middle_name:,
      date_of_birth:,
      sex:
    )
      @id = id
      @email = email
      @first_name = first_name
      @last_name = last_name
      @middle_name = middle_name
      @date_of_birth = date_of_birth
      @sex = sex
      @profiles = []
    end
  end
end