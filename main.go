package main

import (
	"fmt"
	"github.com/hothero/generated-protos/listing"
	"github.com/hothero/generated-protos/offer"
)

func main() {
	l := new(listing.Listing)
	fmt.Println(l.Price)

	o := new(offer.Offer)
	fmt.Println(o.LatestMessage)

	return
}
