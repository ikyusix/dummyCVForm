create schema user_dtls;

create table user_dtls.profile_dtls
(
    profile_code     int not null
        primary key
        unique,
    wanted_job_title varchar,
    first_name       varchar,
    last_name        varchar,
    email            varchar,
    phone            varchar,
    country          varchar,
    city             varchar,
    address          varchar,
    postal_code      int,
    driving_license  varchar,
    nationality      varchar,
    place_of_birth   varchar,
    date_of_birth    varchar,
    photo_url        varchar,
    obj              text,
    del_flg          varchar,
    mod_id           varchar,
    mod_time         timestamp,
    cre_id           varchar,
    cre_time         timestamp
);

create table user_dtls.job_dtls
(
    prof_id    int
        constraint job_dtls_profile_dtls_profile_code_fk
            references user_dtls.profile_dtls,
    id         int not null,
    job_title  varchar,
    employer   varchar,
    start_date varchar,
    end_date   varchar,
    city       varchar,
    job_desc   varchar,
    del_flg    varchar,
    mod_id     varchar,
    mod_time   timestamp,
    cre_id     varchar,
    cre_time   timestamp
);

create table user_dtls.edu_dtls
(
    prof_id    int
        constraint edu_dtls_profile_dtls_profile_code_fk
            references user_dtls.profile_dtls,
    id         int not null,
    school     varchar,
    degree     varchar,
    start_date varchar,
    end_date   varchar,
    city       varchar,
    edu_desc   varchar,
    del_flg    varchar,
    mod_id     varchar,
    mod_time   timestamp,
    cre_id     varchar,
    cre_time   timestamp
);

create table user_dtls.skill_dtls
(
    prof_id    int
        constraint skill_dtls_profile_dtls_profile_code_fk
            references user_dtls.profile_dtls,
    id         int not null,
    skill      varchar,
    level_desc varchar,
    del_flg    varchar,
    mod_id     varchar,
    mod_time   timestamp,
    cre_id     varchar,
    cre_time   timestamp
);





