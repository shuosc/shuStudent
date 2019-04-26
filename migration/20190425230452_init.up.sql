create table student
(
    id          varchar(16) not null,
    name        varchar(128),
    mail        varchar(256),
    phoneNumber varchar(128)
);

create unique index student_id_uindex
    on student (id);

alter table student
    add constraint student_pk
        primary key (id);

