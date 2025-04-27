-- Classes Table
CREATE TABLE classes (
    id CHAR(36) PRIMARY KEY,
    level INT NOT NULL,
    branch VARCHAR(2) NOT NULL
);

-- Trigger
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


-- Students Table
CREATE TABLE students (
    id CHAR(36) PRIMARY KEY,
    class_id CHAR(36) NOT NULL,
    student_image VARCHAR(255),
    school_number BIGINT NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    created_at BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes(id)
);

-- Trigger
DELIMITER //
CREATE TRIGGER students_before_insert
BEFORE INSERT ON students
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
    IF NEW.created_at IS NULL THEN
        SET NEW.created_at = UNIX_TIMESTAMP();
    END IF;
END;
//
DELIMITER ;


-- Attendances Table
CREATE TABLE attendances (
    id CHAR(36) PRIMARY KEY,
    class_id CHAR(36) NOT NULL,
    created_at BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes(id)
);

-- Trigger
DELIMITER //
CREATE TRIGGER attendances_before_insert
BEFORE INSERT ON attendances
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
    IF NEW.created_at IS NULL THEN
        SET NEW.created_at = UNIX_TIMESTAMP() * 1000; 
    END IF;
END;
//
DELIMITER ;

ALTER TABLE attendances
ADD UNIQUE INDEX idx_class_date (class_id, (DATE(FROM_UNIXTIME(created_at / 1000))));


-- AttendanceDetails Table
CREATE TABLE attendance_details (
    id CHAR(36) PRIMARY KEY,
    attendance_id CHAR(36) NOT NULL,
    student_id CHAR(36) NOT NULL,
    is_present BOOLEAN NOT NULL,
    FOREIGN KEY (attendance_id) REFERENCES attendances(id),
    FOREIGN KEY (student_id) REFERENCES students(id)
);

--Trigger
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
