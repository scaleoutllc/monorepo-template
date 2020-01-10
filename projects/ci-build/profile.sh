service redis-server start
service postgresql start
runuser -l postgres -c "psql -c \"alter role postgres with password 'secret'\"" > /dev/null