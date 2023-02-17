package tx

import "context"

// Transaction 事务接口 在data层实现该接口 以在service层使用事务
// 在data层写事务会经常涉及多表操作使业务逻辑下沉到data层 使得存储库职责混乱 因此data层与业务逻辑耦合 不利于后期维护
// service层的职责类似Facade 将DTO装换为领域模型 同时调用多个usecase协作 因此事务写在该层 非常便于将多个usecase方法组合在一个事务中
type Transaction interface {
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}
