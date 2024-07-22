CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE resumes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    owner_id VARCHAR(255) NOT NULL,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL,
    update_time TIMESTAMP WITH TIME ZONE,
    resume_title TEXT NOT NULL,
    data_content TEXT NOT NULL,
    is_active bool DEFAULT true,
    is_draft bool DEFAULT true
);

CREATE TABLE vacancies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    owner_id VARCHAR(255) NOT NULL,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL,
    update_time TIMESTAMP WITH TIME ZONE,
    vacancy_title TEXT NOT NULL,
    data_content TEXT NOT NULL,
    is_active bool DEFAULT true,
    is_draft bool DEFAULT true
);