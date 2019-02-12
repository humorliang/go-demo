package proxy

import "errors"

/*
 代理模式
	为其他对象提供一种代理以控制对这个对象的访问。
	在某些情况下，一个对象不适合或者不能直接引用另一个对象，而代理对象可以在客户端和目标对象之间起到中介的作用。
	比如你喜欢一个妹子， 不好意思跟人家开口， 这时候你可能就通过她舍友来表达你的爱慕了。
     优点： 1、职责清晰。 2、高扩展性。 3、智能化。
     注意事项：
          1、和适配器模式的区别：适配器模式主要改变所考虑对象的接口，
             而代理模式不能改变所代理类的接口。
          2、和装饰器模式的区别：装饰器模式为了增强功能，而代理模式
             是为了加以控制。
代理模式共有4个角色， 其中一个是抽象的：
1. Client 就是上面那个段子中的你， 你是行为的主导者。
2. Subject 是代理人和被代理的抽象接口
3. RealSubject 被代理的对象， 也就是上面的妹子
4. Proxy 代理者， 对应上面的妹子室友
*/

//代理人和被代理的抽象接口
type Device interface {
	Read() ([]byte, error)
	Writer(msg []byte) error
}

//被代理的对象
type HardDisk struct {
	Storage []byte
}

func (hd *HardDisk) Read() ([]byte, error) {
	return hd.Storage, nil
}
func (hd *HardDisk) Writer(msg []byte) error {
	hd.Storage = msg
	return nil
}

//代理者
type HardDiskProxy struct {
	OpId string
	Hd   *HardDisk //被代理对象的指针类型
}

//代理者处理的方法（她室友给你探探口风）
//权限判断
func (hp *HardDiskProxy) Permission(tag string) bool {
	if hp.OpId == "admin" {
		return true
	}
	if hp.OpId == "reader" && tag == "read" {
		return true
	}
	if hp.OpId == "writer" && tag == "write" {
		return true
	}
	return false
}

func (hp *HardDiskProxy) Read() ([]byte, error) {
	if !hp.Permission("read") {
		return nil, errors.New("You don't have permission read ")
	}
	return hp.Hd.Read()
}

func (hp *HardDiskProxy) Writer(msg []byte) error {
	if !hp.Permission("write") {
		return errors.New("You don't have permission write ")
	}
	return hp.Hd.Writer(msg)
}

