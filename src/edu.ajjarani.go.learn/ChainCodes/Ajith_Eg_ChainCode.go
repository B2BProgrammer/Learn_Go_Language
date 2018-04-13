package ChainCodes

import (
	"runtime"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"os"
	"net/http"
	"io"
)


/////////////////////////////////////////////////////////////////////////////////////////////////////
// A Map that holds ObjectNames and the number of Keys
// This information is used to dynamically Create, Update
// Replace , and Query the Ledger
// In this model all attributes in a table are strings
// The chain code does both validation
// A dummy key like 2016 in some cases is used for a query to get all rows
//
//              "User":        1, Key: UserID
//              "Item":        1, Key: ItemID
//              "UserCat":     3, Key: "2016", UserType, UserID
//              "ItemCat":     3, Key: "2016", ItemSubject, ItemID
//              "Auction":     1, Key: AuctionID
//              "AucInit":     2, Key: Year, AuctionID
//              "AucOpen":     2, Key: Year, AuctionID
//              "Trans":       2, Key: AuctionID, ItemID
//              "Bid":         2, Key: AuctionID, BidNo
//              "ItemHistory": 4, Key: ItemID, Status, AuctionHouseID(if applicable),date-time
//
// The additional key is the ObjectType (aka ObjectName or Object). The keys  would be
// keys: {"picname", "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art1.png"}
/////////////////////////////////////////////////////////////////////////////////////////////////////
var PictureMap = map[string] string {
	"art1.png":        "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art1.png",
	"art2.png":        "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art2.png",
	"art3.png":        "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art3.png",
	"art4.png":        "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art4.png",
	"art5.png":        "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art5.png",
	"art6.png":        "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art6.png",
	"art7.png":        "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/art7.png",
	"item-001.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-001.jpg",
	"item-002.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-002.jpg",
	"item-003.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-003.jpg",
	"item-004.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-004.jpg",
	"item-005.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-005.jpg",
	"item-006.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-006.jpg",
	"item-007.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-007.jpg",
	"item-008.jpg":    "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/item-008.jpg",
	"people.gif":      "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/people.gif",
	"mad-fb.jpg":      "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/mad-fb.gif",
	"sample.png":      "https://raw.githubusercontent.com/ITPeople-Blockchain/auction/v0.6/art/artchaincode/sample.png",
}

/*
ChainCode kick-off
 */
func main (){

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Starting Item Auction Applicatoin ChainCode")

	// Start the shim --- running the fabric
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Error in starting the Item Auction Application")
	}
}

func downloadFile(filepath string, url string) (err error) {
//Create hte file
 out, err := os.Create(filepath)
 if err != nil {
 	return err
 }
 defer out.Close()

 //Get the data
 resp, err := http.Get(url)
 if err != nil {
 	return err
 }
 defer resp.Body.Close()

 //Write the body to file
 _,err = io.Copy(out,resp.Body)
 if err != nil {
 	return err
 }
 return nil
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
fmt.Println("Trade & Auction Appication Initizae")
var err error

	for key, value := range PictureMap {
		fmt.Println("\n Downloading Image '%s' from URL : %s", key ,value)
		err = downloadFile(key,value)
		if err != nil {
			fmt.Println(err)
			return shim.Error("Invoke : Invalid Function Name")
		}
	}
	fmt.Println("\n Init() Initalization complete")
	return  shim.Success(nil)

}

