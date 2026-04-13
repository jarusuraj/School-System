CREATE TABLE
    class_subjects (
        id SERIAL PRIMARY KEY,
        class_id INT NOT NULL REFERENCES classes (id) ON DELETE CASCADE,
        subject_id INT NOT NULL REFERENCES subjects (id) ON DELETE CASCADE,
        subject_code TEXT NOT NULL
    );