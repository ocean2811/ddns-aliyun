package cmd

import (
	"fmt"
	"os"

	"github.com/ocean2811/ddns-aliyun/pkg/ddns"
	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
)

var (
	gListAccessKeyID     string
	gListAccessKeySecret string
	gListDomain          string
	gProfessional        bool
)

func init() {
	listCmd.Flags().StringVarP(&gListAccessKeyID, "access-key-id", "i", "", "aliyun access key ID (required)")
	listCmd.MarkFlagRequired("access-key-id")

	listCmd.Flags().StringVarP(&gListAccessKeySecret, "access-key-secret", "s", "", "aliyun access key Secret (required)")
	listCmd.MarkFlagRequired("access-key-secret")

	listCmd.Flags().StringVarP(&gListDomain, "domain-name", "d", "", "your domain name (required). e.g. baidu.com")
	listCmd.MarkFlagRequired("domain-name")

	listCmd.Flags().BoolVarP(&gProfessional, "professional", "p", false, "professional mode")

	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all DNS records",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ddns.NewDDNSAndConnect(gListAccessKeyID, gListAccessKeySecret)
		if err != nil {
			return err
		}

		records, err := client.DescribeDomainRecords(gListDomain)
		if err != nil {
			return err
		}

		if len(records) == 0 {
			fmt.Println("empty!")
			return nil
		}

		table := tablewriter.NewWriter(os.Stdout)
		if gProfessional {
			// show raw field
			table.SetHeader([]string{"Type", "RR", "Value", "RecordId", "Status"})
			table.SetAutoFormatHeaders(false)
			table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
			table.SetAlignment(tablewriter.ALIGN_CENTER)

			for _, r := range records {
				table.Append([]string{r.Type, r.RR, r.Value, r.RecordId, r.Status})
			}
		} else {
			// user-friendly
			table.SetHeader([]string{"Src", "Dest", "RecordId", "Status"})
			table.SetAutoFormatHeaders(false)
			table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
			table.SetAlignment(tablewriter.ALIGN_CENTER)

			for _, r := range records {
				switch r.Type {
				case "FORWARD_URL":
					table.Append([]string{r.RR + "." + gListDomain, r.Value, r.RecordId, r.Status})
				case "A":
					src := r.RR + "." + gListDomain
					if r.RR == "@" {
						src = gListDomain
					}
					table.Append([]string{src, r.Value, r.RecordId, r.Status})
				default:
					continue
				}
			}
		}

		table.Render()

		return nil
	},
}
