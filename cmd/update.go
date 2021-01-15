package cmd

import (
	"fmt"

	"github.com/ocean2811/ddns-aliyun/pkg/ddns"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	gUpdateAccessKeyID     string
	gUpdateAccessKeySecret string
	gType                  string
	gSrc                   string
	gDest                  string
)

const (
	gcTypeAddress    = "A"
	gcTypeForwardURL = "F"
)

func init() {
	updateCmd.Flags().StringVarP(&gUpdateAccessKeyID, "access-key-id", "i", "", "aliyun access key ID (required)")
	updateCmd.MarkFlagRequired("access-key-id")

	updateCmd.Flags().StringVarP(&gUpdateAccessKeySecret, "access-key-secret", "s", "", "aliyun access key Secret (required)")
	updateCmd.MarkFlagRequired("access-key-secret")

	updateCmd.Flags().StringVarP(&gType, "type", "t", "", "DNS type: [F|A]. 'F':URL to to another URL; 'A':URL to IP ")
	updateCmd.MarkFlagRequired("type")

	updateCmd.Flags().StringVarP(&gSrc, "src", "", "", "your source URL. DNS will convert source URL to destination URL(IP)")
	updateCmd.MarkFlagRequired("src")

	updateCmd.Flags().StringVarP(&gDest, "dest", "", "", "your destination URL(IP). DNS will convert source URL to destination URL(IP).\ndest will be current public IP if not specified.")

	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the specified DNS record",
	RunE: func(cmd *cobra.Command, args []string) error {
		recordType := ""
		switch gType {
		case gcTypeAddress:
			recordType = "A"
		case gcTypeForwardURL:
			recordType = "FORWARD_URL"
		default:
			return errors.New("type has error")
		}

		client, err := ddns.NewDDNSAndConnect(gUpdateAccessKeyID, gUpdateAccessKeySecret)
		if err != nil {
			return err
		}

		if len(gDest) == 0 {
			if gType == gcTypeForwardURL {
				return errors.New("type 'F' must with 'dest' flag")
			}

			gDest, err = getPublicIP()
			if err != nil {
				return err
			}
		}

		records, err := client.DescribeSubDomainRecords(gSrc)
		if err != nil {
			return err
		}
		// fmt.Println(records, err)

		// Update Record
		for _, r := range records {
			if recordType != r.Type {
				continue
			}

			if r.Value == gDest { //No need to change
				fmt.Printf("No need to change.url=%s,current=%s\n", gSrc, gDest)
				return nil
			}

			//update record
			err = client.UpdateDomainRecord(r.RecordId, r.Type, r.RR, gDest)
			if err != nil {
				return err
			}

			fmt.Printf("Update domain record. url=%s,from=%s,to=%s\n", gSrc, r.Value, gDest)
			return nil
		}

		// Add a new record
		err = client.AddDomainRecord(recordType, gSrc, gDest)
		if err != nil {
			return err
		}
		fmt.Printf("Add domain record. url=%s,record=%s\n", gSrc, gDest)

		return err
	},
}
