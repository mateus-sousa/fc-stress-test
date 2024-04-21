/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mateus-sousa/fc-stress-test/internal"
	"log"

	"github.com/spf13/cobra"
)

// stressCmd represents the stress command
var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Sistema de Stress test",
	Long: `Um sistema CLI em Go para realizar testes de carga em um serviço web. 
		O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt64("requests")
		concurrency, _ := cmd.Flags().GetInt64("concurrency")
		if url == "" {
			fmt.Println("Parametro url é obrigatório.")
			return
		}
		if requests == 0 {
			fmt.Println("Parametro requests é obrigatório.")
			return
		}
		if concurrency == 0 {
			fmt.Println("Parametro concurrency é obrigatório.")
			return
		}
		stressTestUseCase := internal.NewStressTestUseCase()
		report, err := stressTestUseCase.Exec(url, requests, concurrency)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("Relatório final")
		fmt.Printf("Tempo total gasto na execução: %v \n", report.TotalTimeExec)
		fmt.Printf("Quantidade total de requests realizados: %v \n", report.TotalAmountRequests)
		fmt.Printf("Quantidade de requests com status HTTP 200: %v \n", report.TotalAmountHTTPStatusOk)
		fmt.Printf("Distribuição de outros códigos de status HTTP: %v \n", getStatusCodeFailList(report.AllHTTPSStatus))
	},
}

func getStatusCodeFailList(statusCodeList map[int]int64) string {
	delete(statusCodeList, 200)
	result := ""
	for k, v := range statusCodeList {
		result += fmt.Sprintf("%v: %v; ", k, v)
	}
	return result
}

func init() {
	rootCmd.AddCommand(stressCmd)
	stressCmd.Flags().StringP("url", "u", "", "Url do serviço a ser testado")
	stressCmd.Flags().Int64P("requests", "r", 0, "Número total de requests")
	stressCmd.Flags().Int64P("concurrency", "c", 0, "Número de chamadas simultâneas")
}
