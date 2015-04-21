package main

import (
	"encoding/csv"
	"fmt"
	// "io/ioutil"
	// "io"
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

	equals(t, 6, len(data))

}

func TestOpenCSVFile(t *testing.T) {
	file, err := os.Open("./sso_db_raw.csv")
	ok(t, err)

	defer file.Close()

	reader := csv.NewReader(file)
	ok(t, err)

	data, err := reader.ReadAll()
	ok(t, err)

	equals(t, 3758, len(data))
}

func TestMarshalCSV(t *testing.T) {
	csvStream := `Overflow Type,Municipality/Facility,NPDES #,Date Discovered,Time Discovered,Days,Minutes,Hours,Location,Zip Code,Latitude,Longitude,Collection-System,Quantity in Gallons (Estimated),Net in Gallons (Estimated),Cause,Watershed,Receiving waters,County,Comments,Penalty Collected,Penalty Notes
SSO,American Water Military Services,N/A,3/29/2011,4:45:00 PM,0,1,0,"8133B Lawson Loop, Fort Meade",20724,,,Fort Meade WWTP,89,89,Baby whips & debris,,Unknown,Anne Arundel,None,,
SSO,"Army, Department of",N/A,10/21/2009,11:00:00 AM,0,1,0,"MH # 7202 off 33 Calvery Rd., Ft. Meade",20755,,,Ft. Meade WWTP,500,500,Baby wipes & rags,,Unknown,Anne Arundel,None,,
SSO,City of Baltimore,N/A,1/1/2005,8:16:00 PM,0,2,0,5300 Falls Rd,21209,,,Patapsco WWTP,600,600,Blockage,,Jones Falls,City of Baltimore,None,,
SSO,City of Baltimore,N/A,1/2/2005,12:03:00 PM,0,3,0,6200 Ship View Way,21224,,,Back River WWTP,891,891,Blockage,,Inner Harbor,City of Baltimore,None,,
SSO,City of Baltimore,N/A,1/2/2005,10:00:00 PM,0,1,0,1800 Park Ave,21217,,,Patapsco WWTP,300,300,Blockage,,Jones Falls,City of Baltimore,None,,
SSO,"Ridgely, Commissioners of",N/A,1/3/2005,9:00:00 AM,0,0,20,Between Liberty & Bell Sts,21660,,,,10,10,Blockage,,Unknown,Caroline,Collection system not provided,,
SSO,City of Baltimore,N/A,1/3/2005,10:30:00 AM,1,4,15,4601 Franklintown Rd,21229,,,Patapsco WWTP,24975,24975,Blockage,,Gwynn's Falls,City of Baltimore,None,,
SSO,WSSC,N/A,1/3/2005,3:23:00 PM,0,21,0,"18101-57 Town Center Dr, MH/26003011",20832,,,,1890,1890,Blockage,,Unknown,Montgomery,Collection system not provided,,
SSO,WSSC,N/A,1/4/2005,11:40:00 AM,0,8,0,"6800 Killarney St, MH/04021030U",20735,,,,1080,1080,Blockage,,Unnamed tributary,Prince George's,Collection system not provided,,
SSO,Baltimore County DPW,N/A,1/5/2005,2:30:00 PM,0,1,0,2539 Cedar Lane,21207,,,Patapsco WWTP,200,200,Blockage,,Unknown,Baltimore County,None,,
SSO,Baltimore County DPW,N/A,1/9/2005,1:05:00 PM,0,1,0,976 Sandalwood Rd,21221,,,Back River WWTP,100,100,Blockage,,Deep Creek,Baltimore County,None,,
SSO,City of Baltimore,N/A,1/9/2005,2:30:00 PM,0,3,0,1 Edgevale Rd,21210,,,Patapsco WWTP,1260,1260,Blockage,,Jones Falls,City of Baltimore,None,,
SSO,City of Baltimore,N/A,1/10/2005,8:00:00 PM,0,1,0,3524 Lyndales Ave,21213,,,Back River WWTP,300,300,Blockage,,Herring Run,City of Baltimore,None,,
SSO,WSSC,N/A,1/11/2005,2:30:00 PM,0,3,0,"700 Quince Orchard Rd, Gaithersburg, MH #15010027",20878,,,,360,360,Blockage,,Unknown,Montgomery,Collection system not provided,,
SSO,City of Baltimore,N/A,1/13/2005,9:00:00 AM,0,4,0,4600 Franklintown Rd,21216,,,Patapsco WWTP,3600,3600,Blockage,,Gwynn's Falls,City of Baltimore,None,,
SSO,WSSC,N/A,1/14/2005,1:40:00 PM,0,0,12,"6902 Kent Town Dr, MH 03-025-002u",20785,,,,500,500,Blockage,,Unknown,Prince George's,Collection system not provided,,
SSO,Town of Federalsburg,N/A,1/16/2005,12:30:00 AM,0,24,0,MH before PS,21632,,,,20000,20000,Blockage,,Marshyhope Creek,Caroline,Collection system not provided,,
SSO,Baltimore County DPW,N/A,1/17/2005,11:30:00 AM,0,1,20,Cantwell Rd & Giard Dr,21244,,,,150,150,Blockage,,Unknown,Baltimore County,Collection system not provided,,
SSO,City of Baltimore,N/A,1/18/2005,10:00:00 AM,0,1,0,4600 Franklintown Rd,21216,,,Patapsco WWTP,300,300,Blockage,,Gwynn's Falls,City of Baltimore,Lat/Long could not be obtained with the address provided,,`

	reader := csv.NewReader(strings.NewReader(csvStream))
	reader.FieldsPerRecord = 22
}

func TestUnmarshalInRecords(t *testing.T) {
	file, err := os.Open("./sso_db_raw.csv")
	ok(t, err)

	defer file.Close()

	reader := csv.NewReader(file)
	ok(t, err)

	// data, err := reader.Read()
	data, err := UnmarshalInRecords(reader)
	ok(t, err)

	equals(t, 3758, len(data))
}

func TestGeocodeInRecords(t *testing.T) {
	file, err := os.Open("./sso_db_raw.csv")
	ok(t, err)

	defer file.Close()

	reader := csv.NewReader(file)
	ok(t, err)

	// data, err := reader.Read()
	data, err := UnmarshalInRecords(reader)
	ok(t, err)

	gc := NewGeocoder()
	for i := 0; i < 10; i++ {
		gc.SetUrlValues(data[i])

		fmt.Println(data[i])
		fmt.Println(string(gc.Geocode()) + "\n")
	}

}
