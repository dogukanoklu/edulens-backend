-- Classes Table and Trigger
CREATE TABLE classes (
    id CHAR(36) PRIMARY KEY,
    level INT NOT NULL,
    branch VARCHAR(2) NOT NULL
);

DELIMITER //
CREATE TRIGGER classes_before_insert
BEFORE INSERT ON classes
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END;
//
DELIMITER ;

-- Students Table and Trigger
CREATE TABLE students (
    id CHAR(36) PRIMARY KEY,
    class_id CHAR(36) NOT NULL,
    profile_picture VARCHAR(255),
    school_number BIGINT NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    created_at BIGINT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    FOREIGN KEY (class_id) REFERENCES classes(id)
);

DELIMITER //
CREATE TRIGGER students_before_insert
BEFORE INSERT ON students
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END;
//
DELIMITER ;

-- Attendances Table and Trigger
CREATE TABLE attendances (
    id CHAR(36) PRIMARY KEY,
    class_id CHAR(36) NOT NULL,
    created_at BIGINT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    FOREIGN KEY (class_id) REFERENCES classes(id)
);

DELIMITER //
CREATE TRIGGER attendances_before_insert
BEFORE INSERT ON attendances
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END;
//
DELIMITER ;

ALTER TABLE Attendances
ADD UNIQUE INDEX idx_class_date (class_id, (DATE(FROM_UNIXTIME(created_at))));

-- AttendanceDetails Table and Trigger
CREATE TABLE attendance_details (
    id CHAR(36) PRIMARY KEY,
    attendance_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    is_present BOOLEAN NOT NULL,
    FOREIGN KEY (attendance_id) REFERENCES attendances(id),
    FOREIGN KEY (user_id) REFERENCES students(id)
);

DELIMITER //
CREATE TRIGGER attendance_details_before_insert
BEFORE INSERT ON attendance_details
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END;
//
DELIMITER ;