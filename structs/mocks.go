package structs

import (
	"encoding/xml"
)

var MockPricingAnswerVariant = &PricingAnswerVariant{
	FlightGroups: []PricingAnswerVariantFlightGroup{{
		Flight: []PricingAnswerVariantFlight{
			MockPricingAnswerVariantFlightOne,
			MockPricingAnswerVariantFlightTwo,
			MockPricingAnswerVariantFlightThree,
			MockPricingAnswerVariantFlightFour,
		},
	}},
	Directions: []PricingAnswerVariantDirection{MockPricingAnswerVariantDirectionOne, MockPricingAnswerVariantDirectionTwo},
	Total:      PricingAnswerVariantTotal{},
}

var MockPricingAnswerVariantWithBrand = &PricingAnswerVariant{
	FlightGroups: []PricingAnswerVariantFlightGroup{{
		Flight: []PricingAnswerVariantFlight{
			MockPricingAnswerVariantFlightOne,
			MockPricingAnswerVariantFlightTwo,
			MockPricingAnswerVariantFlightThree,
			MockPricingAnswerVariantFlightFour,
		},
	}},
	Directions: []PricingAnswerVariantDirection{MockPricingAnswerVariantDirectionOneWithBrand},
	Total:      PricingAnswerVariantTotal{},
}

var MockPricingAnswerVariantDirectionOne = PricingAnswerVariantDirection{
	Num:    1,
	Prices: []*PricingAnswerPrice{&MockPricingAnswerPrice},
}

var MockPricingAnswerVariantDirectionOneWithBrand = PricingAnswerVariantDirection{
	Num:    1,
	Prices: []*PricingAnswerPrice{&MockPricingAnswerPriceWithBrand},
}

var MockPricingAnswerVariantDirectionTwo = PricingAnswerVariantDirection{
	Num: 2,
	Prices: []*PricingAnswerPrice{{
		Upt:               PriceUpt{},
		Fare:              nil,
		Taxes:             nil,
		Vat:               nil,
		Baggage:           "",
		ValidatingAirline: "",
		OriginalPaxType:   "",
		Brand:             "",
		Total:             0,
		Currency:          "",
		PassengerID:       0,
		PaxType:           "CHD",
	}},
}

var MockPricingAnswerVariantFlightGroup = PricingAnswerVariantFlightGroup{
	BookingTimeLimit:      "",
	NeedLatinRegistration: false,
	ETPossible:            false,
	ETBlanks:              false,
	Flight:                []PricingAnswerVariantFlight{MockPricingAnswerVariantFlightOne},
}

var MockPricingAnswerVariantFlightOne = PricingAnswerVariantFlight{
	ID:         "1",
	Num:        1,
	Class:      "Y",
	SubClass:   "Y",
	BaseClass:  "B",
	Available:  0,
	SegmentNum: "",
	Cabin:      "N",
}

var MockPricingAnswerVariantFlightTwo = PricingAnswerVariantFlight{
	ID:         "5",
	Num:        1,
	Class:      "f",
	SubClass:   "q",
	BaseClass:  "B",
	Available:  0,
	SegmentNum: "",
	Cabin:      "N",
}

var MockPricingAnswerVariantFlightThree = PricingAnswerVariantFlight{
	ID:         "2",
	Num:        2,
	Class:      "Z",
	SubClass:   "H",
	BaseClass:  "J",
	Available:  0,
	SegmentNum: "",
	Cabin:      "J",
}

var MockPricingAnswerVariantFlightFour = PricingAnswerVariantFlight{
	ID:         "3",
	Num:        2,
	Class:      "Z",
	SubClass:   "H",
	BaseClass:  "J",
	Available:  0,
	SegmentNum: "",
	Cabin:      "J",
}

var MockPricingAnswerPrice = PricingAnswerPrice{
	Baggage:           "",
	ValidatingAirline: "",
	OriginalPaxType:   "",
	Currency:          "",
	PassengerID:       0,
	PaxType:           "ADT",
	IsRoundTrip:       false,
	FormPay:           "",
	Upt:               PriceUpt{},
	Fare:              nil,
	Taxes:             nil,
	Vat:               nil,
	Total:             0,
	UPT18CatText:      "",
}

var MockPricingAnswerPriceWithBrand = PricingAnswerPrice{
	Brand:             "BRAND TEST S7 N4",
	Baggage:           "",
	ValidatingAirline: "",
	OriginalPaxType:   "",
	Currency:          "",
	PassengerID:       0,
	PaxType:           "ADT",
	IsRoundTrip:       false,
	FormPay:           "",
	Upt:               PriceUpt{},
	Fare:              nil,
	Taxes:             nil,
	Vat:               nil,
	Total:             0,
	UPT18CatText:      "",
}

var MockDirectionOneFlightGroups = [][]PricingAnswerVariantFlight{
	{
		MockPricingAnswerVariantFlightOne,
		MockPricingAnswerVariantFlightTwo,
	},
}

var MockDirectionTwoFlightGroups = [][]PricingAnswerVariantFlight{
	{
		MockPricingAnswerVariantFlightOne,
		MockPricingAnswerVariantFlightTwo,
		MockPricingAnswerVariantFlightThree,
		MockPricingAnswerVariantFlightFour,
	},
}

var MockPricingResponse = PricingResponse{
	Answer: PricingAnswer{
		Pricing:          &MockPricingAnswerPricing,
		PricingMonobrand: nil,
	},
}

