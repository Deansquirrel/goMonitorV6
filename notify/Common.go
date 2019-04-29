package notify

//发送消息
// - id TaskId
// - msg 待发送文本
// return
// - total 待通知对象列表
// - success 通知成功对象数量
// - fail 通知失败对象数量
// - sendErr 发送过程遇到错误
// - notifyErrMap 通知对象时遇到错误 [TaskId]error
func SendMsg(id string, msg string) (total int, success, fail int, sendErr error, notifyErrMap map[string]error) {
	//notifyErrMap = make(map[string]error)
	//n := notify{}
	//list, err := n.GetNotifyList(id)
	//if err != nil {
	//	sendErr = errors.New(fmt.Sprintf("get notify list error: %s", err.Error()))
	//	return
	//}
	//total = len(list)
	//for _, notify := range list {
	//	notifyErr := notify.SendMsg(msg)
	//	if notifyErr != nil {
	//		notifyErrMap[notify.GetId()] = notifyErr
	//		fail = fail + 1
	//	} else {
	//		success = success + 1
	//	}
	//}
	return
}
