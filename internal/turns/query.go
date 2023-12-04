package turns

var (
	QueryInsertTurn       = `INSERT INTO turns(date, description, patients_id, dentists_id) VALUES (?,?,?,?)`
	QueryGetTurnById      = `SELECT * FROM turns INNER JOIN patients ON patients.id = turns.patients_id INNER JOIN dentists ON dentists.id = turns.dentists_id WHERE turns.id = ?`
	QueryGetTurnByPatient = `SELECT * FROM turns INNER JOIN patients ON patients.id = turns.patients_id INNER JOIN dentists ON dentists.id = turns.dentists_id WHERE turns.patients_id = ?`
	QueryUpdateTurn       = `UPDATE turns SET date = ?, description = ?, dentists_id = ? WHERE id = ?`
	QueryDeleteTurn       = `DELETE FROM turns WHERE id = ?`
)
