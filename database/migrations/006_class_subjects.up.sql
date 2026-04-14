CREATE TABLE class_subjects (
    id SERIAL PRIMARY KEY,
    class_no INT NOT NULL REFERENCES classes(class_no) ON DELETE CASCADE,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    subject_code TEXT NOT NULL,
    UNIQUE(class_no, subject_id)
);