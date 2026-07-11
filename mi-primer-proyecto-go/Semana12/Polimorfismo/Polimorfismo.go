package main

import (
	"errors"
	"fmt"
)

type PaymentProcessor interface {
	Process(amount float64) error
	GetFree() float64
}

type CreditCardProcessor struct {
	CardNumer string
	FeeRate   float64
}

func (cc CreditCardProcessor) Process(amount float64) error {

	if amount <= 0 {
		return errors.New("monto invalido")
	}
	fmt.Printf("Procesando $%.2f con tarjeta ****%s\n", amount, cc.CardNumer[len(cc.CardNumer)-4:])
	return nil
}

func (cc CreditCardProcessor) GetFree() float64 {
	return cc.FeeRate
}

type PayPalProcessor struct {
	Email string
}

func (pp PayPalProcessor) Process(amount float64) error {
	if amount <= 0 {
		return errors.New("Monto invalido")
	}
	fmt.Printf("Procesando $%.2f via PayPal (%s)\n ", amount, pp.Email)
	return nil
}

func (pp PayPalProcessor) GetFree() float64 {
	return 0.029 // 2.9%
}

type CryptoProcessor struct {
	WalletAddress string
	Currency      string
}

func (cp CryptoProcessor) Process(amount float64) error {
	if amount <= 0 {
		return errors.New("Monto invalido")
	}
	fmt.Printf("Procensando $%.2f en %s (wallet: %s...)\n", amount, cp.Currency, cp.WalletAddress[:10])
	return nil
}

func (cp CryptoProcessor) GetFree() float64 {
	return 0.01 //1%
}

func ProcessoOrder(processor PaymentProcessor, amount float64) error {
	fee := amount * processor.GetFree()
	total := amount + fee

	fmt.Printf("Procesando orden por $%.2f (+ $%.2f fee) = $%.2f total\n", amount, fee, total)
	return processor.Process(total)
}

func ProcessWithBestRate(amount float64, processors []PaymentProcessor) error {
	if len(processors) == 0 {
		return errors.New("no hay procesadores disponibles")
	}

	//Procesador con menor comision
	bestProcessor := processors[0]
	lowestFee := bestProcessor.GetFree()

	for _, processors := range processors[1:] {
		if processors.GetFree() < lowestFee {
			bestProcessor = processors
			lowestFee = processors.GetFree()
		}
	}

	fmt.Printf("Selecionado el mejor proceso con %.1f%% de comision\n ", lowestFee*100)
	return ProcessoOrder(bestProcessor, amount)
}

func main() {
	crediCard := CreditCardProcessor{
		CardNumer: "1234567890123456",
		FeeRate:   0.035, //3.5%
	}
	paypal := PayPalProcessor{
		Email: "example@paypal.com",
	}

	crypto := CryptoProcessor{
		WalletAddress: "1Q2W3E4R5T6Y7U8I9O",
		Currency:      "BTC",
	}

	//polimorfismo

	processors := []PaymentProcessor{crediCard, paypal, crypto}

	fmt.Println("====Procesando con el mejor Rate===")
	ProcessWithBestRate(100.0, processors)
	for _, processor := range processors {
		ProcessoOrder(processor, 50.0)
		fmt.Println()
	}

	fmt.Scanln()
}
