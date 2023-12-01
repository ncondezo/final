package dentists

var (
	QueryInsertDentist  = `INSERT INTO dentists(id,name,surname,registry) VALUES(?,?,?,?)`
	QueryGetDentistById = `SELECT name, surname, registry FROM dentists WHERE id = ?`
	QueryUpdateDentist  = `UPDATE dentists SET name = ?, surname = ?, registry = ? WHERE id = ?`
	QueryDeleteDentist  = `DELETE FROM dentists WHERE id = ?`
)
