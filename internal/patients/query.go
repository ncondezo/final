package patients

var (
	QueryInsertPatient  = `INSERT INTO patients(name, lastname, address, dni, dateup) VALUES(?,?,?,?,?)`
	QueryGetPatientById = `SELECT * FROM patients WHERE id = ?`
	QueryUpdatePatient  = `UPDATE patients SET name = ?, lastname = ?, address = ?, dni = ? WHERE id = ?`
	QueryPatchPatient   = `UPDATE patients SET dni = ? WHERE id = ?`
	QueryDeletePatient  = `DELETE FROM patients WHERE id = ?`
)