var MockPricingAnswerPricing = PricingAnswerPricing{
	Variants: []PricingAnswerVariant{
		{
			FlightGroups: []PricingAnswerVariantFlightGroup{
				{
					BookingTimeLimit:      "",
					NeedLatinRegistration: false,
					ETPossible:            false,
					ETBlanks:              false,
					Flight: []PricingAnswerVariantFlight{
						{
							ID:         "1",
							Num:        1,
							Class:      "",
							SubClass:   "",
							BaseClass:  "",
							Available:  0,
							SegmentNum: "",
							Cabin:      "",
						},
						{
							ID:         "2",
							Num:        2,
							Class:      "",
							SubClass:   "",
							BaseClass:  "",
							Available:  0,
							SegmentNum: "",
							Cabin:      "",
						},
					},
				},
			},
			Directions: []PricingAnswerVariantDirection{
				{
					Num:            1,
					RequestedBrand: "",
					Prices: []*PricingAnswerPrice{
						{

							Brand:             "",
							Baggage:           "",
							ValidatingAirline: "",
							OriginalPaxType:   "",
							Currency:          "",
							PassengerID:       0,
							PaxType:           "ADT",
							IsRoundTrip:       false,
							FormPay:           "",
							Upt:               PriceUpt{},
							Fare: &PricingAnswerPriceFare{
								Remark:      "",
								FareExpDate: "",
								Code:        "",
								BaseCode:    "",
								Total:       0,
							},
							Taxes:        nil,
							Vat:          nil,
							Total:        0,
							UPT18CatText: "",
						},
					},
				},
				{
					Num:            2,
					RequestedBrand: "",
					Prices: []*PricingAnswerPrice{
						{
							Brand:             "",
							Baggage:           "",
							ValidatingAirline: "",
							OriginalPaxType:   "",
							Currency:          "",
							PassengerID:       0,
							PaxType:           "ADT",
							IsRoundTrip:       false,
							FormPay:           "",
							Upt: PriceUpt{
								Idar1:     "",
								AddonIda:  "",
								Ntrip:     "",
								Nvr:       "",
								Ftnt:      "",
								CodeUpt:   "",
								Tariff:    "",
								MainAwk:   "",
								Cat:       "",
								Vcat:      "",
								City1:     "",
								City2:     "",
								Dport:     "",
								Aport:     "",
								BaseFare:  "",
								Iit:       "",
								Owrt:      "",
								Ddate:     "",
								Fdate:     "",
								DelivType: "",
								F0:        "",
								F1:        "",
								F2:        "",
								F3:        "",
								FlAwk:     "",
							},
							Fare: &PricingAnswerPriceFare{
								Remark:      "",
								FareExpDate: "",
								Code:        "",
								BaseCode:    "",
								Total:       0,
							},
							Taxes:        nil,
							Vat:          nil,
							Total:        0,
							UPT18CatText: "",
						},
					},
				},
			},
			Total: PricingAnswerVariantTotal{
				Currency: "",
				Total:    0,
			},
		},
	},
	Flights: []*PricingAnswerFlight{
		{
			ID:               "1",
			MarketingAirline: "",
			FlightNumber:     "",
			OperatingAirline: "",
			Flight:           "",
			Origin:           PricingAnswerLocation{},
			Destination:      PricingAnswerLocation{},
			Legs:             nil,
			DeptDate:         "10.03.20",
			ArrvDate:         "10.03.20",
			DeptTime:         "10:00",
			ArrvTime:         "10:00",
			Airplane:         "",
			FlightTime:       "",
			UPT18CatText:     "",
		},
		{
			ID:               "2",
			MarketingAirline: "",
			FlightNumber:     "",
			OperatingAirline: "",
			Flight:           "",
			Origin:           PricingAnswerLocation{},
			Destination:      PricingAnswerLocation{},
			Legs:             nil,
			DeptDate:         "13.03.20",
			ArrvDate:         "13.03.20",
			DeptTime:         "11:00",
			ArrvTime:         "11:00",
			Airplane:         "",
			FlightTime:       "",
			UPT18CatText:     "",
		},
	},
	Error: nil,
}

var MockIgnoredAirlines = []PricingRequestIgnoredAirline{
	{
		Name:   "v",
		Flight: "",
	},
}

var MockGetItinReceiptsDataResponse = GetItinReceiptsDataResponse{
	Answer: GetItinReceiptsDataAnswer{
		Answer: "",
		GetItinReceiptsData: GetItinReceiptsDataBody{
			Receipts: &GetItinReceiptsDataAnswerReceipts{
				TicketForm: []TicketForm{
					{
						IssueDate:       "",
						CRTime:          "",
						PassengerID:     "441",
						NameOfPassenger: "One",
						DocOfPassenger:  "",
						Total:           "",
						SerialNumber:    "",
					},
					{
						IssueDate:       "",
						CRTime:          "",
						PassengerID:     "222",
						NameOfPassenger: "Two",
						DocOfPassenger:  "",
						Total:           "",
						SerialNumber:    "",
					},
					{
						IssueDate:       "",
						CRTime:          "",
						PassengerID:     "333",
						NameOfPassenger: "Two",
						DocOfPassenger:  "",
						Total:           "",
						SerialNumber:    "",
					},
				},
			},
			Error: nil,
		},
	},
	XMLName: xml.Name{},
}
