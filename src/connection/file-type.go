package connection

type fileLines []*fileType

type fileType struct {
	cpf                string
	private            bool
	incompleto         bool
	dataUltimaCompra   string
	compraTicketMedio  float64
	ticketUltimaCompra float64
	cnpjLojaFrequente  string
	cnpjUltimaLoja     string
}
