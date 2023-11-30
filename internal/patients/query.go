package patients

var (
	QueryInsertPatient = `INSERT INTO patient(id,name,lastname,address,dni,dateup)
	VALUES(?,?,?,?,?,?)`
	QueryDeletePatient  = `DELETE FROM patient WHERE id = ?`
	QueryGetPatientById = `SELECT id, name, lastname, address, dni, dateup
	FROM patient WHERE id = ?`
	QueryUpdatePatient = `UPDATE patient SET name = ?, lastname = ?, address = ?, dni = ?, dateup = ?
	WHERE id = ?`
)
