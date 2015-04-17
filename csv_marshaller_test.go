package main

import (
	"encoding/csv"
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestCSVDecoding(t *testing.T) {

	csvStream := `Overflow Type,Municipality/Facility,NPDES #,Date Discovered,Time Discovered,Days,Hours,Minutes,Location,Zip Code,Latitude,Longitude,Collection-System,Quantity in Gallons (Estimated),Net in Gallons (Estimated),Cause,Receiving waters,County,Penalty Collected Comments,Penalty Collected,Notes
SSO,City of Baltimore,N/A,1/1/2005,8:16:00 PM,0,2,0,5300 Falls Rd,21209,,,Patapsco WWTP,600,600,Blockage,Jones Falls,City of Baltimore,,,None`

	reader := csv.NewReader(strings.NewReader(csvStream))
	data, err := reader.ReadAll()

	ok(t, err)
	fmt.Println(data[0][0])
}

func TestGetCSVAddressHeaders(t *testing.T) {

	csvStream := `Overflow Type,Municipality/Facility,NPDES #,Date Discovered,Time Discovered,Days,Hours,Minutes,Location,Zip Code,Latitude,Longitude,Collection-System,Quantity in Gallons (Estimated),Net in Gallons (Estimated),Cause,Receiving waters,County,Penalty Collected Comments,Penalty Collected,Notes
SSO,City of Baltimore,N/A,1/1/2005,8:16:00 PM,0,2,0,5300 Falls Rd,21209,,,Patapsco WWTP,600,600,Blockage,Jones Falls,City of Baltimore,,,None`

	reader := csv.NewReader(strings.NewReader(csvStream))
	data, err := reader.ReadAll()
	headerCol := data[0]

	streetIndex := sort.SearchStrings(headerCol, "Location")
	zipIndex := sort.SearchStrings(headerCol, "Zip Code")

	ok(t, err)

	fmt.Println(streetIndex)
	fmt.Println(zipIndex)
}
