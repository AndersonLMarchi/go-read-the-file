# GO-READ-THE-FILE

O GO-READ-THE-FILE é um sistema responsavel pela leitura de um arquivo texto (.txt) com seguinte estrutura:

```go
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
```

Os programas irão ler, linha por linha, o arquivo e separar cada informação no seu respectivo campo da estrutura de FileType. Após a leitura é carregado os dados em uma lista de FileType e acionado os programas de persistência dos dados no Postgresql.

## Estrutura do projeto
- Pastas:
    - repository -> local para armazenamento dos arquivos texto para leitura
    - src/connection -> local para armazenamento dos arquivos do pacote de Connexão com o banco de dados
    - src/utils -> local para armazenamento de arquivos com funções úteis ao projeto

- Arquivos:
    - .env -> arquivo com a definição das variáveis de ambiente para conexão com o postgresql
    - docker-compose.yml -> arquivo com a definição do conteiner contendo as imagens e dependências que serão  usadas em nosso projeto
    - Dockerfile -> conjunto de instruções para execução do projeto pelo docker
    - go.mod e go.sum -> arquivos para definição do pacote principal da aplicação e as dependências de terceiros - usadas em nosso projeto
    - read-file.go -> arquivo principal contendo a função de execução main(). Nele está toda a lógica de processamento do arquivo para persistência

## Pre-requisitos
Necessário ter o [docker](https://www.docker.com/products/docker-desktop) instalado e configurado na máquina

## Execução
O projeto está configurado para executar apenas com o seguinte comando

```bash
docker-compose up
```
