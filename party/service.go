package party

import "github.com/solkn/Voting_System/entity"

type PartyServices interface {

	Parties()([]entity.Party,[]error)
	Party(id uint)(*entity.Party,[]error)
	StoreParty(party *entity.Party)(*entity.Party,[]error)
	UpdateParty(party *entity.Party)(*entity.Party,[]error)
	DeleteParty(id uint)(*entity.Party,[]error)
}