module Main
  class Profile
    attr_reader :subjects

    def initialize(
      id:,
      type:,
      enlisted_on:,
      left_on:,
      school: {},
      class_unit: {}
    )
      @hash = {}
      @hash[:Id] = id if id
      @hash[:Type] = type if type
      @hash[:EnlistedOn] = enlisted_on if enlisted_on
      @hash[:LeftOn] = left_on if left_on

      @hash[:School] = School.new(
        id: school[:id],
        name: school[:name],
        guid: school[:guid]) if school[:id]

      @hash[:ClassUnit] = ClassUnit.new(
        id: class_unit[:id],
        name: class_unit[:name]) if class_unit[:id]

      @subjects = []
    end

    def id
      @hash[:Id]
    end

    def user_id
      @hash[:UserId]
    end

    def to_hash
      @hash[:Subjects] = @subjects if @subjects[0]
      @hash
    end

    def to_json
      Oj.dump(to_hash)
    end
  end
end