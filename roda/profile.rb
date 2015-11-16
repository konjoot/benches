require './school'
require './class_unit'

module Main
  class Profile < Jsonable
    attr_reader :id,
                :user_id,
                :type,
                :enlisted_on,
                :left_on,
                :school,
                :class_unit,
                :subjects

    def initialize(
      id:,
      user_id:,
      type:,
      enlisted_on:,
      left_on:,
      school: {},
      class_unit: {}
    )
      @id = id
      @user_id = user_id
      @type = type
      @enlisted_on = enlisted_on
      @left_on = left_on

      @school = School.new(
        id: school[:id],
        name: school[:name],
        guid: school[:guid]) if school

      @class_unit = ClassUnit.new(
        id: class_unit[:id],
        name: class_unit[:name]
      ) if class_unit

      @subjects = []
    end
  end
end