package api

func (o *MailcowResponseArray) GetTransportId() (*string, error) {
	if !o.HasFinalMsgItem(0) || !o.HasFinalMsgItem(2) {
		return nil, errors.New(fmt.Sprint("msg error: ", o.GetFinalMsgs()))
	}
	receipt := *o.GetFinalMsgItem(0)
	if receipt != "relayhost_added" {
		return nil, errors.New(fmt.Sprint("msg error: ", receipt))
	}
	return o.GetFinalMsgItem(2), nil
}
