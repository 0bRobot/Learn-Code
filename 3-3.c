#define _CRT_SECURE_NO_WARNINGS 1
#include <stdio.h>
#include <math.h>

// 9 9 乘法表

//int main()
//{	
//	int a, b;
//	for (a = 1; a <=9; a++)
//	{
//		for (b = 1; b <=a; b++)  // a打印几行b就打印几行  a 是1  就打印一行   a是2 就打印2行
//		{
//			printf("%d*%d=%-3d ",b,a,a*b); // %3d  表示打印3位  不够用空格补齐
//											// %-3d  三位左对齐
//		}
//		printf("\n"); // 打印一行就换行
//	}
//	return 0;
//}

//int main()
// 打印100 以内的奇数

//{
//	int a=1;
//	while (a<100)
//	{
//		printf("%d ", a);
//		a += 2;
//	}
//	return 0;
//}


//void Menu()
//{// 方法一 ###############################################################
//	printf("###############################\n");
//	printf("######------求素数------#######\n");
//	printf("###############################\n");
//	printf("==========输入数值的范围=======\n");
//	printf("！！！要求第二个数大于第一个数!\n");
//	printf("\n");
//	printf("-------------------------------\n");
//}
//int main()
//{ //求对应范围的素数 间的素数  >>>只能被1和本身整除
//	
//	Menu();
//	
//	int i, u, mun = 0;
//
//	printf("[空格区分]请输入_: ");
//
//	scanf("%d %d", &i, &u);
//      printf("\n");
//	for (i; i <= u; i++)  //循环生成 100 - 200的数字
//	
//	{// 判断 i是否为素数
//		int j = 0;
//		for (j = 2; j <=sqrt(i); j++)
//		{
//			if (i % j == 0)
//			{
//				break;
//			}
//		}
//		if (j > sqrt(i))
//		{
//			printf("%-2d  ", i);
//			mun++;
//		}
//		
//	}
//	printf("\n");
//	printf("-------------------------\n");
//	printf("这个范围一共有%d 个素数\n", mun);
//	return 0;
//}





void Menu()
{// 函数实现  scanf 判断实现
	printf("===============================\n");
	printf("######------求素数------#######\n");
	printf("###############################\n");
	printf("\n");
	printf("==========输入数值的范围=======\n");
	printf("！！！要求第二个数大于第一个数!\n");
	printf("\n");
	printf("-------------------------------\n");
}

void  qSS(int i,int u)
{
	int mun =0;
	for (i; i <= u; i++)  //生成 对应范围的 的数字

	{// 判断 i是否为素数      a和b中至少有一个数小于等于 开平方 c
		int j = 0;
		for (j = 2; j <= sqrt(i); j++)   //sqrt()  开平方函数     需要用到 math函数
		{
			if (i % j == 0)   // 能被整除
			{
				break;
			}
		}
		if (j > sqrt(i))
		{
			printf("%-2d  ", i);
			mun++;  //计数
		}

	}
	printf("\n");
	printf("-------------------------\n");
	printf("这个范围一共有%d 个素数\n", mun);
}
int main()
{ //求对应范围的素数 间的素数  >>>只能被1和本身整除

	Menu();  //打印函数

	int iSTART,iEND;
	printf("[空格区分]请输入_: ");


	while (1)
	{
		if ((scanf("%d %d", &iSTART, &iEND)) == 2)   //判断scang 获取值得正确性
		{
			printf("\n"); 
			qSS(iSTART, iEND);  //  判断素数函数   
			break;
		}
		else
		{
			printf("\n");
			printf("大哥看一看输错了！！\n");
			printf("++++++++++++++++++++++++\n");
			printf("\n");
			printf("[空格区分]请输入_: ");
			int c;
			while ((c = getchar()) != '\n' && c != EOF);  // 清空缓冲区   

		}
	}

	
	
	return 0;
}