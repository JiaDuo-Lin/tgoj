package global

import (
	"fmt"
	"log"
	"os"
	"testing"
	"tgoj/server/model"
)

func TestGlobalDB(t *testing.T) {
	err := DB.Set(
		"gorm:table_options",
		"ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
		AutoMigrate(
			&model.Question{},
			&model.User{},
		)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}


	q1 := model.Question{
		Title: "两数之和",
		Description: "编写程序，实现两个整数相加的功能，并输出相加的结果。",
		Level: "简单",
		TestData: "2 10\n8 20\n-90 90\n0 0\n20 -30\n9999 1\n-10000 1\n",
		TestAnswer: "12\n28\n0\n0\n-10\n10000\n-9999\n",
		Example: "示例1\n    输入:1 2\n    输出:3",
		Tags: "入门级",
		MemoryLimit: 20<<20,
		TimeLimit: 1.0,
	}

	q2 := model.Question{
		Title: "整数反转",
		Description: "给你一个 32 位的有符号整数 x ，返回 x 中每位上的数字反转后的结果。\n如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。假设环境不允许存储 64 位整数（有符号或无符号）。",
		Level: "简单",
		TestData: "123\n0\n-123\n1234567899\n",
		TestAnswer: "321\n0\n-321\n0\n",
		Example: "示例1\n    输入:123\n    输出:321",
		Tags: "入门级",
		MemoryLimit: 20<<20,
		TimeLimit: 1.0,
	}

	result := DB.Create(&q1)
	fmt.Println(*result)
	result = DB.Create(&q2)
	fmt.Println(*result)
}
