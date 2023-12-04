package dentists

var (
	QueryInsertDentist  = `INSERT INTO dentists(name, lastname, registry) VALUES(?,?,?)`
	QueryGetDentistById = `SELECT * FROM dentists WHERE id = ?`
	QueryUpdateDentist  = `UPDATE dentists SET name = ?, lastname = ?, registry = ? WHERE id = ?`
	QueryPatchDentist   = `UPDATE dentists SET registry = ? WHERE id = ?`
	QueryDeleteDentist  = `DELETE FROM dentists WHERE id = ?`
)
