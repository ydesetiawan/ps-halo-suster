CREATE TABLE medical_patients (
    identity_number BIGINT NOT NULL PRIMARY KEY CHECK (length(identity_number::text) = 16),
    phone_number VARCHAR(15) NOT NULL CHECK (phone_number ~ '^\+62[0-9]{8,13}$'),
    name VARCHAR(30) NOT NULL CHECK (length(name) >= 3 AND length(name) <= 30),
    birth_date DATE NOT NULL,
    gender VARCHAR(6) NOT NULL CHECK (gender IN ('male', 'female')),
    identity_card_scan_img TEXT NOT NULL CHECK (identity_card_scan_img ~ '^(http|https)://'),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create indexes if necessary
CREATE INDEX idx_medical_patients_identity_number ON medical_patients (identity_number);
CREATE INDEX idx_medical_patients_phone_number ON medical_patients (phone_number);
CREATE INDEX idx_medical_patients_name ON medical_patients (name);
CREATE INDEX idx_medical_patients_gender ON medical_patients (gender);
CREATE INDEX idx_medical_patients_created_at ON medical_patients (created_at);

CREATE TABLE medical_records (
     id char(26) PRIMARY KEY,
     identity_number BIGINT NOT NULL,
     symptoms TEXT NOT NULL CHECK (length(symptoms) >= 1 AND length(symptoms) <= 2000),
     medications TEXT NOT NULL CHECK (length(medications) >= 1 AND length(medications) <= 2000),
     created_by char(26) NOT NULL ,
     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
     FOREIGN KEY (identity_number) REFERENCES medical_patients(identity_number),
     FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE INDEX idx_medical_records_id ON medical_records (id);
CREATE INDEX idx_medical_records_identity_number ON medical_records (identity_number);
CREATE INDEX idx_medical_records_created_at ON medical_records (created_at);