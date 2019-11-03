package iorpc

//Ping Ping
func (as *InterfaceRPC) Ping(in string, out *string) error {
	*out = in
	return nil
}
