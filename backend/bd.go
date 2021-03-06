package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func obtenerBD() (*sql.DB, error) {
	return sql.Open("sqlite3", NombreBaseDeDatos)
}

func crearTablas() error {
	db, err := obtenerBD()
	if err != nil {
		return err
	}
	defer db.Close()
	tablas := []string{
		`CREATE TABLE IF NOT EXISTS vehiculos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			placas TEXT NOT NULL,
			descripcion TEXT NOT NULL,
			fecha_entrada TEXT NOT NULL,
			fecha_salida TEXT NOT NULL DEFAULT ""
		);`,
		`CREATE TABLE IF NOT EXISTS ajustes_costo(
			costo_hora REAL NOT NULL,
			minutos_redondear INTEGER NOT NULL,
			minutos_tolerancia INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS pagos_vehiculos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			id_vehiculo INTEGER NOT NULL,
			pago REAL NOT NULL,
			minutos INTEGER NOT NULL,
			FOREIGN KEY(id_vehiculo) REFERENCES vehiculos(id) ON DELETE CASCADE ON UPDATE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS ajustes_impresora(
			nombre_impresora TEXT NOT NULL
		);`,
	}
	for _, tabla := range tablas {
		_, err = db.Exec(tabla)
		if err != nil {
			return err
		}
	}
	return nil
}
