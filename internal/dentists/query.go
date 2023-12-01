package dentists

var(
	QueryInsertDentist = `INSERT INTO storage.products(name,last_name,registration) VALUES(?,?,?)`
	QueryGetDentistById = `SELECT name, last_name, registration FROM storage.dentists WHERE id = ?`
	QueryUpdateDentist = `UPDATE storage.dentists SET name = ?, last_name = ?, registration = ? WHERE id = ?`
	QueryDeleteDentist  = `DELETE FROM storage.dentists WHERE id = ?`
)