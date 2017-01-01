#!/usr/bin/env ruby
require 'rubygems'
require 'pp'
require 'pg'

$tables = {
  categories: ['id', 'name', 'created_at', 'updated_at'],
#  projects: ['id','name','description','version','size','dlink','logo_url','category_id', 'created_at', 'updated_at', 'assets'],
  tags: ['id', 'name', 'created_at', 'updated_at'],
  project_tags: ['id', 'project_id', 'tag_id'],
}

project_keys = ['id','name','description','version','size','dlink','logo_url','category_id', 'created_at', 'updated_at', 'assets'],

def seed_sql file, sql
  file.puts sql
end

def read_table name
  file = File.open( "./tmp_#{name}.sql", "w")
  vs = $conn.exec("select * from #{name}").values
  vs.each do |v|
    valuestr = v.join("','")
    sql = "insert into #{name}(#{$tables[name].join(',')}) values('#{valuestr}');"
    seed_sql file, sql
  end
  file.close
end

$conn = PG.connect(host: '127.0.0.1', dbname: 'journey', user: "postgres", password: "postgres")
pp '1.postgres connect origin success'

$tables.keys.each do |t|
  read_table t
  pp "#{t} done."
end


projects = $conn.exec("select * from projects").values

project_file = File.open( "./tmp_project.sql", "w") 
projects.each do |p|
  project_assets = $conn.exec("select url from project_assets where project_id = " + p[0]).values
  p.pop()
  p.push("{#{project_assets.join(',')}}")
  valuestr = p.join("','")
  sql = "insert into projects(#{project_keys.join(',')}) values('#{valuestr}');"
  project_file.puts sql
end
project_file.close
