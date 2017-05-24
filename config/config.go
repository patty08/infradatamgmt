package config

var Config = struct {
	Elasticsearch string
	ElasticAuth bool
	ElasticUser string
	ElasticPassword string
	Kibana string
	Agent []string
	Client string
}{}