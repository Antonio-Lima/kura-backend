package model

import (
	"time"

	"github.com/google/uuid"
)

type CategoryInput struct {
	Category string `json:"category" binding:"required"`
	Color    string `json:"color" binding:"required,hexcolor"`
	Icon     string `json:"icon" binding:"required"`
}

type Category struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	UserID uuid.UUID `json:"-"`
	User   User      `json:"user" gorm:"foreign_key:UserId"`
	CategoryInput
}

const (
	// Essenciais e Renda
	IconSalary     = "Banknote"
	IconInvestment = "TrendingUp"
	IconSavings    = "PiggyBank"
	IconBonus      = "Gift"
	IconWallet     = "Wallet"

	// Alimentação
	IconFood     = "Utensils"
	IconCoffee   = "Coffee"
	IconGrocery  = "ShoppingCart"
	IconFastFood = "Pizza"
	IconBar      = "Beer"

	// Casa e Contas
	IconHome        = "Home"
	IconElectricity = "Zap"
	IconWater       = "Droplets"
	IconInternet    = "Wifi"
	IconRent        = "Key"
	IconMaintenance = "Wrench"

	// Transporte
	IconCar  = "Car"
	IconFuel = "Fuel"
	IconBus  = "Bus"
	IconBike = "Bike"
	IconUber = "Navigation"

	// Saúde e Bem-estar
	IconHealth   = "HeartPulse"
	IconPharmacy = "Pill"
	IconGym      = "Dumbbell"
	IconDoctor   = "Stethoscope"

	// Lazer e Estilo de Vida
	IconGame   = "Gamepad2"
	IconMovie  = "Film"
	IconTravel = "Plane"
	IconHotel  = "Bed"
	IconMusic  = "Music"
	IconParty  = "GlassWater"

	// Compras
	IconShopping = "ShoppingBag"
	IconShirt    = "Shirt"
	IconGift     = "Package"
	IconDevice   = "Smartphone"

	// Educação e Trabalho
	IconEducation = "GraduationCap"
	IconBook      = "Book"
	IconWork      = "Briefcase"
	IconPrinter   = "Printer"

	// Outros
	IconPet          = "Dog"
	IconChild        = "Baby"
	IconTax          = "Receipt"
	IconInsurance    = "ShieldCheck"
	IconSubscription = "Calendar"
	IconCreditCard   = "CreditCard"
	IconCharity      = "Heart"
	IconOther        = "CircleDot"
)
