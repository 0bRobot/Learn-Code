#define _CRT_SECURE_NO_WARNINGS 1
// C语言入门编程100题

 NO.01
# include <stdio.h> 

// 1 2 3 4 个数字组成多少个互不相同的无重复的三位数字   分别是多少
int main()
{
	int i, j, k;  // 定义三个变量
	int a = 0;  //计数
	printf("\n"); //打印空行
	for (i = 1; i <= 4; i++)  // 以下为三重循环
	{
		for (j = 1; j <= 4; j++)
		{
			for (k = 1; k <= 4; k++)
			{
				if (i != k && i != j && j != k)  // 确保i、j、k三位互不相同
				{
					printf("%d %d %d\n",i ,j,k);
					a++;  // 打印一次 自增1
					
				}
			}
		}
	}
	printf("一共有%d组\n", a);  // 打印总数
	return 0;
}

// NO.02
#include <stdio.h>

 //企业发放的奖金根据利润提成。 从键盘输入当月利润I，求应发放奖金总数？
//利润(I)低于或等于10万元时，奖金可提10% ；
//利润高于10万元，低于20万元时，低于10万元的部分按10% 提成，高于10万元的部分，可提成7.5 % ；
//20万到40万之间时，高于20万元的部分，可提成5 % ；
//40万到60万之间时高于40万元的部分，可提成3 % ；
//60万到100万之间时，高于60万元的部分，可提成1.5 % ；
//高于100万元时，超过100万元的部分按1 % 提成

int main()
{
	double I;  //当月近利润
	double  Bouns, Bouns1, Bouns2, Bouns4, Bouns6, Bouns10;  // 6个不同的类别的提成策略
	Bouns = 1;
	printf("输入你当月的利润：<");
	scanf("%lf", &I);
	Bouns1 = 100000 * 0.1;  //  低于等于 10万 的奖金
	Bouns2 = Bouns1 + 100000 * 0.075;  // 高于10万 低于20万的奖金  =  低于10万的奖金 +  高出10万部分的奖金
	Bouns4 = Bouns2 + 200000 * 0.05;
	Bouns6 = Bouns4 + 200000 * 0.03;
	Bouns10 = Bouns6 + 400000 * 0.015;
	
	if (I <= 100000)
	{
		Bouns = I * 0.1;  //利润低于10万的  奖金=  利润*10%
	}
	else if (I <= 200000)
	{
		Bouns = Bouns1 + (I - 100000) * 0.075;  // 低于等于10万的奖金加上    高出的部分 提成 7.5%  
		
	}
	else if (I <= 400000)
	{
		Bouns = Bouns2 + (I - 200000) * 0.05;
	}
	else if (I < 600000)
	{
		Bouns = Bouns4 + (I - 400000) * 0.03;
	}
	else if (I <= 1000000)
	{
		Bouns = Bouns6 + (I - 600000) * 0.015;
	}
	else if (I > 1000000)
	{
		Bouns = Bouns10 + (I - 1000000) * 0.01;
	}
	printf("你的奖金为：%lf\n",Bouns);
	return 0;
}


//NO.3

#include <stdio.h>

//一个整数，它加上100后是一个完全平方数，再加上168又是一个完全平方数，请问该数是多少？

int main()
{
	int i, j, m, n, x;
	for (i = 1; i < 168 / 2 + 1; i++)
	{
		if (168 % i == 0)
		{
			j = 168 / i;
			if (i > j && (i + j) % 2 == 0 && (i - j) % 2 == 0)
			{
				m = (i + j) / 2;
				n = (i - j) / 2;
				x = n * n - 100;
				printf("%d + 100 = %d * %d\n",x,n,n);
				printf("%d + 268 = %d * %d\n",x,m,m);
			}
		}
	}
	return 0;
}


//NO.4
#include<stdio.h>

