package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ocean2811/ddns-aliyun/pkg/myip"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get public IP.",
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := getPublicIP()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(ip)
	},
}

func getPublicIP() (string, error) {
	ipResult := make(chan string, 1)
	ipErr := make(chan error, myip.DetectorNum)

	go func() {
		ip, err := myip.GetByIPIP()
		if err != nil {
			debugPrintf("GetByIPIP: %+v\n", err)

			select {
			case ipErr <- err:
			default:
			}
			return
		}

		debugPrintf("GetByIPIP: %v\n", ip)

		select {
		case ipResult <- ip:
		default:
		}
	}()

	go func() {
		ip, err := myip.GetByIFConfig()
		if err != nil {
			debugPrintf("GetByIFConfig: %+v\n", err)

			select {
			case ipErr <- err:
			default:
			}
			return
		}

		debugPrintf("GetByIFConfig: %v\n", ip)

		select {
		case ipResult <- ip:
		default:
		}
	}()

	go func() {
		ip, err := myip.GetByJSONIP()
		if err != nil {
			debugPrintf("GetByJSONIP: %+v\n", err)

			select {
			case ipErr <- err:
			default:
			}
			return
		}

		debugPrintf("GetByJSONIP: %v\n", ip)

		select {
		case ipResult <- ip:
		default:
		}
	}()

	errArray := make([]error, 0, myip.DetectorNum)
	for {
		select {
		case ip := <-ipResult:
			return ip, nil
		case err := <-ipErr:
			errArray = append(errArray, err)
			if len(errArray) >= myip.DetectorNum {
				out := strings.Builder{}
				for _, e := range errArray {
					out.WriteString(fmt.Sprintln(e))
				}
				return "", errors.New(out.String())
			}
		}
	}
}
