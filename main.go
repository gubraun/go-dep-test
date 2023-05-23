package main

import (
    "fmt"
    "strings"
    "os"
    "time"
    "github.com/lestrrat-go/strftime"
    "github.com/spf13/cobra"
)

var ref = time.Unix(1136239445, 123456789).UTC()

func main() {
    var echoTimes int

    var cmdPrint = &cobra.Command{
	Use:   "print [string to print]",
	Short: "Print anything to the screen",
	Long: `print is for printing anything back to the screen.
	For many years people have printed back to the screen.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	    fmt.Println("Print: " + strings.Join(args, " "))
	},
    }

    var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	    fmt.Println("Echo: " + strings.Join(args, " "))
	},
    }

    var cmdTimes = &cobra.Command{
	Use:   "times [string to echo]",
	Short: "Echo anything to the screen more times",
	Long: `echo things multiple times back to the user by providing
	a count and a string.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	    for i := 0; i < echoTimes; i++ {
		fmt.Println("Echo: " + strings.Join(args, " "))
	    }
	    // I want %L as milliseconds!
	    p, err := strftime.New(`%L`, strftime.WithMilliseconds('L'))
	    if err != nil {
		fmt.Println(err)
		return
	    }
	    p.Format(os.Stdout, ref)
	    os.Stdout.Write([]byte{'\n'})
	},
    }

    cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

    var rootCmd = &cobra.Command{Use: "app"}
    rootCmd.AddCommand(cmdPrint, cmdEcho)
    cmdEcho.AddCommand(cmdTimes)
    rootCmd.Execute()
}
