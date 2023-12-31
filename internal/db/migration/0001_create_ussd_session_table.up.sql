CREATE TABLE ussd_session
(
    session_id            VARCHAR(255) PRIMARY KEY,
    msisdn                VARCHAR(255) NOT NULL,
    next_stage            VARCHAR(255),
    country_ame           VARCHAR(255),
    amount                NUMERIC,
    foreign_currency_code VARCHAR(255),
    session_start_time    TIMESTAMP
);
