package patients

var (
	QueryInsertPatient = `INSERT INTO patients(id,name,lastname,address,dni,dateup)
	VALUES(?,?,?,?,?,?)`
	QueryDeletePatient  = `DELETE FROM patients WHERE id = ?`
	QueryGetPatientById = `SELECT id, name, lastname, address, dni, dateup
	FROM patients WHERE id = ?`
	QueryUpdatePatient = `UPDATE patients SET name = ?, lastname = ?, address = ?, dni = ?, dateup = ?
	WHERE id = ?`
)
