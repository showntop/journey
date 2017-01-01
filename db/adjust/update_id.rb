#!/usr/bin/env ruby
require 'rubygems'
require 'pp'
require 'pg'


tables = ['category', 'tag', 'project', 'project_tag']

$conn = PG.connect(host: '127.0.0.1', dbname: 'journey', user: "postgres", password: "postgres")
pp '1.postgres connect origin success'

categories = $conn.exec("select * from categories").values
pp categories.size

categories.each do |c|
  next if c[0].to_i >= 10000
  $conn.exec("update categories set id= #{c[0].to_i+ 10000} where id = #{c[0]}")
end


projects = $conn.exec("select * from projects").values

projects.each do |c|
  next if c[0].to_i >= 10000
  $conn.exec("update projects set id= #{c[0].to_i+ 10000} where id = #{c[0]}")
end

tags = $conn.exec("select * from tags").values

tags.each do |c|
  next if c[0].to_i >= 10000
  $conn.exec("update tags set id= #{c[0].to_i+ 10000} where id = #{c[0]}")
end
