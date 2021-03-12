package connection

// FileLines append any FileType
type FileLines []*FileType

// FileType from imported archive
type FileType struct {
	Cpf                string
	Private            bool
	Incompleto         bool
	DataUltimaCompra   string
	CompraTicketMedio  float64
	TicketUltimaCompra float64
	CnpjLojaFrequente  string
	CnpjUltimaLoja     string
}
