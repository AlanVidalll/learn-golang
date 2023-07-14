package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generated vai retornar a aplicação de linha de comando pronta pra ser executada
func Generated() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Buscas Ips e Nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca Ips de endereços na internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca nomes do sevidores na internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host) //NS = name server
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server)
	}

}
