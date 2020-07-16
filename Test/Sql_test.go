package Test

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Model"
	"fmt"
	"testing"
)

// 新增
func TestSqlInsert(t *testing.T) {
	// 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton"

	if Model.CreateTestFactory("").InsertData() {
		fmt.Println("数据插入成功")
	} else {
		t.Errorf("数据插入操作，单元测试失败")
	}
}

// 查询（多条）
func TestSqlSelect(t *testing.T) {
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton"

	list := Model.CreateTestFactory("").QueryData()
	if list != nil {
		for index, item := range list {
			fmt.Printf("%d, %s,%d, %d, %s, %s\n", index, item.Name, item.Age, item.Sex, item.Addr, item.Remark)
		}
	} else {
		t.Errorf("数据查询操作，单元测试失败")
	}
}

// 查询（单条）
func TestSqlSelectOne(t *testing.T) {
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton"

	oneList := Model.CreateTestFactory("").QueryRowData()
	if oneList == nil {
		t.Errorf("单元测试：单条数据查询失败")
	} else {
		fmt.Printf("%#+v\n", *oneList)
	}
}

// 测试提交事务的操作
func TestSqlTransAction(t *testing.T) {
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton"
	// 修改以下函数的参数，测试事务的提交（true）与回滚（false）
	if Model.CreateTestFactory("").TransAction(true) {
		fmt.Println("数据插入成功(提交事务操作)")
	} else {
		t.Errorf("数据插入（提交事务操作），单元测试失败")
	}
}

// 测试回滚事务的操作
func TestSqlTransAction2(t *testing.T) {
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton"
	// 参数 true 表示 提交事务；  false 表示 回滚事务
	if Model.CreateTestFactory("").TransAction(true) {
		fmt.Println("数据插入成功(回滚事务操作)")
	} else {
		t.Errorf("数据插入（回滚事务操作），单元测试失败！")
	}
}

// 批量插入数据的正确姿势
func TestSqlInsertMultiple(t *testing.T) {
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton"

	if Model.CreateTestFactory("").InsertDataMultiple() {
		fmt.Println("批量插入数据OK")
	} else {
		t.Errorf("批量插入数据，单元测试失败！")
	}
}

// 批量插入数据的错误姿势
func TestSqlInsertMultipleError(t *testing.T) {
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton"

	if Model.CreateTestFactory("").InsertDataMultipleErrorMethod() {
		fmt.Println("批量插入数据OK")
	} else {
		t.Errorf("批量插入数据，单元测试失败！")
	}
}
