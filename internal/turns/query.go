package turns

var (
	QueryInsertTurn  = `INSERT INTO turns(patients_id, dentists_id, date, description) VALUES (?,?,?,?)`
	QUeryGetTurnById = `SELECT * FROM turns WHERE id = ?`
	QuertyUpdateTurn = `UPDATE turns SET patients_id = ?, dentists_id = ?, date = ?, description = ? WHERE id = ? `
	QueryDeleteTurn  = `DELETE FROM turns WHERE id = ?`
)
