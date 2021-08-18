USE `GolangSql`;
CREATE TABLE IF NOT EXISTS `student` (
    'student_ID' int unsigned NOT NULL,
    'First Name' text,
    'Last Name' text,
    'Age' int,
    'Major' text,
    PRIMARY KEY ('student_ID')
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHAR SET=UTF8MB4;