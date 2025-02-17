/*
 * APIs Admin do Open Finance Brasil
 *
 * As API's administrativas são recursos que podem ser consumidos apenas pelo diretório para avaliação e controle da qualidade dos serviços fornecidos pelas instituições financeiras
 *
 * API version: 2.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RejectionMetrics struct {
	// Número de chamadas rejeitadas no dia atual.
	CurrentDay int32 `json:"currentDay"`
	// Número de chamadas rejeitadas nos dias anteriores. O primeiro item do array é referente ao dia anterior ao corrente, e assim por diante. Devem ser retornados no máximo sete dias caso estejam disponíveis.​
	PreviousDays []int32 `json:"previousDays"`
}
