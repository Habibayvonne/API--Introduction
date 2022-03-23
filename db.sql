create table user (
    id int primary key auto_increment,
    name varchar(100) not null,
    created datetime default current_timestamp
);

insert into user (name)
values
('Rapando'),
('Habiba'),
('Maxine'),
('Denno');