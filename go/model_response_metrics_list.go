/*
 * APIs Admin do Open Finance Brasil
 *
 * As API's administrativas são recursos que podem ser consumidos apenas pelo diretório para avaliação e controle da qualidade dos serviços fornecidos pelas instituições financeiras
 *
 * API version: 2.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseMetricsList struct {

	Data []ResponseMetricsListData `json:"data"`

	Links *Links `json:"links"`

	Meta *Meta `json:"meta,omitempty"`
}
