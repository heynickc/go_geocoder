package main

import (
	"encoding/csv"
	"fmt"
	// "io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
)

type InAddress struct {
	Address string
	Zip     string
}

func TestCSVDecoding(t *testing.T) {

	csvStream := `Overflow Type,Municipality/Facility,NPDES #,Date Discovered,Time Discovered,"Duration (Days, hours, Minutes)",,,Location,Zip Code,Latitude,Longitude,Collection-System,Quantity in Gallons (Estimated),Net in Gallons (Estimated),Cause,Watershed,Receiving waters,County,Comments,Penalty Collected,Penalty Notes
SSO,American Water Military Services,N/A,3/29/11,4:45:00 PM,0,1,0,"8133B Lawson Loop, Fort Meade",20724,,,Fort Meade WWTP,89,89,Baby whips & debris,,Unknown,Anne Arundel,None,,`

	reader := csv.NewReader(strings.NewReader(csvStream))
	data, err := reader.ReadAll()

	ok(t, err)
	equals(t, "Overflow Type", data[0][0])
}

func TestGetCSVAddressHeaders(t *testing.T) {

	csvStream := `Overflow Type,Municipality/Facility,NPDES #,Date Discovered,Time Discovered,"Duration (Days, hours, Minutes)",,,Location,Zip Code,Latitude,Longitude,Collection-System,Quantity in Gallons (Estimated),Net in Gallons (Estimated),Cause,Watershed,Receiving waters,County,Comments,Penalty Collected,Penalty Notes
SSO,American Water Military Services,N/A,3/29/11,4:45:00 PM,0,1,0,"8133B Lawson Loop, Fort Meade",20724,,,Fort Meade WWTP,89,89,Baby whips & debris,,Unknown,Anne Arundel,None,,`

	reader := csv.NewReader(strings.NewReader(csvStream))
	data, err := reader.ReadAll()
	headerCol := data[0]

	streetIndex := sort.SearchStrings(headerCol, "Location")
	zipIndex := sort.SearchStrings(headerCol, "Zip Code")

	ok(t, err)

	equals(t, 8, streetIndex)
	equals(t, 22, zipIndex)
}

func TestLineFeedCSVDecoder(t *testing.T) {

	csvStream := `Overflow Type,Municipality/Facility,NPDES #,Date Discovered,Time Discovered,"Duration (Days, hours, Minutes)",,,Location,Zip Code,Latitude,Longitude,Collection-System,Quantity in Gallons (Estimated),Net in Gallons (Estimated),Cause,Watershed,Receiving waters,County,Comments,Penalty Collected,Penalty Notes
SSO,American Water Military Services,N/A,3/29/11,4:45:00 PM,0,1,0,"8133B Lawson Loop, Fort Meade",20724,,,Fort Meade WWTP,89,89,Baby whips & debris,,Unknown,Anne Arundel,None,,
SSO,"Army, Department of",N/A,10/21/09,11:00:00 AM,0,1,0,"MH # 7202 off 33 Calvery Rd., Ft. Meade",20755,,,Ft. Meade WWTP,500,500,Baby wipes & rags,,Unknown,Anne Arundel,None,,
SSO,City of Baltimore,N/A,1/1/05,8:16:00 PM,0,2,0,5300 Falls Rd,21209,,,Patapsco WWTP,600,600,Blockage,,Jones Falls,City of Baltimore,None,,
SSO,City of Baltimore,N/A,1/2/05,12:03:00 PM,0,3,0,6200 Ship View Way,21224,,,Back River WWTP,891,891,Blockage,,Inner Harbor,City of Baltimore,None,,
SSO,City of Baltimore,N/A,1/2/05,10:00:00 PM,0,1,0,1800 Park Ave,21217,,,Patapsco WWTP,300,300,Blockage,,Jones Falls,City of Baltimore,None,,`

	reader := csv.NewReader(strings.NewReader(csvStream))

	reader.FieldsPerRecord = 22
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	data, err := reader.ReadAll()
	ok(t, err)

	fmt.Println(data)

}

func TestOpenCSVFile(t *testing.T) {
	file, err := os.Open("./sso_db.csv")
	ok(t, err)

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 22
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	ok(t, err)

	data, err := reader.ReadAll()
	ok(t, err)

	fmt.Println(data)

	// for _, v := range data {
	// 	fmt.Print(v[8])
	// }
}
