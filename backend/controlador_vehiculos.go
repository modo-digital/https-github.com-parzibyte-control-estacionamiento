package main

type Vehiculo struct {
	Id           int    `json:"id"`
	Descripcion  string `json:"descripcion"`
	Placas       string `json:"placas"`
	FechaEntrada string `json:"fechaEntrada"`
	FechaSalida  string `json:"fechaSalida"`
}

func registrarNuevoVehiculo(vehiculo Vehiculo) error {
	bd, err := obtenerBD()
	if err != nil {
		return err
	}

	defer bd.Close()
	_, err = bd.Exec(`INSERT INTO vehiculos(descripcion,placas,fecha_entrada)
	VALUES
	(?, ?, ?)`, vehiculo.Descripcion, vehiculo.Placas, vehiculo.FechaEntrada)
	if err != nil {
		return err
	}
	return nil
}

func establecerFechaSalida(IdVehiculo int64, FechaSalida string) error {
	bd, err := obtenerBD()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE vehiculos SET fecha_salida = ? WHERE id = ?", FechaSalida, IdVehiculo)
	if err != nil {
		return err
	}
	return nil
}

func obtenerVehiculos(fechaInicio, fechaFin string) ([]Vehiculo, error) {
	vehiculos := []Vehiculo{}
	bd, err := obtenerBD()
	if err != nil {
		return vehiculos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`SELECT id, placas, descripcion, fecha_entrada, fecha_salida FROM vehiculos WHERE fecha_entrada >= ? AND fecha_entrada <= ?`, fechaInicio, fechaFin)
	if err != nil {
		return vehiculos, err
	}
	defer filas.Close()
	var vehiculo Vehiculo
	for filas.Next() {
		err := filas.Scan(&vehiculo.Id, &vehiculo.Placas, &vehiculo.Descripcion, &vehiculo.FechaEntrada, &vehiculo.FechaSalida)
		if err != nil {
			return vehiculos, err
		}
		vehiculos = append(vehiculos, vehiculo)
	}
	return vehiculos, nil
}
