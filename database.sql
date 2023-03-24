CREATE DATABASE activitytracker_db;

USE activitytracker_db;

CREATE TABLE activities (
    activity_id int NOT NULL AUTO_INCREMENT,
    user_name varchar(255),
    activity_type varchar(255),
    activity_length int,
    distance float,
    PRIMARY KEY (activity_id)
)

INSERT INTO activities(user_name, activity_type, activity_length, distance) VALUES ("Karol Pawlak", "Running", 5000, 16);