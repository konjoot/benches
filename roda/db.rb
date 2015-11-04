module DB
  DB_PARAMS = {
    adapter:   'postgres',
    host:      'localhost',
    database:  'lms2_db_dev',
    user:      'lms2_db_user',
    password:  'lms_2014'
  }

  @@conn = nil

  private

  def db
    @@conn ||= Sequel.connect(DB_PARAMS)
  end

  module_function :db
end