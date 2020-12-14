package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// retorna a aplicação de linha de comando para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "uol.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca nomes de servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarServidores(c *cli.Context){
	host := c.String("host")

	names, error := net.LookupNS(host)
	if error != nil{
		log.Fatal(error)
	}

	for _, name := range names {
		fmt.Println(name.Host)
	}
}

func buscarIps(c *cli.Context){
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil{
		log.Fatal(error)
	}

	for _, ip := range ips{
		fmt.Println(ip)
	}
}