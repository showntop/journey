#!/usr/bin/env ruby
require 'rubygems'
require 'pp'
require 'pg'



$conn = PG.connect(host: '182.92.82.68', dbname: 'journey', user: "journey", password: "Journey2016")
pp '1.postgres connect origin success'

projects = $conn.exec("select * from projects").values
# projects.each do |p|
# 	dddd = p[2].gsub(/\r\n/, "<br>").gsub(/\\r\\n/,"<br>").gsub(/\\n/, "<br>").gsub("'","''")	
# 	update_sql = "update projects set description='#{dddd}' where projects.id=#{p[0]}"
# 	pp update_sql
# 	$conn.exec(update_sql)
# end

projects.each do |p|
	info = p[2].gsub(/<br><br><br>/, ",").gsub(/<br><br>/, ",").gsub(/<br>/, ",")

	index = info.index("破解介绍:")
	index.nil? ? tmp = info[5, 30] : tmp = info[index + 5,30]
	tmp = tmp.gsub(/'/, "")
	update_sql = "update projects set intro='#{tmp}' where projects.id=#{p[0]}"
	pp update_sql
	$conn.exec(update_sql)
end

