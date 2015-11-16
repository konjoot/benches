module DB
  DB_PARAMS = {
    adapter:   'postgres',
    host:      'localhost',
    database:  'lms2_development_2',
    user:      'lms',
    password:  ''
  }

  @@conn = nil

  private

  def db
    @@conn ||= Sequel.connect(DB_PARAMS)
  end

  module_function :db
end