//输入某年某月某日，判断这一天是这一年的第几天？
int main()
{
	int day, month, year,leap;
	int sum =0;
	printf("\n请输入年、月、日,格式为：年 月 日（2015 12 10）\n");
	scanf("%d %d %d",&year,&month,&day); 
	switch (month)
	{
	case 1:
		sum = 0;   // 1月不计算
		break;
	case 2:
		sum = 31;  // 2月 要增加 1月的天数
		break;
	case 3:        // 3月  1月和2月的天数之和；
		sum = 59;
		break;
	case 4:
		sum = 90;
		break;
	case 5:
		sum = 120;
		break;
	case 6:
		sum = 151;
		break;
	case 7:
		sum = 181;
		break;
	case 8:
		sum = 212;
		break;
	case 9:
		sum = 243;
		break;
	case 10:
		sum = 273;
		break;
	case 11:
		sum = 304;
		break;
	case 12:
		sum = 334;
		break;
	default:
		printf("date Error\n");
		break;
	}
	sum = sum + day; // 加上某天的天数
	if (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0))  //判断是不是闰年
	{
		leap = 1;
	}
	else
	{
		leap = 0;
	}
	if ((leap == 1) && (month > 2)) //如果是闰年且月份大于2,总天数应该加一天
	{
		sum++;
	}
	printf("这是一年的第%d天\n",sum);

	return 0;
}

//NO.5
#include <stdio.h>

int main()
{
	int a = 0;
	int b = 0;
	int c = 0;
	// a>b>c
	printf("-------------------------------\n");
	printf("请输入三个整数数:\n");
	printf("-------------------------------\n");
	printf("空格分开\n");
	scanf("%d %d %d",&a,&b,&c);

	if ( a > b )  //交换 a和b 的值
	{
		int tmp = 0;
		tmp = a;   
		a = b;    
		b = tmp;   
	}
	if ( a > c )  
	{
		int tmp=0;
		tmp = a;    
		a = c;
		c = tmp;

	}
	if ( b > c ) 
	{
		int tmp = 0;
		tmp = b;
		b = c;
		c = tmp;

	}
	printf("%d  %d  %d", a, b, c);
	return 0;
}


//NO.6
#include<stdio.h>

//用* 号输出字母C的图案
int main()
{
	printf("     *\n");
	printf("  **    **\n");
	printf("**\n");
	printf("*\n");
	printf("**\n");
	printf("  **    **\n");
	printf("     *\n");

	return 0;
}

NO.7
#include<stdio.h>

//输出特殊图案，请在c环境中运行，看一看，Very Beautiful!

int main()
{
    int a, b;
    for (a = 0; a < 5; a++) //列
    {
        for (b = 0; b < 5; b++) //行
        {
            if (a == b || (a + b+1) == 5)
            {
                printf("■");
            }
            else
                printf("□");
            
        }
        printf("\n");
    }
    return 0;
}


//NO.8
#include<stdio.h>
//输出9*9口诀
int main()
{
	int i, j;
	for (i = 1; i <= 9;i++)
	{
		
		for (j = 1; j <= i; j++)
		{
			printf("%d*%d=%-3d",j, i, i*j);
		}
		printf("\n");
	}
	return 0;
}


NO.9
#include <stdio.h>
//要求输出国际象棋棋盘。
int main()
{
	int i,j;
    for (i = 0; i < 8; i++) // 控制列
    {
        for (j = 0; j < 8; j++)  // 控制行
        {
            if ((i + j) % 2 == 0)
                printf("■");
            else 
                printf("□");
        }
        printf("\n");
    }
}




//NO.10
#include<stdio.h>
// 打印楼梯，同时在楼梯上方打印两个笑脸
int main()
{
	int i,j;
	printf("♂♀\n");
	for (i = 1; i <= 11; i++)
	{
		for (j = 0; j < i; j++)
		{
			printf("■■");
		}
		printf("\n");
		
	}
	return 0;
}


// No.11
#include <stdio.h>
int main()
//古典问题（兔子生崽）：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？（输出前40个月即可）
//兔子的出生规律为 每月   1  ，1  ，3 ，5  ，8   ，13  ，21   即下个月是上两个月之和（从第三个月开始）
//1 月 2月 都有一只   五月的为 三月和四月之和
{
	long f1 = 1, f2 = 1;  // 兔子数量
	int i;  //循环
	int m =0; //计算月份

	printf("输入要计算的月份_:");
	scanf("%d",&m);
	printf("\n");

	if (m % 2 == 1)  // 按照月份计算循环次数
	{
		m = (m+1) / 2;
	}
	else
	{
		m = m / 2;
	}
	for (i = 1;i <= m; i++)
	{
		printf("第%d个月有%d只兔子\n",i*2-1,f1);
		printf("第%d个月有%d只兔子\n", i * 2, f2);
		f1 = f1 + f2; //前两个月加起来赋值给第三个月
		f2 = f1 + f2; //前两个月加起来赋值给第三个月
	}

	return 0;
}


