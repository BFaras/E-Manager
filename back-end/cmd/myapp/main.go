package main

import (
   "fmt"
   "os"
   "back-end/cmd/myapp/cli"
   "go.uber.org/zap"
)

func main() {
   if err := cli.Execute(); err != nil {
      _, _ = fmt.Fprintf(os.Stderr, "Some error occurred during execute app. Error: %v\n", err)
      os.Exit(2)
   }
}

func init() {
   zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}
