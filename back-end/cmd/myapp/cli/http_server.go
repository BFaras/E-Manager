package cli


import (
   "os"
   "github.com/spf13/cobra"
   "go.uber.org/zap"
   "go.uber.org/zap/zapcore"
   "back-end/internal/infrastructure/api/rest"
)


const HttpServerCommand = "http-server"
const VersionHttpServer = "1.0.0"

var httpServer = &cobra.Command{
   Use:     HttpServerCommand,
   Short:   "Start http server",
   Version: VersionHttpServer,
   RunE: func(cmd *cobra.Command, args []string) (err error) {
      core := zapcore.NewCore(
         zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
         zapcore.AddSync(os.Stderr),
         zapcore.DebugLevel,
      )
      log := zap.New(core)

      cnf := cfg.Sub("app.api.rest")

      var server *rest.Server
      if server, err = rest.New(cnf, log); err != nil {
         return err
      }

      return server.Start(cmd.Context())
   },
}

func init() {
   cmd.AddCommand(httpServer)
}
