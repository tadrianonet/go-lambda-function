package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// Estrutura de entrada
type Input struct {
	InitialAmount float64 `json:"initial_amount"` // Valor inicial
	Rate          float64 `json:"rate"`           // Taxa de juros em porcentagem
	Periods       int     `json:"periods"`        // Número de períodos
}

// Estrutura de saída
type Output struct {
	FinalAmount float64 `json:"final_amount"` // Valor final do investimento
}

// handler processa o cálculo do investimento
func handler(ctx context.Context, input Input) (Output, error) {
	// Calcular o valor final: FV = PV * (1 + i)^n
	finalAmount := input.InitialAmount * pow(1+(input.Rate/100), input.Periods)

	// Criar o output
	output := Output{
		FinalAmount: finalAmount,
	}

	return output, nil
}

// Função auxiliar para cálculo de potência
func pow(base float64, exp int) float64 {
	result := 1.0
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	lambda.Start(handler)
}
