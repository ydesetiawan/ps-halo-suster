DROP INDEX IF EXISTS idx_medical_patients_identity_number;
DROP INDEX IF EXISTS idx_medical_patients_phone_number;
DROP INDEX IF EXISTS idx_medical_patients_name;
DROP INDEX IF EXISTS idx_medical_patients_gender;
DROP INDEX IF EXISTS idx_medical_records_id;
DROP INDEX IF EXISTS idx_medical_records_identity_number;
DROP INDEX IF EXISTS idx_medical_records_created_at;

DROP TABLE IF EXISTS medical_records;
DROP TABLE IF EXISTS medical_patients;