package runner

func Payload() []string {
	payloads := []string{
		"constructor[prototype][I1Younes]%dTest4me",
		"__proto__.I1Younes%3dTest4me",
		"constructor.prototype.I1Younes%sdTest4me",
		"__proto__[I1Younes]%3dTest4me",
	}
	return payloads
}
