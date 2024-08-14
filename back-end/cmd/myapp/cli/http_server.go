package cli


import (
   "github.com/spf13/cobra"
   "back-end/internal/infrastructure/api/rest"
)


const HttpServerCommand = "http-server"
const VersionHttpServer = "1.0.0"

var httpServer = &cobra.Command{
   Use:     HttpServerCommand,
   Short:   "Start http server",
   Version: VersionHttpServer,
   RunE: func(cmd *cobra.Command, args []string) (err error) {

      cnf := cfg.Sub("app.api.rest")

      var server *rest.Server
      if server, err = rest.New(cnf); err != nil {
         return err
      }

      return server.Start(cmd.Context())
   },
}

func init() {
   cmd.AddCommand(httpServer)
}
