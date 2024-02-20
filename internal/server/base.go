package server

type Server interface {
	Run()
	Shutdown()
}
