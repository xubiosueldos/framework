package configuracion

type Configuracion struct {
	Ip     string
	Namedb string
}

var instance Configuracion

func GetInstance() Configuracion {

	instance = Configuracion{}
	instance.Ip = "192.168.30.111"
	instance.Namedb = "faf_multitenant_go"
	return instance
}