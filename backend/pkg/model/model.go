package model

type SayHelloResponse struct {
	Message string `json:"message"`
}
type InfoResponse struct {
	AuthenticationFlowType string `json:"authenticationFlowType"`
	Region                 string `json:"region"`
	UserPoolId             string `json:"userPoolId"`
	UserPoolWebClientId    string `json:"userPoolWebClientId"`
}
type ApiError struct {
	ErrorMessage string `json:"errorMessage"`
}
