package patients

var (
	QueryInsertPatient  = `INSERT INTO patients(id,name,lastname,address,dni,dateup) VALUES(?,?,?,?,?,?)`
	QueryGetPatientById = `SELECT * FROM patients WHERE id = ?`
	QueryUpdatePatient  = `UPDATE patients SET name = ?, lastname = ?, address = ?, dni = ? WHERE id = ?`
	QueryDeletePatient  = `DELETE FROM patients WHERE id = ?`
)
