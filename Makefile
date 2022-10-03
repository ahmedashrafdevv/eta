ms:
	mysql -u root -h 127.0.0.1 --password=secret < db/db.sql && \
	mysql -u root -h 127.0.0.1 --password=secret < db/seed.sql  

migrate:
	mysql -u root -h 127.0.0.1 --password=secret < db/db.sql 


seed:
	mysql -u root -h 127.0.0.1 --password=secret < db/seed.sql

proc:
	mysql -u root -h 127.0.0.1 --password=secret < db/proc.sql 

dball:
	mysql -u root -h 127.0.0.1 --password=secret < db/db.sql && \
	mysql -u root -h 127.0.0.1 --password=secret < db/seed.sql && \
	mysql -u root -h 127.0.0.1 --password=secret < db/proc.sql 