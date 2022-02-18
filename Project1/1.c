
#define  _CRT_SECURE_NO_WARNINGS 1


#include <stdio.h>
//
//int main()
//{
//	int a = 4;
//	printf("%d \n", a);
//	return 0;
//}

//#include <stdio.h>
//
//int main()
//{
//	int a = 20;
//	int* d = &a; // d 是一个地址变量。叫指针变量。  int*   int 对应 a 变量的类型  * 表明a 是是指针变量。
//	//printf("%p \n", d);  // %p 打印 地址/指针 
//	*d = 2021; // *解应用操作符/ 间接访问操作符。   修改对应的地址里面的值。
//	printf("%d \n", a); //打印 a  输出 2021.
//	return 0;
//}
//
//int main()
//
//{
//	double cc = 3.141546;
//	double* ff = &cc;
//	*ff = 7.834525;
//	printf("%lf\n", cc);  //double 类型%lf 打印
//	printf("%lf \n", *ff);  // 直接打印 *ff
//	printf("%d \n", sizeof(ff)); //  所有的指针类型 x86  4个字节  x64 8个字节  // 不管变量类型的长度
//	return 0;
//}


//结构体  
 //char  int double  人 或者书  复杂对象     名字身高 年龄 号码
//复杂对象 描述结构体
//自己创造出来的一种类型
//
//#include <string.h>
//
//	struct ren       //struct  结构体关键字 
//	{
//		char name[20];
//		short nianlin;
//		char  xinbei[20];
//	};
//
//	int main()
//	{
//		struct ren NO1 = { "张三",22,"男" }; //第一种写法
//		//NO1.nianlin = 54; // 修改年龄。  变量可以直接赋值  
//		// 数组不能赋值  
//		strcpy(NO1.name, "张狗子"); 
//		// strcpy -- string copy  字符串拷贝。  库函数  添加头文件。
//		
//		struct ren* DD = &NO1;  //取地址 给DD
//
//		//printf("%s\n", (*DD).name);
//		//printf("%d\n", (*DD).nianlin); // 解应用   在打印
//		//printf("%s\n", (*DD).xinbei);
//
//
//		printf("%s \n" ,DD->name);
//		printf("%d \n", DD->nianlin);
//
//		//################### . 结构体变量 . 找到成员
//		//###################  -> 结构体指针-> 找到成员。
//
//		//printf("谁%s\n", NO1.name);
//		//printf("年龄%d\n", NO1.nianlin);
//		//printf("性别 %s \n", NO1.xinbei);
//
//		return 0;
//	}

//struct stu {
//	char name[40]; //姓名
//	int mum; //编号
//	int age; // 年龄
//	char groups[30]; //组
//}class[2] = {
//	{"张狗子",001,43, "田田"},
//	{"王柳",002,53,'过节'}
//};
//int main()
//{
//	printf("谁%s \n", class->name); // 可以打印 默认名第一行 class
//	return 0;
//}


// c语言是一门  结构化的 程序设计语言
//1. 顺序结构 
//2.选择结构   分支语句  if   switch  
//3. 循坏结构   while   for   dowhile


//选择结构  //   分支语句   
//循坏结构  // 循坏语句
 
//语句  在c语言中由一个分号; 隔开的就是一条语句
//; 是语句  是一个空语句

//分支语句  选择语句


//if

//#include <stdio.h>

//int main()
//{
//	//;// 是语句  空语句
//
//	int age = 45;
//	/*if (age < 18)
//		printf("未成年\n");*/   // 单分支
//	if (age < 18)
//		printf("未成年\n"); // 双分支   if 后面跟 表达式
//	else
//		printf("成年了\n");
//
//	return 0;
//}

//int main()
//{
//	int age = 124;
//	if (age < 18)
//		printf("未成年\n");
//	//else if (18 <= age < 28) // 这中逻辑错误。。
//	else if (age >= 18 && age < 28)
//		printf("青年\n");
//	else if (age >= 28 && age < 50)
//		printf("中年\n");
//	else if (age >= 50 && age < 90)
//		printf("老年\n");
//	else
//		printf("不说了 您牛逼!!! \n");
//
//	return 0;
//}

//int main()
//{
//	int age = 54;
//	if (age < 18)
//		printf("未成年\n");
//	else   // 嵌套 if  else   牛逼了 可以多层嵌套
//	{
//		if (age >= 18 && age < 28)
//			printf("青年\n");
//		else if (age >= 28 && age < 50)
//			printf("中年\n");
//		else if (age >= 50 && age < 90)
//			printf("老年\n");
//		else
//			printf("不说了 您牛逼!!! \n");
//
//	}
//	
//		return 0;
//}


// c 语言中  0 表示假   非0 表示真  （-1  也表示真）
//如果表达式成立执行多条语句  使用代码块
//
//int main()
//{
//	int a = 45;
//	
//	if (a < 18)
//	{
//		printf("你未成年\n");  //   一个大括号 就是一个代码块
//		printf("不能早恋吆！ \n");//  if  else  控制多条语句 必须使用代码块
//	}
//	else
//	{
//		printf("18不禁止\n");
//		printf("你懂得 \n");
//	}
//	return 0;
//}


//else 悬空问题
//int main()
//{
//	// c 语言中  0 表示假   非0 表示真
//	int a = 0;
//	int b = 2;
//	if (a == 1) // 为假  不执行。 // else 和离得最近未匹配的if 匹配。
//
//		if (b == 2)
//			printf("nihao\n");   // else 和离得最近未匹配的if 匹配。
//		else
//			printf("你好\n"); // else 和离得最近未匹配的if 匹配。
//
//
//	return 0;
//}


//int main()
//{
//	// c 语言中  0 表示假   非0 表示真
//	int a = 0;
//	int b = 2;
//	if (a == 1) // 为假  不执行。 // else 和离得最近未匹配的if 匹配。
//
//	{
//		if (b == 2)
//			printf("nihao\n");   // else 和离得最近未匹配的if 匹配。
//	}
//		else
//			printf("你好\n"); // else 和离得最近未匹配的if 匹配。
//
//
//	return 0;
//}
//
//int main()
//
//{
//	int mun = 4;
//	if (mun = 7)  //  =  是赋值       ==  是判断  ！！！！
//	if (7 == mun) // 建议判断 常量写左边
//	{
//	
//		printf("你猜我会不会被打印 \n");
//	}
//
//
//	return 0;
//}


// 输出1-100的奇数。

int main()

{
	int a = 0;
	while (a<=100)
	{
		if (a % 2 == 1)
			printf("%d  ", a);
		a++;
	}
	

	return 0;
}


//int main()
//
//{
//	int a = 1;
//	while (a <= 100)
//	{
//		printf("%d ", a);
//		a += 2;
//	}
//
//
//	return 0;
//}

//switch 语句   也是一种分支语句 常常用于多分支的情况

