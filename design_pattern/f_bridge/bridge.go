package f_bridge

/*
桥接模式
思想：组合优于继承。继承方式会导致实现呈指数级增加（M*N），而组合只需要常量级（M+N）
*/

// IMsgSender 消息通道接口
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 发送邮件，可能还有企微、电话等
type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{
		emails: emails,
	}
}

func (e *EmailMsgSender) Send(msg string) error {
	return nil
}

// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification error级别通知，以后可能还有warn、info
type ErrorNotification struct {
	// 这里通过组合的方式组合消息通道
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{
		sender: sender,
	}
}

func (en *ErrorNotification) Notify(msg string) error {
	return en.sender.Send(msg)
}
