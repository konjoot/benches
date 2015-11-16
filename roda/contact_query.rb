require "sequel"

require "./db"
require "./contact"
require "./profile"
require "./subject"

module Main
  class ContactQuery
    include DB

    attr_reader :limit,
                :offset,
                :collection

    def initialize(params)
      page = get(:page, from: params, default: 1)
      per_page = get(:per_page, from: params, default: 100)

      @limit = per_page
      @offset = per_page * (page - 1)
      @collection = []
    end

    def all
      if fill_users
        fill_dependent_data
      end

      collection
    end

    private

    def fill_users
      return if db.nil?

      @collection = db[:users].select(
        :id,
        :email,
        :first_name,
        :last_name,
        :middle_name,
        :date_of_birth,
        :sex
      ).order(:id)
      .limit(limit)
      .offset(offset).map do |rec|
        Contact.new(rec)
      end

      collection.any?
    end

    def fill_dependent_data
      return if db.nil?

      icollection = collection.to_enum
      current = icollection.next

      db[:profiles___p].select(
        :p__id,
        :p__type,
        :p__user_id,
        :s__id___school_id,
        :s__short_name___school_name,
        :s__guid___school_guid,
        :cu__id___class_unit_id,
        :cu__name___class_unit_name,
        :p__enlisted_on,
        :p__left_on,
        :sb__id___subject_id,
        :sb__name___subject_name
      ).left_outer_join(
        :schools___s,
        {id: :school_id}
      ) { |ta|
        Sequel.qualify(ta, :deleted_at) =~ nil
      }.left_outer_join(
        :class_units___cu,
        {id: :p__class_unit_id}
      ) { |ta|
        Sequel.qualify(ta, :deleted_at) =~ nil
      }.left_outer_join(
        :competences___c,
        {profile_id: :p__id}
      ).left_outer_join(
        :subjects___sb,
        {id: :subject_id}
      ).where(
        p__deleted_at: nil,
        p__user_id: collection.map(&:id)
      ).order(
        :p__user_id,
        :p__id
      ).each do |rec|

        profile = Profile.new(
          id: rec[:id],
          type: rec[:type],
          user_id: rec[:user_id],
          left_on: rec[:left_on],
          enlisted_on: rec[:enlisted_on],
          school: {
            id: rec[:school_id],
            guid: rec[:school_guid],
            name: rec[:school_name]
          },
          class_unit: {
            id: rec[:class_unit_id],
            name: rec[:class_unit_name]
          }
        )

        while profile.user_id != current.id
          begin
            current = icollection.next
          rescue StopIteration
            break
          end
        end

        next if current.id != profile.user_id

        last_pr = current.profiles.last

        if last_pr == nil || last_pr.id != profile.id
          current.profiles << profile
        end

        next unless rec[:subject_id]

        subject = Subject.new(
          id: rec[:subject_id],
          name: rec[:subject_name]
        )

        current.profiles.last.subjects << subject
      end
    end

    def get(key, from:, default:)
      value = from[key].to_i
      value > 0 ? value : default
    end
  end
end