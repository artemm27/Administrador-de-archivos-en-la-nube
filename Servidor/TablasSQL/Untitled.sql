#crear el esquema
#CREATE SCHEMA `GoFiles` ;

#crear tabla de usuarios registrados
#drop table Users;
create table Users(
	UserId varchar(255),
    Passw varchar(255),
    UserPath varchar (255)
);


select * from Users;

#vaciar la tabla
truncate table Users;