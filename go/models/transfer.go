package models

import "github.com/google/uuid"

// Structure representing a player on the transfer list.
type Transfer struct {
	team      *Team            // Team of the player.
	id        uuid.UUID        // Id of player in the team.
	time      int              // Time until player gets removed from the list.
	fee, wage [QUALITY_END]int // Estimated fees and wages for different scout qualities.
	offers    []*TransferOffer // Offers for the player.
}

type TransferOffer struct {
	team      *Team               // The team that makes the offer.
	fee, wage int                 // Transfer fee and wage offer.
	status    TransferOfferStatus // Whether the offer got accepted or rejected etc.
}

type TransferOfferStatus int

const (
	TRANSFER_OFFER_NOT_CONSIDERED TransferOfferStatus = iota
	TRANSFER_OFFER_ACCEPTED
	TRANSFER_OFFER_REJECTED
	TRANSFER_OFFER_REJECTED2
	TRANSFER_OFFER_END
)
