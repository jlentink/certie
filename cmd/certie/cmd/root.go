package cmd

import (
	"certie/internal"
	"certie/internal/render"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	timeout   int64
	port      int64
	showChain bool
	format    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "certie",
	Short: "A simple command line tool to pull information form a server's TLS certificate.",
	Long: `A simple command line tool to pull information form a server's TLS certificate.
It has been a simple hobby project to learn more about Go and the TLS protocol.
I hope it will be useful to you.`,
	Run: mainCmd,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func mainCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("No hostname provided to probe.")
		os.Exit(1)
	}
	host := internal.TargetHost{}
	host.ParseUrl(args[0])
	host.SetPort(port)
	host.SetTimeout(timeout)
	certs := internal.GetCertificates(host)
	serverCerts := certs.GetCertificates(showChain)

	fmt.Print(render.Render(format, serverCerts))

	if certs.GetHostCertificate().IsExpired {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().Int64VarP(&port, "port", "p", 443, "The port to connect to.")
	rootCmd.Flags().Int64VarP(&timeout, "timeout", "t", 5, "Connection time out.")
	rootCmd.Flags().BoolVarP(&showChain, "chain", "c", false, "Show the certificate chain.")
	rootCmd.Flags().StringVarP(&format, "format", "f", "text", "Output format. [json|text|yaml|xml]")
}