// No.12
#include <stdio.h>
//判断 101 到 200 之间的素数。
// 用一个数分别去除 2 到 sqrt(这个数)，如果能被整除，则表明此数不是素数   只能被1 和 本身整除
int main()
{
	int i, j;
	int count=0;
	for (i = 101; i <= 200; i+=2) //偶数一定不是素数
	{
		for (j = 2; j < i; j++)
			if (i % j == 0) // 判断能否被整除   再跳出循环
				break;
		if (j >= i)   //判断循环是否提前跳出，如果 j < i 说明在 2~j 之间, i 有可整除的数
		{
			count ++;   //计数器
			printf("%d ", i);

			if (count % 5 == 0) //  换行，用 count 计数，每五个数换行
				printf("\n");
		}

	}
	printf("\n");
	printf("100-200 一共有%d 个素数\n",count);
}


// No.13
#include<stdio.h>

int main()
//打印出所有的"水仙花数"
//其各位数字立方和等于该数 本身。例如：153是一个"水仙花数"，因为153=1的三次方＋5的三次方＋3的三次方
{
	int i, x, y, z;
	int count=0;
	for (i = 100;i < 1000; i++) // 生成所有的三位数
	{
		x = i % 10;           //取个数
		y = i / 10 % 10;     //取十数
		z = i / 100 % 10;     //取百数

		if (i == (x * x * x + y * y * y + z * z * z))
		{
			count++;  //计数
			printf("%d ", i);
			
			if (count % 5 == 0) //每五个数换行
				printf("\n");
		}
	}
	printf("一共有%d个水仙花数\n",count);
	return 0;
}



// No.13
#include<stdio.h>

int main()
//打印出所有的"水仙花数"四位数
//其各位数字立方和等于该数 本身
{
	int i, x, y, z,q;
	int count=0;
	for (i = 1000;i < 10000; i++) // 生成所有的三位数
	{
		x = i % 10;           //取个数
		y = i / 10 % 10;     //取十数
		z = i / 100 % 10;     //取百数
		q = i / 1000 % 10;    //取千数
		if (i == (x * x * x * x + y * y * y * y + z * z * z * z + q * q * q* q))
		{
			count++;  //计数
			printf("%d ", i);
			
			if (count % 5 == 0) //每五个数换行
				printf("\n");
		}
	}
	printf("一共有%d个水仙花数\n",count);
	return 0;
}

//No.14
#include<stdio.h>
//将一个正整数分解质因数。例如：输入90, 打印出90 = 2 * 3 * 3 * 5
int main()
{
	int n, j;
	printf("请输入一个整数：");
	scanf("%d",&n);
	printf("%d=",n);

	for (j = 2; j <=n; j++)  // j大于2 小于  n
	{
		while (n % j == 0)     // n 能被j 整除 进入循环
		{
			printf("%d", j);   //打印第一次的商   
			n /= j;  // n = n/j;  n =  n /2  
			if (n !=  1)   // 商不为1   打印 乘号  
				printf("*");
		}
	}
	printf("\n");
	return 0;

}

//No.15
#include <stdio.h>

int main()
//条件运算符的嵌套来完成此题：学习成绩 >= 90分的同学用A表示，60 - 89分之间的用B表示，60分以下的用C表示
//(a > b) ? a : b这是条件运算符的基本例子。
{
	int score;
	char  grade;
	printf("请输入分数：");
	scanf("%d",&score);
	grade = (score >= 90) ? 'A' : ((score >= 60) ? 'B' : 'C');

	printf("%c \n",grade);
	return 0;
}

// No.16
#include <stdio.h>
//输入两个正整数m和n，求其最大公约数和最小公倍数。
int main()
{
	int a, b, r, n;
	printf("请输入两个数字:");
	scanf("%d %d", &a, &b);
	if (a < b)   // 数值交换   a>b;
	{
		int tmp;
		tmp = b;
		b = a;
		a = tmp;
	}
	r = a % b;  //  判断是否可以整除    11
	n = a * b;  // 可以整除 较大值为最大公约数  较小值为公倍数    420
	while (r != 0)  //判断 取模的值是否为0；
	{
		a = b;
		b = r;
		r = a % b;  //
	}
	printf("这两个数的最大公约数是%d，最小公倍数是%d\n", b, n / b);
	return 0;
}