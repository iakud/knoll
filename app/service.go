package app

type Service interface {
	Init() error
	Start() error
	Stop() error
}

var services []Service

func Register(service Service) {
	services = append(services, service)
}

func initServices() error {
	for _, service := range services {
		if err := service.Init(); err != nil {
			return err
		}
	}
	return nil
}

func startServices() error {
	for _, service := range services {
		if err := service.Start(); err != nil {
			return err
		}
	}
	return nil
}

func stopServices() error {
	for _, service := range services {
		if err := service.Stop(); err != nil {
			return err
		}
	}
	return nil
}
